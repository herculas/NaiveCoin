package command

import (
	"flag"
	"fmt"
	"log"
	"naivecoin-go/src/blockchain"
	"naivecoin-go/src/database"
	"naivecoin-go/src/transaction"
	"os"
	"strconv"
)

func printUsage() {
	fmt.Println("Usage and commands:")
	fmt.Println("\tinit \t\t-address \tADDRESS \t-- initialize a chain with a genesis block")
	fmt.Println("\tsend \t\t-from \t\tSOURCE_ADDR")
	fmt.Println("\t\t\t-to \t\tDEST_ADDR")
	fmt.Println("\t\t\t-amount \tAMOUNT \t\t-- initialize a transaction")
	fmt.Println("\tprintChain \t\t\t\t\t-- print all the blocks in chain")
	fmt.Println("\tprintTrans \t\t\t\t\t-- print all the transactions in chain")
	fmt.Println("\tgetBalance \t-address \tADDRESS \t-- get balance of specific account")
	fmt.Println("\thelp \t\t\t\t\t\t-- hint of this program")
}

func validateArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func initChain(address string) {
	if database.IsExist() {
		fmt.Println("Fatal: Database Status Exception.")
		fmt.Println("Blockchain has already existed.")
		os.Exit(1)
	}
	var coinbaseTransaction = transaction.CreateCoinbaseTransaction(address)
	blockchain.InitializeBlockchain(coinbaseTransaction)
}

func initTransaction(from string, to string, amount string) {
	var amountInt int64
	var err error
	if amountInt, err = strconv.ParseInt(amount, 10, 64); err != nil {
		fmt.Println("Fatal: Illegal Amount.")
		fmt.Println("Please enter legal arguments suggested below.")
		os.Exit(1)
	}
	var chain = blockchain.GetBlockchain()
	var uTxOs = chain.FindUTxOByAddress(from)
	var newTx = transaction.CreateNormalTransaction(uTxOs, from, to, amountInt)
	chain.AddBlock([]*transaction.Transaction{newTx})
}

func printChain() {
	var chain = blockchain.GetBlockchain()
	var description = chain.Description()
	fmt.Print(description)
}

func printTransactions() {
	var chain = blockchain.GetBlockchain()
	var description = chain.DescribeTransactions()
	fmt.Print(description)
}

func getBalance(address string) {
	var uTxOs = blockchain.GetBlockchain().FindUTxOByAddress(address)
	fmt.Println(transaction.GetTotalAmount(uTxOs))
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
