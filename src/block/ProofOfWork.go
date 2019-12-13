package block

import (
	"../utils"
	"bytes"
	"crypto/sha256"
	"math/big"
)

const targetBit = 22

type ProofOfWork struct {
	block  *Block   // block to be verified
	target *big.Int // difficulty threshold in big integer storage
}

func (proofOfWork *ProofOfWork) MineBlock() ([]byte, int64) {
	var nonce int64 = 0
	var hash [32]byte
	var hashInt = new(big.Int)
	for {
		var dataBytes = bytes.Join([][]byte{
			proofOfWork.block.PreviousHash,
			proofOfWork.block.Data,
			utils.IntToHex(proofOfWork.block.Timestamp),
			utils.IntToHex(proofOfWork.block.Height),
			utils.IntToHex(nonce),
		}, []byte{})
		hash = sha256.Sum256(dataBytes)
		//fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if proofOfWork.target.Cmp(hashInt) == 1 {
			break
		}
		nonce++
	}
	//fmt.Println()
	return hash[:], nonce
}

func (proofOfWork *ProofOfWork) ValidateBlock() bool {
	var hashInt = new(big.Int)
	hashInt.SetBytes(proofOfWork.block.Hash)
	return proofOfWork.target.Cmp(hashInt) == 1
}

// create new ProofOfWork instance
func CreateProofOfWork(block *Block) *ProofOfWork {
	var target = big.NewInt(1)
	target = target.Lsh(target, 256 - targetBit)
	return &ProofOfWork{
		block:  block,
		target: target,
	}
}