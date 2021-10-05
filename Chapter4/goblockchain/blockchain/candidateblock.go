package blockchain

import (
	"bytes"
	"encoding/gob"
	"goblockchain/constcoe"
	"goblockchain/transaction"
	"goblockchain/utils"
	"io/ioutil"
	"os"
)

type CandidateBlock struct {
	PubTx []*transaction.Transaction
}

func (cb *CandidateBlock) SaveFile() {
	var content bytes.Buffer
	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(cb)
	utils.Handle(err)
	err = ioutil.WriteFile(constcoe.CandidateBlockFile, content.Bytes(), 0644)
	utils.Handle(err)
}

func (cb *CandidateBlock) LoadFile() error {
	if !utils.FileExists(constcoe.CandidateBlockFile) {
		return nil
	}

	var candidateBlock CandidateBlock

	fileContent, err := ioutil.ReadFile(constcoe.CandidateBlockFile)
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(bytes.NewBuffer(fileContent))
	err = decoder.Decode(&candidateBlock)

	if err != nil {
		return err
	}

	cb.PubTx = candidateBlock.PubTx
	return nil
}

func CreateCandidateBlock() *CandidateBlock {
	candidateblock := CandidateBlock{}
	err := candidateblock.LoadFile()
	utils.Handle(err)
	return &candidateblock
}

func (cb *CandidateBlock) AddTransaction(tx *transaction.Transaction) {
	cb.PubTx = append(cb.PubTx, tx)
}

func RemoveCandidateBlockFile() error {
	err := os.Remove(constcoe.CandidateBlockFile)
	return err
}
