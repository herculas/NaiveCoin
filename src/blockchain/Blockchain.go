package blockchain

import (
	"fmt"
	"naivecoin-go/src/block"
	"naivecoin-go/src/database"
	"naivecoin-go/src/transaction"
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

func (blockchain *Blockchain) AddBlock(transactions []*transaction.Transaction) {
	var latestBlockBytes = database.Retrieve(blockchain.latestHash)
	var latestBlock = block.Deserialize(latestBlockBytes)
	var newBlock = block.CreateBlock(transactions, latestBlock.Height + 1, latestBlock.Hash)
	database.Update(newBlock.Hash, block.Serialize(newBlock))
	database.Update([]byte(cursorName), newBlock.Hash)
	blockchain.latestHash = newBlock.Hash
}

func (blockchain *Blockchain) Description() string {
	var description string
	var iterator = blockchain.iterator()
	var currentBlock *block.Block
	for iterator.hasNext() {
		currentBlock = iterator.next()
		description += currentBlock.Description()
	}
	return description
}

func (blockchain *Blockchain) DescribeTransactions() string {
	var description string
	var iterator = blockchain.iterator()
	var currentBlock *block.Block
	for iterator.hasNext() {
		currentBlock = iterator.next()
		description += fmt.Sprintf("Block: %d", currentBlock.Height)
		for _, tx := range currentBlock.Transactions {
			description += tx.Description()
		}
	}
	return description
}

func GetBlockchain() *Blockchain {
	return &Blockchain{
		latestHash: database.Retrieve([]byte(cursorName)),
	}
}

func InitializeBlockchain(coinbase *transaction.Transaction) {
	var genesisBlock = block.CreateGenesisBlock([]*transaction.Transaction{coinbase})
	database.Update(genesisBlock.Hash, block.Serialize(genesisBlock))
	database.Update([]byte(cursorName), genesisBlock.Hash)
}

func (blockchain *Blockchain) FindUTxOByAddress(pubKeyHash []byte) []*transaction.UTxOut {
	var iterator = blockchain.iterator()
	var currentBlock *block.Block
	var unspentTxOuts []*transaction.UTxOut
	var spentTxOuts = map[string]map[int64]bool{}
	for iterator.hasNext() {
		currentBlock = iterator.next()
		for _, tx := range currentBlock.Transactions {
			var txKey = fmt.Sprintf("%x", tx.TxID)
			for _, txIn := range tx.TxIns {
				if txIn.UsePubKey(pubKeyHash) {
					var txInKey = fmt.Sprintf("%x", txIn.TxID)
					if spentTxOuts[txInKey] == nil {
						spentTxOuts[txInKey] = map[int64]bool{}
					}
					spentTxOuts[txInKey][txIn.TxOutIndex] = true
				}
			}
			for index, txOut := range tx.TxOuts {
				if txOut.IsLockedWithKey(pubKeyHash) {
					if spentTxOuts[txKey] == nil || spentTxOuts[txKey][int64(index)] == false {
						var newUTxO = &transaction.UTxOut{
							TxID:       tx.TxID,
							TxOutIndex: int64(index),
							Amount:     txOut.Amount,
							PubKeyHash: txOut.PubKeyHash,
						}
						unspentTxOuts = append(unspentTxOuts, newUTxO)
					}
				}
			}
		}
	}
	return unspentTxOuts
}
