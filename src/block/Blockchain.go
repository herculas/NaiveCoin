package block

import (
	"naivecoin-go/src/database"
)

const cursorName = "latest"

type Blockchain struct {
	latestHash []byte
}

func (blockchain *Blockchain) AddBlockToBlockchain(data string) {
	var latestBlockBytes = database.Retrieve(blockchain.latestHash)
	var latestBlock = Deserialize(latestBlockBytes)
	var newBlock = createBlock(data, latestBlock.Height + 1, latestBlock.Hash)
	database.Update(newBlock.Hash, newBlock.Serialize())
	database.Update([]byte(cursorName), newBlock.Hash)
	blockchain.latestHash = newBlock.Hash
}

func (blockchain *Blockchain) Description() {
	var iterator = blockchain.Iterator()
	var currentBlock *Block
	for iterator.hasNext() {
		currentBlock = iterator.next()
		currentBlock.Description()
	}
}

func (blockchain *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{
		currentHash: blockchain.latestHash,
	}
}

func CreateBlockchainWithGenesisBlock() *Blockchain {
	var genesisBlock = createGenesisBlock("Genesis block")
	database.DeleteBucket()
	database.Update(genesisBlock.Hash, genesisBlock.Serialize())
	database.Update([]byte(cursorName), genesisBlock.Hash)
	return &Blockchain{latestHash: genesisBlock.Hash}
}
