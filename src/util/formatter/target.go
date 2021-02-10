package formatter

import "math/big"

func Target(target int64) *big.Int {
	var res = big.NewInt(1)
	res = res.Lsh(res, uint(256 -target))
	return res
}
