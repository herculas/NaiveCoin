package blockchain

import (
	"math/big"
	"naivecoin-go/src/block"
	"naivecoin-go/src/database"
)

type Iterator struct {
	currentHash []byte
}

func (iterator *Iterator) next() *block.Block {
	var currentBlockBytes = database.Retrieve(iterator.currentHash)
	var currentBlock = block.Deserialize(currentBlockBytes)
	iterator.currentHash = currentBlock.PreviousHash
	return currentBlock
}

func (iterator *Iterator) hasNext() bool {
	var previousHashInt = new(big.Int).SetBytes(iterator.currentHash)
	return big.NewInt(0).Cmp(previousHashInt) != 0
}
