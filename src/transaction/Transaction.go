package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"naivecoin-go/src/utils"
	"os"
)

const CoinbaseAmount = 10

type Transaction struct {
	TxID   []byte
	TxIns  []*TxIn
	TxOuts []*TxOut
}

func (transaction *Transaction) SetTxID() {
	var encoded bytes.Buffer
	var encoder = gob.NewEncoder(&encoded)
	if err := encoder.Encode(transaction); err != nil {
		log.Panic(err)
	}
	var hash = sha256.Sum256(encoded.Bytes())
	transaction.TxID = hash[:]
}

func (transaction *Transaction) Description() string {
	var res = fmt.Sprintln("") +
		fmt.Sprintln("+---------------+----------------------------------------------------------------+") +
		fmt.Sprint("|Transaction ID |") +
		utils.FormatStrings(fmt.Sprintf("%x", transaction.TxID), 64) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+----------------------------------------------------------------+")

	for index, txIn := range transaction.TxIns {
		res += fmt.Sprint("|") +
			utils.FormatStrings(fmt.Sprintf("Transaction Input %d", index), 80) +
			fmt.Sprintln("|") +
			fmt.Sprintln("+---------------+----------------------------------------------------------------+") +
			txIn.Description()
	}
	for index, txOut := range transaction.TxOuts {
		res += fmt.Sprint("|") +
			utils.FormatStrings(fmt.Sprintf("Transaction Output %d", index), 80) +
			fmt.Sprintln("|") +
			fmt.Sprintln("+---------------+----------------------------------------------------------------+") +
			txOut.Description()
	}
	return res
}

func CreateCoinbaseTransaction(address string) *Transaction {
	var txIn = &TxIn{
		TxID:       []byte{},
		TxOutIndex: -1,
		ScriptSig:  "",
	}
	var txOut = &TxOut{
		Amount:       CoinbaseAmount,
		ScriptPubKey: address,
	}
	var transaction = &Transaction{
		TxID:   nil,
		TxIns:  []*TxIn{txIn},
		TxOuts: []*TxOut{txOut},
	}
	transaction.SetTxID()
	return transaction
}

func CreateNormalTransaction(uTxOs []*UTxOut, from string, to string, amount int64) *Transaction {
	var currentAmount int64 = 0
	var changeAmount int64 = 0
	var includedUTxOs []*UTxOut

	for _, uTxO := range uTxOs {
		includedUTxOs = append(includedUTxOs, uTxO)
		currentAmount += uTxO.Amount
		if currentAmount >= amount {
			changeAmount = currentAmount - amount
			break
		}
	}

	if currentAmount < amount {
		fmt.Println("Fatal: Amount Overflow.")
		fmt.Println("You do not have enough amount to transfer.")
		os.Exit(1)
	}

	var txIns []*TxIn
	var txOuts []*TxOut

	for _, uTxO := range includedUTxOs {
		txIns = append(txIns, ConvertTxInFromUTxO(uTxO))
	}
	txOuts = append(txOuts, &TxOut{
		Amount:       amount,
		ScriptPubKey: to,
	})
	if changeAmount != 0 {
		txOuts = append(txOuts, &TxOut{
			Amount:       changeAmount,
			ScriptPubKey: from,
		})
	}

	var transaction = &Transaction{
		TxID:   nil,
		TxIns:  txIns,
		TxOuts: txOuts,
	}
	transaction.SetTxID()
	return transaction
}