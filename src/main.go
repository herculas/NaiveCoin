package main

import (
	"naivecoin-go/src/block"
)

func main() {
	var chain = block.CreateBlockchainWithGenesisBlock()
	chain.AddBlockToBlockchain("Send 100 satoshi to Rui")
	chain.AddBlockToBlockchain("Send 200 satoshi to Rui")
	chain.Description()
}
