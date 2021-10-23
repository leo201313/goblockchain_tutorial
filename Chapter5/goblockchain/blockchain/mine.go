package blockchain

import (
	"fmt"
	"goblockchain/utils"
)

func (bc *BlockChain) RunMine() {
	transactionPool := CreateTransactionPool()
	//In the near future, we'll have to validate the transactions first here.
	candidateBlock := CreateBlock(bc.LastHash, transactionPool.PubTx) //PoW has been done here.
	if candidateBlock.ValidatePoW() {
		bc.AddBlock(candidateBlock)
		err := RemoveTransactionPoolFile()
		utils.Handle(err)
		return
	} else {
		fmt.Println("Block has invalid nonce.")
		return
	}
}
