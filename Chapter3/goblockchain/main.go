//main.go
package main

import (
	"fmt"
	"goblockchain/blockchain"
	"goblockchain/transaction"
)

func main() {
	txPool := make([]*transaction.Transaction, 0)
	var tempTx *transaction.Transaction
	var ok bool
	var property int
	chain := blockchain.CreateBlockChain()
	property, _ = chain.FindUTXOs([]byte("Leo Cao"))
	fmt.Println("Balance of Leo Cao: ", property)

	tempTx, ok = chain.CreateTransaction([]byte("Leo Cao"), []byte("Krad"), 100)
	if ok {
		txPool = append(txPool, tempTx)
	}
	chain.AddBlock(txPool)
	txPool = make([]*transaction.Transaction, 0)
	property, _ = chain.FindUTXOs([]byte("Leo Cao"))
	fmt.Println("Balance of Leo Cao: ", property)

	tempTx, ok = chain.CreateTransaction([]byte("Krad"), []byte("Exia"), 200) // this transaction is invalid
	if ok {
		txPool = append(txPool, tempTx)
	}

	tempTx, ok = chain.CreateTransaction([]byte("Krad"), []byte("Exia"), 50)
	if ok {
		txPool = append(txPool, tempTx)
	}

	tempTx, ok = chain.CreateTransaction([]byte("Leo Cao"), []byte("Exia"), 100)
	if ok {
		txPool = append(txPool, tempTx)
	}
	chain.AddBlock(txPool)
	txPool = make([]*transaction.Transaction, 0)
	property, _ = chain.FindUTXOs([]byte("Leo Cao"))
	fmt.Println("Balance of Leo Cao: ", property)
	property, _ = chain.FindUTXOs([]byte("Krad"))
	fmt.Println("Balance of Krad: ", property)
	property, _ = chain.FindUTXOs([]byte("Exia"))
	fmt.Println("Balance of Exia: ", property)

	for _, block := range chain.Blocks {
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("nonce: %d\n", block.Nonce)
		fmt.Println("Proof of Work validation:", block.ValidatePoW())
	}

	//I want to show the bug at this version.

	tempTx, ok = chain.CreateTransaction([]byte("Krad"), []byte("Exia"), 30)
	if ok {
		txPool = append(txPool, tempTx)
	}

	tempTx, ok = chain.CreateTransaction([]byte("Krad"), []byte("Leo Cao"), 30)
	if ok {
		txPool = append(txPool, tempTx)
	}

	chain.AddBlock(txPool)
	txPool = make([]*transaction.Transaction, 0)

	for _, block := range chain.Blocks {
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("nonce: %d\n", block.Nonce)
		fmt.Println("Proof of Work validation:", block.ValidatePoW())
	}

	property, _ = chain.FindUTXOs([]byte("Leo Cao"))
	fmt.Println("Balance of Leo Cao: ", property)
	property, _ = chain.FindUTXOs([]byte("Krad"))
	fmt.Println("Balance of Krad: ", property)
	property, _ = chain.FindUTXOs([]byte("Exia"))
	fmt.Println("Balance of Exia: ", property)
}
