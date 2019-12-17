package block

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	Height       int64
	Timestamp    int64
	Nonce        int64
	PreviousHash []byte
	Hash         []byte
	Data         []byte
}

func (block *Block) Description() {
	fmt.Printf("Height: %d ", block.Height)
	fmt.Printf("Time: %s ", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
	fmt.Printf("Data: %s ", block.Data)
	fmt.Printf("Hash: %x ", block.Hash)
	fmt.Printf("PreviousHash: %x ", block.PreviousHash)
	fmt.Printf("Nonce: %d\n", block.Nonce)
}

func (block *Block) Serialize() []byte {
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

func CreateBlock(data string, height int64, previousHash []byte) *Block {
	var newBlock = &Block{
		Height:       height,
		Timestamp:    time.Now().Unix(),
		Nonce:        0,
		PreviousHash: previousHash,
		Hash:         nil,
		Data:         []byte(data),
	}
	var proofOfWork = createProofOfWork(newBlock)
	hash, nonce := proofOfWork.mineBlock()
	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce
	return newBlock
}

func CreateGenesisBlock(data string) *Block {
	return CreateBlock(data, 0, []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	})
}
