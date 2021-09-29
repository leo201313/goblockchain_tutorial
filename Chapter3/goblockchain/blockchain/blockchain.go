//blockchain.go
package blockchain

import (
	"encoding/hex"
	"fmt"
	"goblockchain/transaction"
	"goblockchain/utils"
)

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) AddBlock(txs []*transaction.Transaction) {
	newBlock := CreateBlock(bc.Blocks[len(bc.Blocks)-1].Hash, txs)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func CreateBlockChain() *BlockChain {
	blockchain := BlockChain{}
	blockchain.Blocks = append(blockchain.Blocks, GenesisBlock())
	return &blockchain
}

func (bc *BlockChain) FindUnspentTransactions(address []byte) []transaction.Transaction {
	var unSpentTxs []transaction.Transaction
	spentTxs := make(map[string][]int) // can't use type []byte as key value
	for idx := len(bc.Blocks) - 1; idx >= 0; idx-- {
		block := bc.Blocks[idx]
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		IterOutputs:
			for outIdx, out := range tx.Outputs {
				if spentTxs[txID] != nil {
					for _, spentOut := range spentTxs[txID] {
						if spentOut == outIdx {
							continue IterOutputs
						}
					}
				}

				if out.ToAddressRight(address) {
					unSpentTxs = append(unSpentTxs, *tx)
				}
			}
			if !tx.IsBase() {
				for _, in := range tx.Inputs {
					if in.FromAddressRight(address) {
						inTxID := hex.EncodeToString(in.TxID)
						spentTxs[inTxID] = append(spentTxs[inTxID], in.OutIdx)
					}
				}
			}
		}
	}
	return unSpentTxs
}

func (bc *BlockChain) FindSpendableOutputs(address []byte, amount int) (int, map[string]int) {
	unspentOuts := make(map[string]int)
	unspentTxs := bc.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTxs {
		txID := hex.EncodeToString(tx.ID)
		for outIdx, out := range tx.Outputs {
			if out.ToAddressRight(address) && accumulated < amount {
				accumulated += out.Value
				unspentOuts[txID] = outIdx
				if accumulated >= amount {
					break Work
				}
				continue Work // one transaction can only have one output referred to adderss
			}
		}
	}
	return accumulated, unspentOuts
}

func (bc *BlockChain) FindUTXOs(address []byte) (int, map[string]int) {
	unspentOuts := make(map[string]int)
	unspentTxs := bc.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTxs {
		txID := hex.EncodeToString(tx.ID)
		for outIdx, out := range tx.Outputs {
			if out.ToAddressRight(address) {
				accumulated += out.Value
				unspentOuts[txID] = outIdx
				continue Work // one transaction can only have one output referred to adderss
			}
		}
	}
	return accumulated, unspentOuts

}

func (bc *BlockChain) CreateTransaction(from, to []byte, amount int) (*transaction.Transaction, bool) {
	var inputs []transaction.TxInput
	var outputs []transaction.TxOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
	if acc < amount {
		fmt.Println("Not enough coins!")
		return &transaction.Transaction{}, false
	}
	for txid, outidx := range validOutputs {
		txID, err := hex.DecodeString(txid)
		utils.Handle(err)
		input := transaction.TxInput{txID, outidx, from}
		inputs = append(inputs, input)
	}

	outputs = append(outputs, transaction.TxOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, transaction.TxOutput{acc - amount, from})
	}
	tx := transaction.Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx, true
}

// func (bc*BlockChain)ValidateTransaction(tx *transaction.Transaction) bool{
// 	if !bytes.Equal(tx.ID,tx.TxHash()) {  //validate hash
// 		return false
// 	}

// 	fromaddress := tx.Inputs[0].FromAddress
// 	validOutputs :=
// }

// func (bc *BlockChain)ValidateTransactions(txs []*transaction.Transaction)(bool,int){
// 	for idx,tx := range txs{

// 	}

// }

func (bc *BlockChain) Mine(txs []*transaction.Transaction) {
	bc.AddBlock(txs)
}
