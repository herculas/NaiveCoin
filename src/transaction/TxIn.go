package transaction

import (
	"fmt"
	"naivecoin-go/src/utils"
)

type TxIn struct {
	TxID       []byte
	TxOutIndex int64
	ScriptSig  string
}

func (txIn *TxIn) CanBeUnlockedByAddress(address string) bool {
	return txIn.ScriptSig == address
}

func (txIn *TxIn) Description() string {
	return fmt.Sprint("|    TxIn ID    |") +
		utils.FormatStrings(fmt.Sprintf("%x", txIn.TxID), 64) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+----------------------------------------------------------------+") +
		fmt.Sprint("|  TxOut Index  |") +
		utils.FormatIntegers(txIn.TxOutIndex, 64) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+----------------------------------------------------------------+") +
		fmt.Sprint("|   Signature   |") +
		utils.FormatStrings(txIn.ScriptSig, 64) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+----------------------------------------------------------------+")
}

func ConvertTxInFromUTxO(uTxO *UTxOut) *TxIn {
	return &TxIn{
		TxID:       uTxO.TxID,
		TxOutIndex: uTxO.TxOutIndex,
		ScriptSig:  uTxO.ScriptPubKey,
	}
}