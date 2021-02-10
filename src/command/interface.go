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

func Run()  {
	validateArgs()

	var initChainCMD = flag.NewFlagSet("initchain", flag.ExitOnError)
	var createAddrCMD = flag.NewFlagSet("createaddress", flag.ExitOnError)
	var createTxCMD = flag.NewFlagSet("send", flag.ExitOnError)

	var printChainCMD = flag.NewFlagSet("printchain", flag.ExitOnError)
	var printAddrCMD = flag.NewFlagSet("printaddresses", flag.ExitOnError)
	var printTransCMD = flag.NewFlagSet("printtrans", flag.ExitOnError)
	var printUTxOutsCMD = flag.NewFlagSet("printutxo", flag.ExitOnError)
	var printBalanceCMD = flag.NewFlagSet("printbalance", flag.ExitOnError)
	var printHelpCMD = flag.NewFlagSet("help", flag.ExitOnError)

	var initAddrArgs = initChainCMD.String("address", "", "address of who generates genesis block")
	var sendFromAddrArgs = createTxCMD.String("from", "", "transaction source address")
	var sendToAddrArgs = createTxCMD.String("to", "", "transaction destination address")
	var sendAmountArgs = createTxCMD.String("amount", "", "transaction amount")
	var balanceAddrArgs = printBalanceCMD.String("address", "", "Address of acount")

	var argErr error
	switch os.Args[1] {
	case "initchain":
		argErr = initChainCMD.Parse(os.Args[2:])
	case "createaddress":
		argErr = createAddrCMD.Parse(os.Args[2:])
	case "send":
		argErr = createTxCMD.Parse(os.Args[2:])
	case "printchain":
		argErr = printChainCMD.Parse(os.Args[2:])
	case "printaddresses":
		argErr = printAddrCMD.Parse(os.Args[2:])
	case "printtrans":
		argErr = printTransCMD.Parse(os.Args[2:])
	case "printutxo":
		argErr = printUTxOutsCMD.Parse(os.Args[2:])
	case "printbalance":
		argErr = printBalanceCMD.Parse(os.Args[2:])
	case "help":
		argErr = printHelpCMD.Parse(os.Args[2:])
	default:
		fmt.Println("Fatal: undefined arguments.")
		fmt.Println("Please re-enter legal arguments suggests below.")
		printUsage()
		os.Exit(1)
	}
	if argErr != nil {
		log.Panic(argErr)
	}

	if initChainCMD.Parsed() {
		initChain(*initAddrArgs)
	}
	if createAddrCMD.Parsed() {
		createAddress()
	}
	if createTxCMD.Parsed() {
		createTransaction(*sendFromAddrArgs, *sendToAddrArgs, *sendAmountArgs)
	}
	if printChainCMD.Parsed() {
		printChain()
	}
	if printTransCMD.Parsed() {
		printTransactions()
	}
	if printAddrCMD.Parsed() {
		printAddresses()
	}
	if printUTxOutsCMD.Parsed() {
		printUTxOuts()
	}
	if printBalanceCMD.Parsed() {
		printBalance(*balanceAddrArgs)
	}
	if printHelpCMD.Parsed() {
		printUsage()
		os.Exit(0)
	}
}