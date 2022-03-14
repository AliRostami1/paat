package parser

import (
	"log"
	"reflect"
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
	currentSum := sumOfSlice(currentLengths)

	t := make([]int, 0, len(currentLengths))

	for i := range currentLengths {
		t = append(t, int(float64(currentLengths[i])/float64(currentSum)*float64(targetLength)))
	}

	tSum := sumOfSlice(t)
	for i := 0; i < targetLength-tSum; i += 1 {
		t[i] += 1
	}

	tSum = sumOfSlice(t)

	log.Printf("targetLength:%d | tSum:%d", targetLength, tSum)

	return t
}

func sumOfSlice(slice []int) (sum int) {
	for _, cl := range slice {
		sum += cl
	}
	return
}

func max(firstNum, secondNum int) int {
	if firstNum > secondNum {
		return firstNum
	}
	return secondNum
}

func extractInterface(in reflect.Value) reflect.Value {
	if in.Type().Kind() == reflect.Interface {
		return reflect.ValueOf(in.Interface())
	}
	return in
}

func calcBorder(parentBorder Border, row, col int) Border {
	var (
		top  int
		left int
	)
	if row == 0 {
		// if row is 0 the it does not have any cells on it's
		// top side and we need borders on that side
		top = parentBorder.Top
	}
	if col == 0 {
		// if col is 0 the it does not have any cells on it's
		//left side and we need borders on that side
		left = parentBorder.Left
	}

	return Border{
		Left:   left,
		Top:    top,
		Right:  1,
		Bottom: 1,
	}
}
