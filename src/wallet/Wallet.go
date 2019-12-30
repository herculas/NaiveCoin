package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const walletFileName = "data/wallet.dat"

type Wallet struct {
	KeyPairs map[string]*KeyPair
}

func NewWallet() (*Wallet, error) {
	var wallet = Wallet{}
	wallet.KeyPairs = make(map[string]*KeyPair)
	var err = wallet.loadFromFile()
	return &wallet, err
}

func (wallet *Wallet) CreateKeyPair() string {
	var keyPair = NewKeyPair()
	var address = fmt.Sprintf("%s", keyPair.GetAddress())
	wallet.KeyPairs[address] = keyPair
	wallet.saveToFile()
	return address
}

func (wallet *Wallet) GetAddressList() []string {
	var addresses []string
	for address := range wallet.KeyPairs {
		addresses = append(addresses, address)
	}
	return addresses
}

func (wallet *Wallet) GetKeyPair(address string) KeyPair {
	return *wallet.KeyPairs[address]
}

func (wallet *Wallet) loadFromFile() error {
	var fileContent []byte
	var walletBuffer Wallet
	var err error
	if _, err = os.Stat(walletFileName); os.IsNotExist(err) {
		return err
	}
	if fileContent, err = ioutil.ReadFile(walletFileName); err != nil {
		log.Panic(err)
	}
	gob.Register(elliptic.P256())
	var decoder = gob.NewDecoder(bytes.NewReader(fileContent))
	if err = decoder.Decode(&walletBuffer); err != nil {
		log.Panic(err)
	}
	wallet.KeyPairs = walletBuffer.KeyPairs
	return nil
}

func (wallet *Wallet) saveToFile() {
	var content bytes.Buffer
	gob.Register(elliptic.P256())
	var encoder = gob.NewEncoder(&content)
	if err := encoder.Encode(wallet); err != nil {
		log.Panic(err)
	}
	if err := ioutil.WriteFile(walletFileName, content.Bytes(), 0644); err != nil {
		log.Panic(err)
	}
}
