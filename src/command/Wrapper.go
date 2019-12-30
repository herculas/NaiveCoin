package command

import (
	"fmt"
	"log"
	"naivecoin-go/src/blockchain"
	"naivecoin-go/src/database"
	"naivecoin-go/src/transaction"
	"naivecoin-go/src/wallet"
	"os"
	"strconv"
)

func printUsage() {
	fmt.Println("Usage and commands:")
	fmt.Println("Mutate:")
	fmt.Println("\tinitchain \t-address \tADDRESS \t-- initialize a chain with a genesis block")
	fmt.Println("\tcreateaddress \t\t\t\t\t-- generate a new key-pair and save it into the wallet file")
	fmt.Println("\tsend \t\t-from \t\tSOURCE_ADDR")
	fmt.Println("\t\t\t-to \t\tDEST_ADDR")
	fmt.Println("\t\t\t-amount \tAMOUNT \t\t-- initialize a transaction")
	fmt.Println("Inquire:")
	fmt.Println("\tprintchain \t\t\t\t\t-- print all the blocks in chain")
	fmt.Println("\tprinttrans \t\t\t\t\t-- print all the transactions in chain")
	fmt.Println("\tgetbalance \t-address \tADDRESS \t-- get balance of specific account")
	fmt.Println("\tlistaddress \t\t\t\t\t-- list all addresses from the wallet file")
	fmt.Println("Help:")
	fmt.Println("\thelp \t\t\t\t\t\t-- hint of this program")
}

func initChain(address string) {
	if database.IsExist() {
		fmt.Println("Fatal: Database Status Exception.")
		fmt.Println("Blockchain has already existed.")
		os.Exit(1)
	}
	if !wallet.ValidateAddress(address) {
		fmt.Println("Fatal: Illegal Address.")
		os.Exit(1)
	}
	var coinbaseTransaction = transaction.CreateCoinbaseTransaction(address)
	blockchain.InitializeBlockchain(coinbaseTransaction)
}

func createAddress() {
	var wlt, _ = wallet.NewWallet()
	var address = wlt.CreateKeyPair()
	fmt.Printf("Your new address: %s\n", address)
}

func initTransaction(fromAddress string, toAddress string, amount string) {
	if !wallet.ValidateAddress(fromAddress) || !wallet.ValidateAddress(toAddress) {
		fmt.Println("Fatal: Illegal Address.")
		os.Exit(1)
	}
	var amountInt int64
	var err error
	if amountInt, err = strconv.ParseInt(amount, 10, 64); err != nil {
		fmt.Println("Fatal: Illegal Amount.")
		fmt.Println("Please enter legal arguments suggested below.")
		os.Exit(1)
	}
	var chain = blockchain.GetBlockchain()
	var fromPubKeyHash = wallet.DecodeAddress(fromAddress)
	var uTxOs = chain.FindUTxOByAddress(fromPubKeyHash)
	var newTx = transaction.CreateNormalTransaction(uTxOs, fromAddress, toAddress, amountInt)
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
	if !wallet.ValidateAddress(address) {
		fmt.Println("Fatal: Illegal Address.")
		os.Exit(1)
	}
	var pubKeyHash = wallet.DecodeAddress(address)
	var uTxOs = blockchain.GetBlockchain().FindUTxOByAddress(pubKeyHash)
	fmt.Println(transaction.GetTotalAmount(uTxOs))
}

func listAddresses() {
	var wlt *wallet.Wallet
	var err error
	if wlt, err = wallet.NewWallet(); err != nil {
		log.Panic(err)
	}
	var addressList = wlt.GetAddressList()
	for _, address := range addressList {
		fmt.Println(address)
	}
}