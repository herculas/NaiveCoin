package blockchain

import (
	"math/big"
	"naivechain/src/block"
	"naivechain/src/config"
	"naivechain/src/database"
)

type Iterator struct {
	currentHash []byte
}

func (iterator *Iterator) next() *block.Block {
	var currentBlockBlob = database.Retrieve(config.BlockBucketName, iterator.currentHash)
	var currentBlock = block.Deserialize(currentBlockBlob)
	iterator.currentHash = currentBlock.PreviousHash
	return currentBlock
}

func (iterator *Iterator) hashNext() bool {
	var previousHashInt = new(big.Int).SetBytes(iterator.currentHash)
	return big.NewInt(0).Cmp(previousHashInt) != 0
}