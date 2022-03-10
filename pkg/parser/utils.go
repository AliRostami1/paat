package parser

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func (c *Canvas) drawContent(s string, x, y, width, height int) {
	contentMatrix := parseContent(s, width-2, height-2)

	for line := x; line < x+height; line += 1 {
		// draw the borders first
		for col := y; col < y+width; col += 1 {
			// for first line and last line we draw horizontal borders
			if line == x || line == x+height {
				// for first and last rune of the first and last line we draw border conjunction (+)
				if col == y || col == y+width {
					c.Matrix[line][col] = '+'
				} else {
					c.Matrix[line][col] = '-'
				}
			} else {
				if col == y || col == y+width {
					c.Matrix[line][col] = '|'
				} else {
					c.Matrix[line][col] = contentMatrix[line-x][col-y]
				}
			}
		}
	}
}

func parseContent(s string, w, h int) [][]byte {
	var byteMatrix [][]byte = make([][]byte, h)
	for i := range byteMatrix {
		byteMatrix[i] = make([]byte, w)
	}

	for i := 0; i < h; i += 1 {
		if i == h/2 {
			byteMatrix[i] = centerString(s, w)
		} else {
			byteMatrix[i] = makeEmptyByteArray(w)
		}
	}

	return byteMatrix
}

func makeEmptyByteArray(length int) (b []byte) {
	for i := 0; i < length; i += 1 {
		b = append(b, ' ')
	}
	return
}

func centerString(s string, w int) (b []byte) {
	str := fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
	return []byte(str)
}

func applyHorizontalPadding(s string, w int) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
}

func stringWidth(str string) int {
	return utf8.RuneCountInString(str)
}

func stringHeight(str string) int {
	return strings.Count(str, "\n") + 1
}
