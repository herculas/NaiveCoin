package transaction

import (
	"bytes"
	"fmt"
	"naivecoin-go/src/utils/formatter"
	"naivecoin-go/src/wallet"
)

type TxOut struct {
	Amount     int64
	PubKeyHash []byte
}

func NewTxOut(amount int64, address string) *TxOut {
	var txOut = &TxOut{
		Amount:     amount,
		PubKeyHash: nil,
	}
	txOut.lock(address)
	return txOut
}

func (txOut *TxOut) lock(address string) {
	txOut.PubKeyHash = wallet.DecodeAddress(address)
}

func (txOut *TxOut) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(txOut.PubKeyHash, pubKeyHash) == 0
}

func (txOut *TxOut) Description() string {
	return fmt.Sprint("|  Amount  |") +
		formatter.FormatIntegers(txOut.Amount, 21) +
		fmt.Sprintln("|") +
		fmt.Sprint("|PubKeyHash|") +
		formatter.FormatStrings(fmt.Sprintf("%x", txOut.PubKeyHash), 21) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+----------+---------------------+")
}
