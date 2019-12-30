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
	return fmt.Sprint("|    TxIn ID    |") +
		formatter.FormatStrings(fmt.Sprintf("%x", txIn.TxID), 64) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+------------+--------------+------------------------------------+") +
		fmt.Sprint("|  TxOut Index  |") +
		formatter.FormatIntegers(txIn.TxOutIndex, 12) +
		fmt.Sprint("|  Signature   |") +
		formatter.FormatStrings("", 36) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+------------+--------------+------------------------------------+")
}
