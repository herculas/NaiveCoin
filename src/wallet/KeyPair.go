package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
	"naivecoin-go/src/utils/coder"
)

const version = byte(0x00)
const addressChecksumLen = 4

type KeyPair struct {
	PriKey ecdsa.PrivateKey
	PubKey []byte
}

func NewKeyPair() *KeyPair {
	priKey, pubKey := createKeyPair()
	return &KeyPair{
		PriKey: priKey,
		PubKey: pubKey,
	}
}

func (keyPair *KeyPair) GetAddress() []byte {
	var pubKeyHash = HashPubKey(keyPair.PubKey)
	var versionedPayload = append([]byte{version}, pubKeyHash...)
	var checksum = checksum(versionedPayload)
	var fullPayload = append(versionedPayload, checksum...)
	var address = coder.Base58Encode(fullPayload)
	return address
}

func ValidateAddress(address string) bool {
	var decodedAddress = coder.Base58Decode([]byte(address))
	var version = decodedAddress[0]
	var pubKeyHash = decodedAddress[1 : len(decodedAddress)-addressChecksumLen]
	var actualChecksum = decodedAddress[len(decodedAddress)-addressChecksumLen:]
	var targetChecksum = checksum(append([]byte{version}, pubKeyHash...))
	return bytes.Compare(actualChecksum, targetChecksum) == 0
}

func createKeyPair() (ecdsa.PrivateKey, []byte) {
	var priKey *ecdsa.PrivateKey
	var pubKey []byte
	var err error
	if priKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader); err != nil {
		log.Panic(err)
	}
	pubKey = append(priKey.PublicKey.X.Bytes(), priKey.PublicKey.Y.Bytes()...)
	return *priKey, pubKey
}

func HashPubKey(pubKey []byte) []byte {
	var publicSHA256 = sha256.Sum256(pubKey)
	var ripemd160Hasher = ripemd160.New()
	if _, err := ripemd160Hasher.Write(publicSHA256[:]); err != nil {
		log.Panic(err)
	}
	var publicRIPEMD160 = ripemd160Hasher.Sum(nil)
	return publicRIPEMD160
}

func DecodeAddress(address string) []byte {
	var decodedAddress = coder.Base58Decode([]byte(address))
	return decodedAddress[1 : len(decodedAddress)-addressChecksumLen]
}

func checksum(payload []byte) []byte {
	var innerDigest = sha256.Sum256(payload)
	var outerDigest = sha256.Sum256(innerDigest[:])
	return outerDigest[:addressChecksumLen]
}
