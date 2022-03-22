//block.go
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"goblockchain/merkletree"
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
	MTree        *merkletree.MerkleTree
}

func (b *Block) SetHash() {
	information := bytes.Join([][]byte{utils.ToHexInt(b.Timestamp), b.PrevHash, b.Target, utils.ToHexInt(b.Nonce), b.BackTrasactionSummary(), b.MTree.RootNode.Data}, []byte{})
	hash := sha256.Sum256(information)
	b.Hash = hash[:]
}

func CreateBlock(prevhash []byte, txs []*transaction.Transaction) *Block {
	block := Block{time.Now().Unix(), []byte{}, prevhash, []byte{}, 0, txs, merkletree.CrateMerkleTree(txs)}
	block.Target = block.GetTarget()
	block.Nonce = block.FindNonce()
	block.SetHash()
	return &block
}

func GenesisBlock(address []byte) *Block {
	tx := transaction.BaseTx(address)
	genesis := CreateBlock([]byte("Leo Cao is awesome!"), []*transaction.Transaction{tx})
	genesis.SetHash()
	return genesis
}

func (b *Block) BackTrasactionSummary() []byte {
	txIDs := make([][]byte, 0)
	for _, tx := range b.Transactions {
		txIDs = append(txIDs, tx.ID)
	}
	summary := bytes.Join(txIDs, []byte{})
	return summary
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	utils.Handle(err)
	return res.Bytes()
}

func DeSerializeBlock(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	utils.Handle(err)
	return &block
}
