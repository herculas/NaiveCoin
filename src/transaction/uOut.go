package transaction

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"naivechain/src/config"
	"naivechain/src/database"
	"naivechain/src/util/coder"
)

type UTxOut struct {
	TxId          []byte
	TxOutIndex    int64
	Amount        int64
	PublicKeyHash []byte
}

func (uTxOut *UTxOut) convertToIn(publicKey []byte) *in {
	return &in{
		TxId:       uTxOut.TxId,
		TxOutIndex: uTxOut.TxOutIndex,
		PublicKey:  publicKey,
		Signature:  nil,
	}
}

func (uTxOut *UTxOut) Description() {
	res := fmt.Sprintf("TxId: %x\n", uTxOut.TxId) +
		fmt.Sprintf("Index: %d\n", uTxOut.TxOutIndex) +
		fmt.Sprintf("Amount: %d\n", uTxOut.Amount) +
		fmt.Sprintf("PubKeyHash: %x", uTxOut.PublicKeyHash)
	fmt.Println(res)
}

func GetTotalAmount(address []byte) int64 {
	uTxOuts := findAllUTxOuts(address)
	totalAmount := int64(0)
	for uTxOutStr := range uTxOuts {
		uTxOut := uTxOutDeserialize([]byte(uTxOutStr))
		totalAmount += uTxOut.Amount
	}
	return totalAmount
}

func uTxOutSerialize(uTxOut UTxOut) []byte {
	blob := new(bytes.Buffer)
	encoder := gob.NewEncoder(blob)
	if err := encoder.Encode(uTxOut); err != nil {
		log.Panic(err)
	}
	return blob.Bytes()
}

func uTxOutDeserialize(blob []byte) *UTxOut {
	uTxOut := new(UTxOut)
	decoder := gob.NewDecoder(bytes.NewReader(blob))
	if err := decoder.Decode(uTxOut); err != nil {
		log.Panic(err)
	}
	return uTxOut
}

func mapSerialize(m map[string]bool) []byte {
	blob := new(bytes.Buffer)
	encoder := gob.NewEncoder(blob)
	if err := encoder.Encode(m); err != nil {
		log.Panic(err)
	}
	return blob.Bytes()
}

func mapDeserialize(blob []byte) *map[string]bool {
	m := new(map[string]bool)
	decoder := gob.NewDecoder(bytes.NewReader(blob))
	if err := decoder.Decode(m); err != nil {
		log.Panic(err)
	}
	return m
}

func findAllUTxOuts(address []byte) map[string]bool {
	uTxOutsBlob := database.Retrieve(config.UTxOBucketName, address)
	if uTxOutsBlob == nil || len(uTxOutsBlob) == 0 {
		return make(map[string]bool)
	}
	return *mapDeserialize(uTxOutsBlob)
}

func addUTxOut(txId []byte, txOutIndex int64, txOut out) {
	address := coder.GenerateAddress(txOut.PublicKeyHash)
	uTxOuts := findAllUTxOuts(address)
	newUTxOut := UTxOut{
		TxId:          txId,
		TxOutIndex:    txOutIndex,
		Amount:        txOut.Amount,
		PublicKeyHash: txOut.PublicKeyHash,
	}
	newUTxOutStr := uTxOutSerialize(newUTxOut)
	uTxOuts[string(newUTxOutStr[:])] = true
	database.Update(config.UTxOBucketName, address, mapSerialize(uTxOuts))
}

func deleteUTxOut(txIn in) {
	address := coder.GenerateAddress(coder.HashPublicKey(txIn.PublicKey))
	uTxOuts := findAllUTxOuts(address)
	for uTxOutStr := range uTxOuts {
		uTxOut := uTxOutDeserialize([]byte(uTxOutStr[:]))
		if bytes.Compare(txIn.TxId, uTxOut.TxId) == 0 && txIn.TxOutIndex == uTxOut.TxOutIndex {
			delete(uTxOuts, uTxOutStr)
			break
		}
		log.Panic("not found")
	}
	database.Update(config.UTxOBucketName, address, mapSerialize(uTxOuts))
}

func ListAllUTxOuts() {
	database.Iterate(config.UTxOBucketName, func(address, mapBlob []byte) error {
		fmt.Println("address: ", string(address))
		uTxOuts := *mapDeserialize(mapBlob)
		for uTxOutStr := range uTxOuts {
			uTxOut := uTxOutDeserialize([]byte(uTxOutStr[:]))
			uTxOut.Description()
		}
		return nil
	})
}