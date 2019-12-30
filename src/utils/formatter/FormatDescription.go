package formatter

import (
	"fmt"
	"strings"
)

func FormatIntegers(num int64, space int) string {
	return FormatStrings(fmt.Sprintf("%d", num), space)
}

func FormatStrings(str string, spaceLength int) string {
	var length = len(str)
	if spaceLength < length {
		var leftLength = (spaceLength - 3) / 2
		var rightLength = spaceLength - 3 - leftLength
		var leftStr = str[0:leftLength]
		var rightStr = str[length-rightLength:]
		return leftStr + "..." + rightStr
	}
	var leftSpace = (spaceLength - length) / 2
	var rightSpace = spaceLength - leftSpace - length
	return strings.Repeat(" ", leftSpace) + str + strings.Repeat(" ", rightSpace)
}
