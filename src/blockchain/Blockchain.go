package blockchain

import (
	"fmt"
	"naivecoin-go/src/block"
	"naivecoin-go/src/database"
	"naivecoin-go/src/transaction"
)

const CursorName = "latest"

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
	database.Update(newBlock.Hash, newBlock.Serialize())
	database.Update([]byte(CursorName), newBlock.Hash)
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
		description += fmt.Sprintf("Block: %d\n", currentBlock.Height)
		for _, tx := range currentBlock.Transactions {
			description += tx.Description()
		}
	}
	return description
}

func GetBlockchain() *Blockchain {
	return &Blockchain{
		latestHash: database.Retrieve([]byte(CursorName)),
	}
}

func InitializeBlockchain(coinbase *transaction.Transaction) {
	var genesisBlock = block.CreateGenesisBlock([]*transaction.Transaction{coinbase})
	database.Update(genesisBlock.Hash, genesisBlock.Serialize())
	database.Update([]byte(CursorName), genesisBlock.Hash)
}

// TODO: UTxO on chain access (should be deprecated)
func (blockchain *Blockchain) FindUTxOByAddress(address string) []*transaction.UTxOut {
	var iterator = blockchain.iterator()
	var currentBlock *block.Block
	var unspentTxOuts []*transaction.UTxOut
	var spentTxOuts = map[string]map[int64]bool{}		// key: txID, value: [index, index, ...]
	for iterator.hasNext() {
		currentBlock = iterator.next()
		for _, tx := range currentBlock.Transactions {
			var txKey = fmt.Sprintf("%x", tx.TxID)
			for _, txIn := range tx.TxIns {
				if txIn.CanBeUnlockedByAddress(address) {
					var txInKey = fmt.Sprintf("%x", txIn.TxID)
					if spentTxOuts[txInKey] == nil {
						spentTxOuts[txInKey] = map[int64]bool{}
					}
					spentTxOuts[txInKey][txIn.TxOutIndex] = true
				}
			}
			for index, txOut := range tx.TxOuts {
				if txOut.CanBeUnlockedByAddress(address) {
					if spentTxOuts[txKey] == nil || spentTxOuts[txKey][int64(index)] == false {
						var newUTxO = &transaction.UTxOut{
							TxID:         tx.TxID,
							TxOutIndex:   int64(index),
							Amount:       txOut.Amount,
							ScriptPubKey: txOut.ScriptPubKey,
						}
						unspentTxOuts = append(unspentTxOuts, newUTxO)
					}
				}
			}
		}
	}
	return unspentTxOuts
}
