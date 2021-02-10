package transaction

import (
	"bytes"
	"fmt"
	"naivechain/src/util/coder"
	"naivechain/src/util/formatter"
)

type out struct {
	Amount        int64
	PublicKeyHash []byte
}

func createOut(amount int64, address string) *out {
	txOut := &out{
		Amount:        amount,
		PublicKeyHash: nil,
	}
	txOut.lock(address)
	return txOut
}

func (txOut *out) lock(address string) {
	txOut.PublicKeyHash = coder.DecodeAddress(address)
}

func (txOut *out) checkPublicKey(publicKeyHash []byte) bool {
	return bytes.Compare(txOut.PublicKeyHash, publicKeyHash) == 0
}

func (txOut *out) description(currentIndex int, totalIndex int) string {
	res := fmt.Sprint("┃  Amount  │") +
		formatter.Integers(txOut.Amount, 64) +
		fmt.Sprintln("┃") +
		fmt.Sprint("┃PubKeyHash│") +
		formatter.Strings(fmt.Sprintf("%x", txOut.PublicKeyHash), 64) +
		fmt.Sprintln("┃")
	if currentIndex + 1 == totalIndex {
		 res += fmt.Sprintln("┗━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	} else {
		res += fmt.Sprintln("┠──────────┴────────────────────────────────────────────────────────────────┨")
	}
	return res
}
