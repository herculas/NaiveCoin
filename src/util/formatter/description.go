package formatter

import (
	"fmt"
	"strings"
)

func Integers(num int64, space int) string {
	return Strings(fmt.Sprintf("%d", num), space)
}

func Strings(str string, spaceLength int) string {
	var length = len(str)

	if spaceLength < length {
		var left = (spaceLength - 3) / 2
		var right = (spaceLength - 3) - left
		var leftStr = str[0:left]
		var rightStr = str[length - right:]
		return leftStr + "..." + rightStr
	}

	var leftSpace = (spaceLength - length) / 2
	var rightSpace = (spaceLength - length) - leftSpace
	return strings.Repeat(" ", leftSpace) + str + strings.Repeat(" ", rightSpace)
}