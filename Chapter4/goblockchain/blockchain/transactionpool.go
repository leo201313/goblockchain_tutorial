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

type TransactionPool struct {
	PubTx []*transaction.Transaction
}

func (tp *TransactionPool) SaveFile() {
	var content bytes.Buffer
	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(tp)
	utils.Handle(err)
	err = ioutil.WriteFile(constcoe.TransactionPoolFile, content.Bytes(), 0644)
	utils.Handle(err)
}

func (tp *TransactionPool) LoadFile() error {
	if !utils.FileExists(constcoe.TransactionPoolFile) {
		return nil
	}

	var transactionPool TransactionPool

	fileContent, err := ioutil.ReadFile(constcoe.TransactionPoolFile)
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(bytes.NewBuffer(fileContent))
	err = decoder.Decode(&transactionPool)

	if err != nil {
		return err
	}

	tp.PubTx = transactionPool.PubTx
	return nil
}

func CreateTransactionPool() *TransactionPool {
	transactionPool := TransactionPool{}
	err := transactionPool.LoadFile()
	utils.Handle(err)
	return &transactionPool
}

func (tp *TransactionPool) AddTransaction(tx *transaction.Transaction) {
	tp.PubTx = append(tp.PubTx, tx)
}

func RemoveTransactionPoolFile() error {
	err := os.Remove(constcoe.TransactionPoolFile)
	return err
}
