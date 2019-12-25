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

	var initCmd = flag.NewFlagSet("init", flag.ExitOnError)
	var createTransactionCmd = flag.NewFlagSet("send", flag.ExitOnError)
	var printChainCmd = flag.NewFlagSet("printChain", flag.ExitOnError)
	var printTransCmd = flag.NewFlagSet("printTrans", flag.ExitOnError)
	var getBalanceCmd = flag.NewFlagSet("getBalance", flag.ExitOnError)
	var helpCmd = flag.NewFlagSet("help", flag.ExitOnError)

	var initArgs = initCmd.String("address", "", "Address of Who Generates Genesis Block")
	// FIXME: It seems that a specific user should not initialize a transaction with others as a source
	var sendFromArgs = createTransactionCmd.String("from", "", "Transaction Source Address")
	var sendToArgs = createTransactionCmd.String("to", "", "Transaction Destination Address")
	var sendAmountArgs = createTransactionCmd.String("amount", "", "Transaction Amount")
	var balanceAddressArgs = getBalanceCmd.String("address", "", "Address of Account")

	switch os.Args[1] {
	case "init":
		if err := initCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "send":
		if err := createTransactionCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "printChain":
		if err := printChainCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "printTrans":
		if err := printTransCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	case "getBalance":
		if err := getBalanceCmd.Parse(os.Args[2:]); err != nil {
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

	if initCmd.Parsed() {
		initChain(*initArgs)
	}

	if createTransactionCmd.Parsed() {
		initTransaction(*sendFromArgs, *sendToArgs, *sendAmountArgs)
	}
	if printChainCmd.Parsed() {
		printChain()
	}
	if printTransCmd.Parsed() {
		printTransactions()
	}
	if getBalanceCmd.Parsed() {
		getBalance(*balanceAddressArgs)
	}
	if helpCmd.Parsed() {
		printUsage()
		os.Exit(0)
	}
}
