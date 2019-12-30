package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"naivecoin-go/src/utils/formatter"
	"naivecoin-go/src/wallet"
	"os"
)

const coinbaseAmount = 10

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
		formatter.FormatStrings(fmt.Sprintf("%x", transaction.TxID), 64) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+----------------------------------------------------------------+")

	for index, txIn := range transaction.TxIns {
		res += fmt.Sprint("|") +
			formatter.FormatStrings(fmt.Sprintf("Transaction Input %d", index), 80) +
			fmt.Sprintln("|") +
			fmt.Sprintln("+---------------+----------------------------------------------------------------+") +
			txIn.Description()
	}
	for index, txOut := range transaction.TxOuts {
		res += fmt.Sprint("|") +
			formatter.FormatStrings(fmt.Sprintf("Transaction Output %d", index), 80) +
			fmt.Sprintln("|") +
			fmt.Sprintln("+---------------+------------+--------------+------------------------------------+") +
			txOut.Description()
	}
	return res
}

func HashTransactions(transactions []*Transaction) []byte {
	var txHashes [][]byte
	for _, tx := range transactions {
		txHashes = append(txHashes, tx.TxID)
	}
	var result = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return result[:]
}

func CreateCoinbaseTransaction(address string) *Transaction {
	var txIn = &TxIn{
		TxID:       []byte{},
		TxOutIndex: -1,
		Signature:  nil,
		PubKey:     nil,
	}
	var txOut = NewTxOut(coinbaseAmount, address)
	var transaction = &Transaction{
		TxID:   nil,
		TxIns:  []*TxIn{txIn},
		TxOuts: []*TxOut{txOut},
	}
	transaction.SetTxID()
	return transaction
}

func CreateNormalTransaction(uTxOs []*UTxOut, fromAddress string, toAddress string, amount int64) *Transaction {
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

	var wlt, _ = wallet.NewWallet()
	var fromPubKey = wlt.GetKeyPair(fromAddress).PubKey

	for _, uTxO := range includedUTxOs {
		txIns = append(txIns, uTxO.ConvertToTxIn(fromPubKey))
	}
	txOuts = append(txOuts, NewTxOut(amount, toAddress))
	if changeAmount != 0 {
		txOuts = append(txOuts, NewTxOut(changeAmount, fromAddress))
	}
	var transaction = &Transaction{
		TxID:   nil,
		TxIns:  txIns,
		TxOuts: txOuts,
	}
	transaction.SetTxID()
	return transaction
}