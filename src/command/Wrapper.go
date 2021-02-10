package command

import (
	"fmt"
	"naivechain/src/blockchain"
	"naivechain/src/transaction"
	"naivechain/src/util/coder"
	"naivechain/src/wallet"
	"os"
	"strconv"
)

func printUsage() {
	fmt.Println("Usage and commands:")
	fmt.Println("Mutate:")
	fmt.Println("\tinitchain \t\t-address \tADDRESS \t-- initialize a chain with a genesis block")
	fmt.Println("\tcreateaddress \t\t\t\t\t\t-- generate a new key-pair and save it into the wallet file")
	fmt.Println("\tsend \t\t\t-from \t\tSOURCE_ADDR")
	fmt.Println("\t\t\t\t-to \t\tDEST_ADDR")
	fmt.Println("\t\t\t\t-amount \tAMOUNT \t\t-- initialize a transaction")
	fmt.Println("Inquire:")
	fmt.Println("\tprintchain \t\t\t\t\t\t-- print all the blocks in chain")
	fmt.Println("\tprinttrans \t\t\t\t\t\t-- print all the transactions in chain")
	fmt.Println("\tprintbalance \t\t-address \tADDRESS \t-- get balance of specific account")
	fmt.Println("\tprintaddresses \t\t\t\t\t\t-- list all addresses from the wallet file")
	fmt.Println("Help:")
	fmt.Println("\thelp \t\t\t\t\t\t\t-- hint of this program")
}

func initChain(address string) {
	if !coder.ValidateAddress(address) {
		fmt.Println("Fatal: illegal address.")
		os.Exit(1)
	}
	coinbaseTransaction := transaction.CreateCoinbase(address)
	blockchain.Create(coinbaseTransaction)
}

func createAddress() {
	address := wallet.NewKey()
	fmt.Printf("Your new address: %s\n", address)
}

func createTransaction(from string, to string, amount string) {
	if !coder.ValidateAddress(from) || !coder.ValidateAddress(to) {
		fmt.Println("Fatal: illegal address")
		os.Exit(1)
	}

	var amountInt int64
	var err error
	if amountInt, err = strconv.ParseInt(amount, 10, 64); err != nil {
		fmt.Println("Fatal: illegal amount")
		fmt.Println("Please enter legal arguments suggested below.")
		os.Exit(1)
	}
	normalTx := transaction.CreateNormal(from, to, amountInt)

	chain := blockchain.GetFromDatabase()
	chain.AddBlock([]*transaction.Transaction{normalTx})
}

func printChain() {
	chain := blockchain.GetFromDatabase()
	description := chain.Description()
	fmt.Print(description)
}

func printAddresses() {
	addresses := wallet.FindAllKeys()
	for _, address := range addresses {
		fmt.Println(address)
	}
}

func printTransactions() {
	chain := blockchain.GetFromDatabase()
	fmt.Print(chain.DescribeTransactions())
}

func printUTxOuts() {
	transaction.ListAllUTxOuts()
}

func printBalance(address string) {
	if !coder.ValidateAddress(address) {
		fmt.Println("Fatal: illegal address")
		os.Exit(1)
	}
	fmt.Println(transaction.GetTotalAmount([]byte(address)))
}