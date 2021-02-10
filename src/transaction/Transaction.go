package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"naivechain/src/config"
	"naivechain/src/util/formatter"
	"naivechain/src/wallet"
	"os"
)

type Transaction struct {
	Id   []byte
	Ins  []*in
	Outs []*out
}

func CreateCoinbase(address string) *Transaction {
	txIn := &in{
		TxId:       []byte{},
		TxOutIndex: -1,
		PublicKey:  nil,
		Signature:  nil,
	}
	txOut := createOut(config.CoinbaseReward, address)
	transaction := &Transaction{
		Id:   nil,
		Ins:  []*in{txIn},
		Outs: []*out{txOut},
	}
	transaction.setId()
	return transaction
}

func CreateNormal(from string, to string, amount int64) *Transaction {
	balance := GetTotalAmount([]byte(from))
	if balance < amount {
		fmt.Println("Fatal: amount overflow.")
		fmt.Println("You do not have enough amount to transfer.")
		os.Exit(1)
	}

	uTxOuts := findAllUTxOuts([]byte(from))
	currentAmount := int64(0)
	changeAmount := int64(0)

	var includedUTxOuts []*UTxOut
	for uTxOutStr := range uTxOuts {
		uTxOut := uTxOutDeserialize([]byte(uTxOutStr))
		includedUTxOuts = append(includedUTxOuts, uTxOut)
		currentAmount += uTxOut.Amount
		if currentAmount >= amount {
			changeAmount = currentAmount - amount
			break
		}
	}

	if currentAmount < amount {
		fmt.Println("Fatal: amount overflow.")
		fmt.Println("You do not have enough amount to transfer.")
		os.Exit(1)
	}
	var txIns []*in
	var txOuts []*out

	fromPublicKey := wallet.Load([]byte(from)).Public

	for _, uTxOut := range includedUTxOuts {
		txIns = append(txIns, uTxOut.convertToIn(fromPublicKey))
	}
	txOuts = append(txOuts, createOut(amount, to))
	if changeAmount != 0 {
		txOuts = append(txOuts, createOut(changeAmount, from))
	}

	transaction := &Transaction{
		Id:   nil,
		Ins:  txIns,
		Outs: txOuts,
	}
	transaction.setId()
	return transaction
}

func (transaction *Transaction) setId() {
	encoded := new(bytes.Buffer)
	encoder := gob.NewEncoder(encoded)
	if err := encoder.Encode(transaction); err != nil {
		log.Panic(err)
	}
	hash := sha256.Sum256(encoded.Bytes())
	transaction.Id = hash[:]
}

func (transaction *Transaction) UpdateNormalUTxOs() {
	for _, txIn := range transaction.Ins {
		deleteUTxOut(*txIn)
	}
	for index, txOut := range transaction.Outs {
		addUTxOut(transaction.Id, int64(index), *txOut)
	}
}

func (transaction *Transaction) UpdateCoinbaseUTxOs() {
	for index, txOut := range transaction.Outs {
		addUTxOut(transaction.Id, int64(index), *txOut)
	}
}

func (transaction *Transaction) Description(blockHeight int64) string {
	inCount := len(transaction.Ins)
	outCount := len(transaction.Outs)
	var res = fmt.Sprintln("┏━━━━━━━━━━┯━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓") +
		fmt.Sprint("┃ Block No │") +
		formatter.Integers(blockHeight, 64) +
		fmt.Sprintln("┃") +
		fmt.Sprintln("┠──────────┼────────────────────────────────────────────────────────────────┨") +
		fmt.Sprint("┃   TxID   │") +
		formatter.Strings(fmt.Sprintf("%x", transaction.Id), 64) +
		fmt.Sprintln("┃") +
		fmt.Sprintln("┣━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")

	for index, txIn := range transaction.Ins {
		res += fmt.Sprint("┃") +
			formatter.Strings(fmt.Sprintf("Input %d", index), 75) +
			fmt.Sprintln("┃") +
			fmt.Sprintln("┠──────────┬────────────────────────────────────────────────────────────────┨") +
			txIn.description(index, inCount)
	}
	for index, txOut := range transaction.Outs {
		res += fmt.Sprint("┃") +
			formatter.Strings(fmt.Sprintf("Output %d", index), 75) +
			fmt.Sprintln("┃") +
			fmt.Sprintln("┠──────────┬────────────────────────────────────────────────────────────────┨") +
			txOut.description(index, outCount)
	}
	return res
}

// TODO: should be Merkle Tree
func HashTransactions(transactions []*Transaction) []byte {
	var txHashes [][]byte
	for _, tx := range transactions {
		txHashes = append(txHashes, tx.Id)
	}
	res := sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return res[:]
}
