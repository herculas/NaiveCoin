package transaction

import (
	"bytes"
	"fmt"
	"naivechain/src/util/coder"
	"naivechain/src/util/formatter"
)

type in struct {
	TxId       []byte
	TxOutIndex int64
	PublicKey  []byte
	Signature  []byte
}

func (txIn *in) usePublicKey(publicKeyHash []byte) bool {
	lockingHash := coder.HashPublicKey(txIn.PublicKey)
	return bytes.Compare(lockingHash, publicKeyHash) == 0
}

func (txIn *in) description(currentIndex int, totalIndex int) string {
	res := fmt.Sprint("┃  TxInID  │") +
		formatter.Strings(fmt.Sprintf("%x", txIn.TxId), 64) +
		fmt.Sprintln("┃") +
		fmt.Sprint("┃TxOutIndex│") +
		formatter.Integers(txIn.TxOutIndex, 64) +
		fmt.Sprintln("┃") +
		fmt.Sprint("┃  PubKey  │") +
		formatter.Strings(fmt.Sprintf("%x", txIn.PublicKey), 64) +
		fmt.Sprintln("┃") +
		fmt.Sprint("┃Signature │") +
		formatter.Strings(fmt.Sprintf("%x", txIn.Signature), 64) +
		fmt.Sprintln("┃")

	if currentIndex + 1 == totalIndex {
		res += fmt.Sprintln("┣━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	} else {
		res += fmt.Sprintln("┠──────────┴────────────────────────────────────────────────────────────────┨")
	}
	return res
}
