package command

import (
	"fmt"
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