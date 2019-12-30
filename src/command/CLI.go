package command

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func validateArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func Run() {
	validateArgs()

	var initChainCmd = flag.NewFlagSet("initchain", flag.ExitOnError)
	var createAddCmd = flag.NewFlagSet("createaddress", flag.ExitOnError)
	var createTxCmd = flag.NewFlagSet("send", flag.ExitOnError)
	var printChainCmd = flag.NewFlagSet("printchain", flag.ExitOnError)
	var printTransCmd = flag.NewFlagSet("printtrans", flag.ExitOnError)
	var getBalanceCmd = flag.NewFlagSet("getbalance", flag.ExitOnError)
	var listAddressCmd = flag.NewFlagSet("listaddress", flag.ExitOnError)
	var helpCmd = flag.NewFlagSet("help", flag.ExitOnError)

	var initAddrArgs = initChainCmd.String("address", "", "Address of Who Generates Genesis Block")
	var sendFromAddrArgs = createTxCmd.String("from", "", "Transaction Source Address")
	var sendToAddrArgs = createTxCmd.String("to", "", "Transaction Destination Address")
	var sendAmountArgs = createTxCmd.String("amount", "", "Transaction Amount")
	var balanceAddrArgs = getBalanceCmd.String("address", "", "Address of Account")

	switch os.Args[1] {
	case "initchain":
		if err := initChainCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "createaddress":
		if err := createAddCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "send":
		if err := createTxCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "printchain":
		if err := printChainCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "printtrans":
		if err := printTransCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "getbalance":
		if err := getBalanceCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "listaddress":
		if err := listAddressCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "help":
		if err := helpCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	default:
		fmt.Println("Fatal: Undefined Arguments.")
		fmt.Println("Please enter legal arguments suggested below.")
		printUsage()
		os.Exit(1)
	}

	if initChainCmd.Parsed() {
		initChain(*initAddrArgs)
	}

	if createAddCmd.Parsed() {
		createAddress()
	}

	if createTxCmd.Parsed() {
		initTransaction(*sendFromAddrArgs, *sendToAddrArgs, *sendAmountArgs)
	}
	if printChainCmd.Parsed() {
		printChain()
	}
	if printTransCmd.Parsed() {
		printTransactions()
	}
	if getBalanceCmd.Parsed() {
		getBalance(*balanceAddrArgs)
	}
	if listAddressCmd.Parsed() {
		listAddresses()
	}
	if helpCmd.Parsed() {
		printUsage()
		os.Exit(0)
	}
}
