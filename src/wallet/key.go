package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"log"
	"naivechain/src/config"
	"naivechain/src/database"
	"naivechain/src/util/coder"
)

type Key struct {
	Private ecdsa.PrivateKey
	Public  []byte
}

func NewKey() []byte {
	privateKey, publicKey := createKey()
	key := &Key{
		Private: privateKey,
		Public:  publicKey,
	}
	save(*key)
	return key.getAddress()
}

func (key *Key) getAddress() []byte {
	return coder.GenerateAddress(coder.HashPublicKey(key.Public))
}

func save(key Key) {
	blob := new(bytes.Buffer)
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(blob)
	if err := encoder.Encode(key); err != nil {
		log.Panic(err)
	}
	database.Update(config.KeyBucketName, key.getAddress(), blob.Bytes())
}

func Load(address []byte) Key {
	key := new(Key)
	blob := database.Retrieve(config.KeyBucketName, address)
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(blob))
	if err := decoder.Decode(key); err != nil {
		log.Panic(err)
	}
	return *key
}

func FindAllKeys() []string {
	var addresses []string
	database.Iterate(config.KeyBucketName, func(address, key []byte) error {
		addresses = append(addresses, fmt.Sprintf("%s", address))
		return nil
	})
	return addresses
}

func createKey() (ecdsa.PrivateKey, []byte) {
	var privateKey *ecdsa.PrivateKey
	var publicKey []byte
	var err error
	if privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader); err != nil {
		log.Panic(err)
	}
	publicKey = append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	return *privateKey, publicKey
}