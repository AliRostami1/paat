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

func distributeLength(targetLength int, currentLengths []int) []int {
	var sum int
	for _, cl := range currentLengths {
		sum += cl
	}

	t := make([]int, 0, len(currentLengths))

	for i := range currentLengths {
		t = append(t, int(float64(currentLengths[i])/float64(sum)*float64(targetLength)))
	}

	return t
}

func max(firstNum, secondNum int) int {
	if firstNum > secondNum {
		return firstNum
	}
	return secondNum
}
