//block.go
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"goblockchain/transaction"
	"goblockchain/utils"
	"time"
)

type Block struct {
	Timestamp    int64
	Hash         []byte
	PrevHash     []byte
	Target       []byte
	Nonce        int64
	Transactions []*transaction.Transaction
}

func (b *Block) SetHash() {
	information := bytes.Join([][]byte{utils.ToHexInt(b.Timestamp), b.PrevHash, b.Target, utils.ToHexInt(b.Nonce), b.BackTrasactionSummary()}, []byte{})
	hash := sha256.Sum256(information)
	b.Hash = hash[:]
}

func CreateBlock(prevhash []byte, txs []*transaction.Transaction) *Block {
	block := Block{time.Now().Unix(), []byte{}, prevhash, []byte{}, 0, txs}
	block.Target = block.GetTarget()
	block.Nonce = block.FindNonce()
	block.SetHash()
	return &block
}

func GenesisBlock() *Block {
	tx := transaction.BaseTx([]byte("Leo Cao"))
	return CreateBlock([]byte{}, []*transaction.Transaction{tx})
}

func (b *Block) BackTrasactionSummary() []byte {
	txIDs := make([][]byte, 0)
	for _, tx := range b.Transactions {
		txIDs = append(txIDs, tx.ID)
	}
	summary := bytes.Join(txIDs, []byte{})

	return summary
}
