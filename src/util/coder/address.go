package coder

import (
	"bytes"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
	"naivechain/src/config"
)

func HashPublicKey(publicKey []byte) []byte {
	var publicSHA256 = sha256.Sum256(publicKey)
	var ripemd160Hasher = ripemd160.New()
	if _, err := ripemd160Hasher.Write(publicSHA256[:]); err != nil {
		log.Panic(err)
	}
	var publicRIPEMD160 = ripemd160Hasher.Sum(nil)
	return publicRIPEMD160
}

func checksum(payload []byte) []byte {
	innerDigest := sha256.Sum256(payload)
	outerDigest := sha256.Sum256(innerDigest[:])
	return outerDigest[:config.AddressChecksumLen]
}

func GenerateAddress(publicKeyHash []byte) []byte {
	versionedPayload := append([]byte{config.KeyVersion}, publicKeyHash...)
	checksum := checksum(versionedPayload)
	fullPayload := append(versionedPayload, checksum...)
	return Base58Encode(fullPayload)
}

func ValidateAddress(address string) bool {
	var decodedAddress = Base58Decode([]byte(address))
	var version = decodedAddress[0]
	var publicKeyHash = decodedAddress[1 : len(decodedAddress)-config.AddressChecksumLen]
	var actualChecksum = decodedAddress[len(decodedAddress)-config.AddressChecksumLen:]
	var targetChecksum = checksum(append([]byte{version}, publicKeyHash...))
	return bytes.Compare(actualChecksum, targetChecksum) == 0
}

func DecodeAddress(address string) []byte {
	decodedAddress := Base58Decode([]byte(address))
	return decodedAddress[1 : len(decodedAddress)-config.AddressChecksumLen]
}
