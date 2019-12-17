package command

import (
	"flag"
	"fmt"
	"log"
	"naivecoin-go/src/blockchain"
	"os"
)

func printUsage(state int) {
	if state == 1 {
		fmt.Println("Illegal arguments. Please re-enter the command line arguments.")
	}
	fmt.Println("Usage and commands:")
	fmt.Println("\tinit \t\t-data BLOCK_DATA \t-- initialize a blockchain with a genesis block")
	fmt.Println("\taddBlock \t-data BLOCK_DATA \t-- add a block to chain")
	fmt.Println("\tprintChain \t\t\t\t-- print all the blocks in chain")
	fmt.Println("\thelp \t\t\t\t\t-- hint of this program")
}

func validateArgs() {
	if len(os.Args) < 2 {
		printUsage(0)
		os.Exit(1)
	}
}

func initChain(data string) {
	blockchain.InitializeBlockchain(data)
}

func addBlock(data string) {
	blockchain.GetBlockchain().AddBlock(data)
}

func printChain() {
	blockchain.GetBlockchain().Description()
}

func Run() {
	validateArgs()
	var initCmd = flag.NewFlagSet("init", flag.ExitOnError)
	var addBlockCmd = flag.NewFlagSet("addBlock", flag.ExitOnError)
	var printChainCmd = flag.NewFlagSet("printChain", flag.ExitOnError)
	var helpCmd = flag.NewFlagSet("help", flag.ExitOnError)

	var initArgs = initCmd.String("data", "", "Transaction Data of Genesis Block")
	var addBlockArgs = addBlockCmd.String("data", "", "Transaction Data")

	switch os.Args[1] {
	case "init":
		if err := initCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "addBlock":
		if err := addBlockCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "printChain":
		if err := printChainCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "help":
		if err := helpCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	default:
		printUsage(1)
		os.Exit(1)
	}
	if initCmd.Parsed() {
		if *initArgs == "" {
			printUsage(1)
			os.Exit(1)
		}
		initChain(*initArgs)
	}
	if addBlockCmd.Parsed() {
		if *addBlockArgs == "" {
			printUsage(1)
			os.Exit(1)
		}
		addBlock(*addBlockArgs)
	}
	if printChainCmd.Parsed() {
		printChain()
	}
	if helpCmd.Parsed() {
		printUsage(0)
		os.Exit(0)
	}
}
