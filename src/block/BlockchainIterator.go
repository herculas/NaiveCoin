package block

import (
	"math/big"
	"naivecoin-go/src/database"
)

type BlockchainIterator struct {
	currentHash []byte
}

func (iterator *BlockchainIterator) next() *Block {
	var currentBlockBytes = database.Retrieve(iterator.currentHash)
	var currentBlock = Deserialize(currentBlockBytes)
	iterator.currentHash = currentBlock.PreviousHash
	return currentBlock
}

func (iterator *BlockchainIterator) hasNext() bool {
	var previousHashInt = new(big.Int)
	previousHashInt.SetBytes(iterator.currentHash)
	return big.NewInt(0).Cmp(previousHashInt) != 0
}
