package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"math/big"
	"naivechain/src/transaction"
	"naivechain/src/util/convertor"
	"naivechain/src/util/formatter"
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

// TODO: target adjustment
func Create(height int64, target int64, previousHash []byte, transactions []*transaction.Transaction) *Block {
	var block = &Block{
		Height:       height,
		Timestamp:    time.Now().Unix(),
		Nonce:        0,
		Target:       target,
		PreviousHash: previousHash,
		Hash:         nil,
		Transactions: transactions,
	}
	block.mine()
	return block
}

func (block *Block) mine() {
	nonce := int64(0)
	var hash [32]byte
	var hashInt = new(big.Int)
	for {
		dataBytes := bytes.Join([][]byte{
			block.PreviousHash,
			transaction.HashTransactions(block.Transactions),
			convertor.IntToHex(block.Timestamp),
			convertor.IntToHex(block.Height),
			convertor.IntToHex(nonce),
		}, []byte{})

		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if formatter.Target(block.Target).Cmp(hashInt) == 1 {
			break
		}
		nonce++
	}
	fmt.Println()
	block.Hash = hash[:]
	block.Nonce = nonce
}

func (block *Block) validate() bool {
	var hashInt = new(big.Int).SetBytes(block.Hash)
	return formatter.Target(block.Target).Cmp(hashInt) == 1
}

func Serialize(block *Block) []byte {
	blob := new(bytes.Buffer)
	encoder := gob.NewEncoder(blob)
	if err := encoder.Encode(block); err != nil {
		log.Panic(err)
	}
	return blob.Bytes()
}

func Deserialize(blob []byte) *Block {
	var block Block
	var decoder = gob.NewDecoder(bytes.NewReader(blob))
	if err := decoder.Decode(&block); err != nil {
		log.Panic(err)
	}
	return &block
}

func (block *Block) Description() string {
	return fmt.Sprintln("") +
		fmt.Sprintln("┏━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━┯━━━━━━━━━━━━━━━━━━━━━━━━━━━┓") +
		fmt.Sprint("┃ Block Height  │") +
		formatter.Integers(block.Height, 24) +
		fmt.Sprint("│   Time    │") +
		fmt.Sprint(formatter.Strings(time.Unix(block.Timestamp, 0).Format("2006-01-02 15:04:05"), 27)) +
		fmt.Sprintln("┃") +
		fmt.Sprintln("┠───────────────┼────────────────────────┴───────────┴───────────────────────────┨") +
		fmt.Sprint("┃     Hash      │") +
		fmt.Sprintf("%x", block.Hash) +
		fmt.Sprintln("┃") +
		fmt.Sprintln("┠───────────────┼────────────────────────────────────────────────────────────────┨") +
		fmt.Sprint("┃ Previous Hash │") +
		fmt.Sprintf("%x", block.PreviousHash) +
		fmt.Sprintln("┃") +
		fmt.Sprintln("┠───────────────┼────────────────────────────────────────────────────────────────┨") +
		fmt.Sprint("┃  Txs Digest   │") +
		fmt.Sprintf("%x", transaction.HashTransactions(block.Transactions)) +
		fmt.Sprintln("┃") +
		fmt.Sprintln("┠───────────────┼────────────────────────┬───────────┬───────────────────────────┨") +
		fmt.Sprint("┃    Target     │") +
		formatter.Integers(block.Target, 24) +
		fmt.Sprint("│   Nonce   │") +
		formatter.Integers(block.Nonce, 27) +
		fmt.Sprintln("┃") +
		fmt.Sprintln("┗━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
}