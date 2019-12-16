package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"naivecoin-go/src/utils"
)

const targetBit = 16

type ProofOfWork struct {
	block  *Block   // block to be verified
	target *big.Int // difficulty threshold in big integer storage
}

func (proofOfWork *ProofOfWork) mineBlock() ([]byte, int64) {
	var nonce int64 = 0
	var hash [32]byte
	var hashInt = new(big.Int)
	for {
		var dataBytes = bytes.Join([][]byte{
			proofOfWork.block.PreviousHash,
			proofOfWork.block.Data,
			utils.IntToHexBytes(proofOfWork.block.Timestamp),
			utils.IntToHexBytes(proofOfWork.block.Height),
			utils.IntToHexBytes(nonce),
		}, []byte{})
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if proofOfWork.target.Cmp(hashInt) == 1 {
			break
		}
		nonce++
	}
	fmt.Println()
	return hash[:], nonce
}

func (proofOfWork *ProofOfWork) validateBlock() bool {
	var hashInt = new(big.Int)
	hashInt.SetBytes(proofOfWork.block.Hash)
	return proofOfWork.target.Cmp(hashInt) == 1
}

func createProofOfWork(block *Block) *ProofOfWork {
	var target = big.NewInt(1)
	target = target.Lsh(target, 256 - targetBit)
	return &ProofOfWork{
		block:  block,
		target: target,
	}
}
