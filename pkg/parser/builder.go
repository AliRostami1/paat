package parser

import (
	"strings"
)

type Canvas struct {
	Matrix [][]byte
}

func (c *Canvas) String() string {
	sb := strings.Builder{}
	for _, row := range c.Matrix {
		for _, col := range row {
			sb.WriteByte(col)
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func (cell *Cell) Draw() *Canvas {
	// build the matrix and initilize it
	var runeMatrix [][]byte = make([][]byte, cell.Height)
	for i := range runeMatrix {
		runeMatrix[i] = make([]byte, cell.Width)
	}

	canvas := &Canvas{
		Matrix: runeMatrix,
	}

	canvas.drawer(cell, 0, 0)

	return canvas
}

func (c *Canvas) drawer(cell *Cell, x, y int) {
	if cell.Type == PrimitiveCell {
		c.drawContent(cell.Content.(string), x, y, cell.Width, cell.Height)
	} else {
		table := cell.Content.(*Table)
		currentX := x
		for i := 0; i < len(table.Cells); i += 1 {
			if i > 0 {
				currentX += table.Cells[i][0].Height
			}
			currentY := y
			for j := 0; j < len(table.Cells[i]); j += 1 {
				if j > 0 {
					currentY += table.Cells[i][j].Width
				}
				c.drawer(table.Cells[i][j], currentX, currentY)
			}
		}
	}
}

func (c *Canvas) drawContent(s string, x, y, width, height int) {
	contentMatrix := parseContent(s, width, height)

	for contentRowIndex, row := range contentMatrix {
		for contentColIndex, col := range row {
			c.Matrix[contentRowIndex+x][contentColIndex+y] = col
		}
	}
}

func parseContent(s string, w, h int) [][]byte {
	var byteMatrix [][]byte = make([][]byte, h)
	for i := range byteMatrix {
		byteMatrix[i] = make([]byte, w)
	}

	byteMatrix[0][0] = '+'
	byteMatrix[0][w-1] = '+'
	byteMatrix[h-1][0] = '+'
	byteMatrix[h-1][w-1] = '+'

	for i := 1; i < w-1; i += 1 {
		byteMatrix[0][i] = '-'
		byteMatrix[h-1][i] = '-'
	}

	lines := strings.Split(s, "\n")

	var (
		leadingEmptyLines  int
		trailingEmptyLines int
	)
	mod, emptyLinesEach := (h-2-len(lines))%2, (h-2-len(lines))/2
	leadingEmptyLines = emptyLinesEach + mod
	trailingEmptyLines = emptyLinesEach

	for i := 1; i < h-1; i += 1 {
		if i <= leadingEmptyLines || i >= h-1-trailingEmptyLines {
			byteMatrix[i] = makeEmptyByteArray(w)
		} else {
			byteMatrix[i] = makeContentByteArray(lines[i-1-leadingEmptyLines], w)
		}
	}

	return byteMatrix
}

func makeEmptyByteArray(length int) (b []byte) {
	b = append(b, '|')
	for i := 1; i < length-1; i += 1 {
		b = append(b, ' ')
	}
	b = append(b, '|')
	return
}

func makeContentByteArray(s string, w int) (b []byte) {
	str := []byte(s)
	var (
		leadingEmptyBytes  int
		trailingEmptyBytes int
	)
	mod, emptyBytesEach := (w-2-len(s))%2, (w-2-len(s))/2
	leadingEmptyBytes = emptyBytesEach + mod
	trailingEmptyBytes = emptyBytesEach

	b = append(b, '|')

	for i := 1; i < w-1; i += 1 {
		if i <= leadingEmptyBytes || i >= w-1-trailingEmptyBytes {
			b = append(b, ' ')
		} else {
			b = append(b, str[i-leadingEmptyBytes-1])
		}
	}

	b = append(b, '|')
	return
}
