package parser

import (
	"log"
	"strings"
)

type Canvas struct {
	Matrix [][]byte
}

func (c *Canvas) String() string {
	sb := strings.Builder{}
	sb.WriteByte('\n')
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
	var runeMatrix [][]byte = make([][]byte, cell.BorderBoxHeight())
	for i := range runeMatrix {
		runeMatrix[i] = make([]byte, cell.BorderBoxWidth())
	}

	canvas := &Canvas{
		Matrix: runeMatrix,
	}

	canvas.drawer(cell, 0, 0)

	return canvas
}

func (c *Canvas) drawer(cell *Cell, x, y int) {
	if cell.Type == PrimitiveCell {
		c.drawContent(cell, x, y)
	} else {
		table := cell.Content.(*Table)
		currentX := x
		for i := range table.Cells {
			if i > 0 {
				currentX += table.Cells[i-1][0].BorderBoxHeight()
			}
			currentY := y
			for j := range table.Cells[i] {
				if j > 0 {
					currentY += table.Cells[i][j-1].BorderBoxWidth()
				}

				c.drawer(table.Cells[i][j], currentX, currentY)
			}
		}
	}
}

func (c *Canvas) drawContent(cell *Cell, x, y int) {
	contentMatrix := parseContent(cell)

	for contentRowIndex, row := range contentMatrix {
		for contentColIndex, col := range row {
			c.Matrix[contentRowIndex+x][contentColIndex+y] = col
		}
	}
}

func parseContent(cell *Cell) [][]byte {
	s, ok := cell.Content.(string)
	if !ok {
		log.Printf("s: %v", cell.Content)
	}

	byteMatrix := [][]byte{}

	if cell.Border.Top > 0 {
		byteMatrix = append(byteMatrix, makeHorizontalBorder(cell.Width, cell.Border))
	}

	lines := strings.Split(s, "\n")

	var (
		leadingEmptyLines  int
		trailingEmptyLines int
	)
	mod, emptyLinesEach := (cell.Height-len(lines))%2, (cell.Height-len(lines))/2
	leadingEmptyLines = emptyLinesEach + mod
	trailingEmptyLines = emptyLinesEach

	for i := 0; i < cell.Height; i += 1 {
		if i < leadingEmptyLines || i >= cell.Height-trailingEmptyLines {
			byteMatrix = append(byteMatrix, makeEmptyByteArray(cell.Width, cell.Border))
		} else {
			byteMatrix = append(byteMatrix, makeContentByteArray(lines[i-leadingEmptyLines], cell.Width, cell.Border))
		}
	}

	if cell.Border.Bottom > 0 {
		byteMatrix = append(byteMatrix, makeHorizontalBorder(cell.Width, cell.Border))
	}

	return byteMatrix
}

func makeHorizontalBorder(length int, border Border) (b []byte) {
	if border.Left > 0 {
		b = append(b, '+')
	}
	for i := 0; i < length; i += 1 {
		b = append(b, '-')
	}
	if border.Right > 0 {
		b = append(b, '+')
	}
	return
}

func makeEmptyByteArray(length int, border Border) (b []byte) {
	if border.Left > 0 {
		b = append(b, '|')
	}
	for i := 0; i < length; i += 1 {
		b = append(b, ' ')
	}
	if border.Right > 0 {
		b = append(b, '|')
	}
	return
}

func makeContentByteArray(s string, w int, border Border) (b []byte) {
	str := []byte(s)
	var (
		leadingEmptyBytes  int
		trailingEmptyBytes int
	)

	mod, emptyBytesEach := (w-stringWidth(s))%2, (w-stringWidth(s))/2
	leadingEmptyBytes = emptyBytesEach + mod
	trailingEmptyBytes = emptyBytesEach

	if border.Left > 0 {
		b = append(b, '|')
	}

	for i := 0; i < w; i += 1 {
		if i < leadingEmptyBytes || i >= w-trailingEmptyBytes {
			b = append(b, ' ')
		} else {
			b = append(b, str[i-leadingEmptyBytes])
		}
	}

	if border.Right > 0 {
		b = append(b, '|')
	}
	return
}
