package blockchain

import (
	"naivechain/src/block"
	"naivechain/src/config"
	"naivechain/src/database"
	"naivechain/src/transaction"
)

type Blockchain struct {
	latestHash []byte
}

func Create(coinbase *transaction.Transaction) {
	if database.CheckBucket(config.BlockBucketName) {
		database.DropBucket(config.BlockBucketName)
		database.DropBucket(config.UTxOBucketName)
	}
	genesisBlock := block.Create(0, config.InitialTarget, []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}, []*transaction.Transaction{coinbase})
	database.Update(config.BlockBucketName, genesisBlock.Hash, block.Serialize(genesisBlock))
	database.Update(config.BlockBucketName, []byte(config.BlockCursorName), genesisBlock.Hash)
	coinbase.UpdateCoinbaseUTxOs()
}

func (blockchain *Blockchain) iterator() *Iterator {
	return &Iterator{
		currentHash: blockchain.latestHash,
	}
}

func (blockchain *Blockchain) AddBlock(transactions []*transaction.Transaction) {
	latestBlockBlob := database.Retrieve(config.BlockBucketName, blockchain.latestHash)
	latestBlock := block.Deserialize(latestBlockBlob)
	newBlock := block.Create(latestBlock.Height + 1, config.InitialTarget, latestBlock.Hash, transactions)
	database.Update(config.BlockBucketName, newBlock.Hash, block.Serialize(newBlock))
	database.Update(config.BlockBucketName, []byte(config.BlockCursorName), newBlock.Hash)
	blockchain.latestHash = newBlock.Hash
	for _, tx := range transactions {
		tx.UpdateNormalUTxOs()
	}
}

func (blockchain *Blockchain) Description() string {
	var description string
	var currentBlock *block.Block
	iterator := blockchain.iterator()
	for iterator.hashNext() {
		currentBlock = iterator.next()
		description += currentBlock.Description()
	}
	return description
}

func (blockchain *Blockchain) DescribeTransactions() string {
	var description string
	var currentBlock *block.Block
	iterator := blockchain.iterator()
	for iterator.hashNext() {
		currentBlock = iterator.next()
		for _, tx := range currentBlock.Transactions {
			description += tx.Description(currentBlock.Height)
		}
	}
	return description
}

func GetFromDatabase() *Blockchain {
	return &Blockchain{
		latestHash: database.Retrieve(config.BlockBucketName, []byte(config.BlockCursorName)),
	}
}