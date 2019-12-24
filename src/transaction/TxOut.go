package transaction

import (
	"fmt"
	"naivecoin-go/src/utils"
)

type TxOut struct {
	Amount       int64
	ScriptPubKey string
}

func (txOut *TxOut) CanBeUnlockedByAddress(address string) bool {
	return txOut.ScriptPubKey == address
}

func (txOut *TxOut) Description() string {
	return fmt.Sprint("|    Amount     |") +
		utils.FormatIntegers(txOut.Amount, 64) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+----------------------------------------------------------------+") +
		fmt.Sprint("| ScriptPubKey  |") +
		utils.FormatStrings(txOut.ScriptPubKey, 64) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+----------------------------------------------------------------+")
}
