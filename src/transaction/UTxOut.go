package transaction

import "fmt"

type UTxOut struct {
	TxID         []byte
	TxOutIndex   int64
	Amount       int64
	ScriptPubKey string
}

func (uTxOut *UTxOut) Description() string {
	return fmt.Sprintf("TxID: %x\n", uTxOut.TxID) +
		fmt.Sprintf("Index: %d\n", uTxOut.TxOutIndex) +
		fmt.Sprintf("Amount: %d\n", uTxOut.Amount) +
		fmt.Sprintf("PubKey: %s\n", uTxOut.ScriptPubKey)
}

func GetTotalAmount(uTxOs []*UTxOut) int64 {
	var amount int64 = 0
	for _, uTxO := range uTxOs {
		amount += uTxO.Amount
	}
	return amount
}