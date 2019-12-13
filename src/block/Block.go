package block

import (
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

// create a new block
func CreateBlock(data string, height int64, previousHash []byte) *Block {
	var newBlock = &Block{
		Height:       height,
		Timestamp:    time.Now().Unix(),
		Nonce:        0,
		PreviousHash: previousHash,
		Hash:         nil,
		Data:         []byte(data),
	}
	// create a instance of ProofOfWork and generate hash with nonce
	var proofOfWork = CreateProofOfWork(newBlock)
	hash, nonce := proofOfWork.MineBlock()
	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce
	return newBlock
}

// create Genesis block
func CreateGenesisBlock(data string) *Block {
	return CreateBlock(data, 0, []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	})
}