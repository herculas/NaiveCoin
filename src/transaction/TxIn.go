package transaction

import (
	"bytes"
	"fmt"
	"naivecoin-go/src/utils/formatter"
	"naivecoin-go/src/wallet"
)

type TxIn struct {
	TxID       []byte
	TxOutIndex int64
	PubKey     []byte
	Signature  []byte
}

// determines whether the current input matches a specific output
func (txIn *TxIn) UsePubKey(pubKeyHash []byte) bool {
	var lockingHash = wallet.HashPubKey(txIn.PubKey)
	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

func (txIn *TxIn) Description() string {
	return fmt.Sprint("|  TxInID  |") +
		formatter.FormatStrings(fmt.Sprintf("%x", txIn.TxID), 21) +
		fmt.Sprintln("|") +
		fmt.Sprint("|TxOutIndex|") +
		formatter.FormatIntegers(txIn.TxOutIndex, 21) +
		fmt.Sprintln("|") +
		fmt.Sprint("|  PubKey  |") +
		formatter.FormatStrings(fmt.Sprintf("%x", txIn.PubKey), 21) +
		fmt.Sprintln("|") +
		fmt.Sprint("|Signature |") +
		formatter.FormatStrings(fmt.Sprintf("%x", txIn.Signature), 21) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+----------+---------------------+")
}
