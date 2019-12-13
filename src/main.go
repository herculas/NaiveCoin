package main

import (
	"./block"
	"fmt"
)

func main() {
	var chain = block.CreateBlockchainWithGenesisBlock()

	chain.AddBlockToBlockchain("Send 100 satoshi to Rui")
	chain.AddBlockToBlockchain("Send 200 satoshi to Rui")

	fmt.Println(chain.Blocks)
}
