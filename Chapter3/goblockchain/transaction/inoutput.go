package transaction

import "bytes"

type TxOutput struct {
	Value     int
	ToAddress []byte
}

type TxInput struct {
	TxID        []byte
	OutIdx      int
	FromAddress []byte
}

func (in *TxInput) FromAddressRight(address []byte) bool {
	return bytes.Equal(in.FromAddress, address)
}

func (out *TxOutput) ToAddressRight(address []byte) bool {
	return bytes.Equal(out.ToAddress, address)
}
