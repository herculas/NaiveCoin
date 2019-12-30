package coder

import (
	"bytes"
	"math/big"
)

var alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte {
	var result []byte
	var inputInt = new(big.Int).SetBytes(input)
	var base = big.NewInt(int64(len(alphabet)))
	var mod = &big.Int{}
	for big.NewInt(0).Cmp(inputInt) != 0 {
		inputInt.DivMod(inputInt, base, mod)
		result = append(result, alphabet[mod.Int64()])
	}
	reverseBytes(result)
	for b := range input {
		if b == 0x00 {
			result = append([]byte{alphabet[0]}, result...)
		} else {
			break
		}
	}
	return result
}

func Base58Decode(input []byte) []byte {
	var result = big.NewInt(0)
	var zeroBytes = 0
	for b := range input {
		if b == 0x00 {
			zeroBytes++
		}
	}
	var payload = input[zeroBytes:]
	for _, b := range payload {
		var charIndex = bytes.IndexByte(alphabet, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	var decoded = result.Bytes()
	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)
	return decoded
}

func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
