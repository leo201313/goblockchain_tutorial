package transaction

import (
	"bytes"
	"encoding/gob"
	"goblockchain/utils"
)

type TxOutput struct {
	Value      int
	HashPubKey []byte
}

type TxInput struct {
	TxID   []byte
	OutIdx int
	PubKey []byte
	Sig    []byte
}

type UTXO struct {
	TxID   []byte
	OutIdx int
	TxOutput
}

func (in *TxInput) FromAddressRight(address []byte) bool {
	return bytes.Equal(in.PubKey, address)
}

func (out *TxOutput) ToAddressRight(address []byte) bool {
	return bytes.Equal(out.HashPubKey, utils.PublicKeyHash(address))
}

func (u *UTXO) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(u)
	utils.Handle(err)
	return res.Bytes()
}

func DeserializeUTXO(data []byte) *UTXO {
	var utxo UTXO
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&utxo)
	utils.Handle(err)
	return &utxo
}
