package wallet

import (
	"bytes"
	"fmt"
	"goblockchain/blockchain"
	"goblockchain/constcoe"
	"goblockchain/transaction"
	"goblockchain/utils"
	"goblockchain/utxoset"

	"github.com/dgraph-io/badger"
)

func (wt *Wallet) GetUtxoSetDir() string {
	strAddress := string(wt.Address())
	dirAddress := constcoe.UTXOSet + strAddress
	return dirAddress
}

func (wt *Wallet) CreateUTXOSet(chain *blockchain.BlockChain) *utxoset.UTXOSet {
	UTXOs := chain.BackUTXOs(wt.PublicKey)
	utxoSet := utxoset.CreateUTXOSet(wt.Address(), wt.GetUtxoSetDir(), UTXOs, chain.BackHeight())
	return utxoSet
}

func (wt *Wallet) LoadUTXOSet() *utxoset.UTXOSet {
	utxoSet := utxoset.LoadUTXOSet(wt.GetUtxoSetDir())
	return utxoSet
}

func (wt *Wallet) GetBalance() int {
	amount := 0
	us := wt.LoadUTXOSet()
	defer us.DB.Close()

	err := us.DB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			if utxoset.IsInfo(item.Key()) {
				continue
			}
			err := item.Value(func(v []byte) error {
				tmpUTXO := transaction.DeserializeUTXO(v)
				amount += tmpUTXO.Value
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	utils.Handle(err)
	return amount
}

func (w *Wallet) ScanBlock(block *blockchain.Block) {
	utxoSet := w.LoadUTXOSet()
	defer utxoSet.DB.Close()

	if block.Height > (utxoSet.Height + 1) {
		fmt.Println("UTXO Set is out of date!")
		return
	}

	for _, tx := range block.Transactions {
		for _, in := range tx.Inputs {
			if bytes.Equal(in.PubKey, w.PublicKey) {
				utxoSet.DelUTXO(in.TxID, in.OutIdx)
			}
		}

		for outIdx, out := range tx.Outputs {
			if bytes.Equal(out.HashPubKey, utils.PublicKeyHash(w.PublicKey)) {
				utxoSet.AddUTXO(tx.ID, outIdx, out)
			}
		}
	}
	utxoSet.UpdateHeight(block.Height)
}
