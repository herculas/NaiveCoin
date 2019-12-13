package block

type Blockchain struct {
	Blocks []*Block
}

func (blockchain *Blockchain) AddBlockToBlockchain(data string) {
	var height = blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1
	var previousHash = blockchain.Blocks[len(blockchain.Blocks) - 1].Hash
	var newBlock = CreateBlock(data, height, previousHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

// create a blockchain with genesis block
func CreateBlockchainWithGenesisBlock() *Blockchain {
	var genesisBlock = CreateGenesisBlock("Genesis block")
	return &Blockchain{
		Blocks: []*Block{genesisBlock},
	}
}
