package blockchain

import (
	"fmt"
	"naivecoin-go/src/block"
	"naivecoin-go/src/database"
	"os"
)

const cursorName = "latest"

type Blockchain struct {
	latestHash []byte
}

func (blockchain *Blockchain) iterator() *Iterator {
	return &Iterator{
		currentHash: blockchain.latestHash,
	}
}

func (blockchain *Blockchain) AddBlock(data string) {
	var latestBlockBytes = database.Retrieve(blockchain.latestHash)
	var latestBlock = block.Deserialize(latestBlockBytes)
	var newBlock = block.CreateBlock(data, latestBlock.Height + 1, latestBlock.Hash)
	database.Update(newBlock.Hash, newBlock.Serialize())
	database.Update([]byte(cursorName), newBlock.Hash)
	blockchain.latestHash = newBlock.Hash
}

func (blockchain *Blockchain) Description() {
	var iterator = blockchain.iterator()
	var currentBlock *block.Block
	for iterator.hasNext() {
		currentBlock = iterator.next()
		currentBlock.Description()
	}
}

func InitializeBlockchain(data string) {
	if database.Exists() {
		fmt.Println("Fatal: blockchain has already existed.")
		os.Exit(1)
	}
	var genesisBlock = block.CreateGenesisBlock(data)
	database.Update(genesisBlock.Hash, genesisBlock.Serialize())
	database.Update([]byte(cursorName), genesisBlock.Hash)
}

func GetBlockchain() *Blockchain {
	if !database.Exists() {
		fmt.Println("Fatal: blockchain does not exist.")
		os.Exit(1)
	}
	return &Blockchain{
		latestHash: database.Retrieve([]byte(cursorName)),
	}
}