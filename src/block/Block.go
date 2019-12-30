package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"math/big"
	"naivecoin-go/src/transaction"
	"naivecoin-go/src/utils/convertor"
	"naivecoin-go/src/utils/formatter"
	"time"
)

type Block struct {
	Height       int64
	Timestamp    int64
	Nonce        int64
	Target       int64
	PreviousHash []byte
	Hash         []byte
	Transactions []*transaction.Transaction
}

func (block *Block) mineBlock() {
	var nonce int64 = 0
	var hash [32]byte
	var hashInt = new(big.Int)
	for {
		var dataBytes = bytes.Join([][]byte{
			block.PreviousHash,
			transaction.HashTransactions(block.Transactions),
			convertor.IntToHexBytes(block.Timestamp),
			convertor.IntToHexBytes(block.Height),
			convertor.IntToHexBytes(nonce),
		}, []byte{})
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if formatter.FormatTarget(block.Target).Cmp(hashInt) == 1 {
			break
		}
		nonce++
	}
	fmt.Println()
	block.Hash = hash[:]
	block.Nonce = nonce
}

func (block *Block) Validate() bool {
	var hashInt = new(big.Int).SetBytes(block.Hash)
	return formatter.FormatTarget(block.Target).Cmp(hashInt) == 1
}

func (block *Block) Description() string {
	return fmt.Sprintln("") +
		fmt.Sprintln("+---------------+--------------------------------+--------+----------------------+") +
		fmt.Sprint("| Block Height  |") +
		formatter.FormatIntegers(block.Height, 32) +
		fmt.Sprint("|  Time  |") +
		fmt.Sprint(time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM")) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+--------------------------------+--------+----------------------+") +
		fmt.Sprint("|  Txs Digest   |") +
		fmt.Sprintf("%x", transaction.HashTransactions(block.Transactions)) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+----------------------------------------------------------------+") +
		fmt.Sprint("|     Hash      |") +
		fmt.Sprintf("%x", block.Hash) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+----------------------------------------------------------------+") +
		fmt.Sprint("| Previous Hash |") +
		fmt.Sprintf("%x", block.PreviousHash) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+------------------------+-----------+---------------------------+") +
		fmt.Sprint("|    Target     |") +
		formatter.FormatIntegers(block.Target, 24) +
		fmt.Sprint("|   Nonce   |") +
		formatter.FormatIntegers(block.Nonce, 27) +
		fmt.Sprintln("|") +
		fmt.Sprintln("+---------------+------------------------+-----------+---------------------------+") +
		fmt.Sprintln("")
}

func Serialize(block *Block) []byte {
	var result bytes.Buffer
	var encoder = gob.NewEncoder(&result)
	if err := encoder.Encode(block); err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func Deserialize(blockBytes []byte) *Block {
	var block Block
	var decoder = gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&block); err != nil {
		log.Panic(err)
	}
	return &block
}

// TODO: Target negotiation
func CreateBlock(transactions []*transaction.Transaction, height int64, previousHash []byte) *Block {
	var newBlock = &Block{
		Height:       height,
		Timestamp:    time.Now().Unix(),
		Nonce:        0,
		Target:       16,
		PreviousHash: previousHash,
		Hash:         nil,
		Transactions: transactions,
	}
	newBlock.mineBlock()
	return newBlock
}

func CreateGenesisBlock(transactions []*transaction.Transaction) *Block {
	return CreateBlock(transactions, 0, []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	})
}
