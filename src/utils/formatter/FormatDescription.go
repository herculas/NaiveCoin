package formatter

import (
	"fmt"
	"log"
	"strings"
)
 
func FormatIntegers(data int64, space int) string {
	return FormatStrings(fmt.Sprintf("%d", data), space)
}

func FormatStrings(data string, space int) string {
	var length = len(data)
	if space < length {
		log.Panic("Output glitch...")
	}
	var leftSpace = (space - length) / 2
	var rightSpace = space - leftSpace - length
	return strings.Repeat(" ", leftSpace) + data + strings.Repeat(" ", rightSpace)
}