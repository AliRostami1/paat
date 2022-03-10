package parser

import (
	"strings"
	"unicode/utf8"
)

func stringWidth(str string) int {
	return utf8.RuneCountInString(str)
}

func stringHeight(str string) int {
	return strings.Count(str, "\n") + 1
}
