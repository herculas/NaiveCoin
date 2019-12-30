package formatter

import "math/big"

func FormatTarget(target int64) *big.Int {
	var targetValue = big.NewInt(1)
	targetValue = targetValue.Lsh(targetValue, uint(256-target))
	return targetValue
}