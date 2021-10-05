package blockchain

import (
	"fmt"
	"goblockchain/utils"
)

func (bc *BlockChain) RunMine() {
	candidateBlock := CreateCandidateBlock()
	//In the near future, we'll have to validate the transactions first here.
	block := CreateBlock(bc.LastHash, candidateBlock.PubTx) //PoW has been done here.
	if block.ValidatePoW() {
		bc.AddBlock(block)
		err := RemoveCandidateBlockFile()
		utils.Handle(err)
		return
	} else {
		fmt.Println("Block has invalid nonce.")
		return
	}
}
