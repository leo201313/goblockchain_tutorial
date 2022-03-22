package transaction

import (
	"bytes"
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

func (in *TxInput) FromAddressRight(address []byte) bool {
	return bytes.Equal(in.PubKey, address)
}

func (out *TxOutput) ToAddressRight(address []byte) bool {
	return bytes.Equal(out.HashPubKey, utils.PublicKeyHash(address))
}
