package transaction

import "fmt"

type UTxOut struct {
	TxID       []byte
	TxOutIndex int64
	Amount     int64
	PubKeyHash []byte
}

func (uTxOut *UTxOut) Description() string {
	return fmt.Sprintf("TxID: %x\n", uTxOut.TxID) +
		fmt.Sprintf("Index: %d\n", uTxOut.TxOutIndex) +
		fmt.Sprintf("Amount: %d\n", uTxOut.Amount) +
		fmt.Sprintf("PubKey: %s\n", uTxOut.PubKeyHash)
}

func (uTxOut *UTxOut) ConvertToTxIn(pubKey []byte) *TxIn {
	return &TxIn{
		TxID:       uTxOut.TxID,
		TxOutIndex: uTxOut.TxOutIndex,
		Signature:  nil,
		PubKey:     pubKey,
	}
}

func GetTotalAmount(uTxOs []*UTxOut) int64 {
	var amount int64 = 0
	for _, uTxO := range uTxOs {
		amount += uTxO.Amount
	}
	return amount
}