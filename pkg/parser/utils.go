package parser

import (
	"strings"
	"unicode/utf8"
)

func stringWidth(str string) int {
	var maxLineWidth int

	lines := strings.Split(str, "\n")
	for _, line := range lines {
		maxLineWidth = max(maxLineWidth, utf8.RuneCountInString(line))
	}
	return maxLineWidth
}

func stringHeight(str string) int {
	return strings.Count(str, "\n") + 1
}
