package parser

type Canvas struct {
	Matrix [][]byte
}

func Drawer(cell *Cell) {
	// build the matrix and initilize it
	var runeMatrix [][]byte = make([][]byte, cell.Height)
	for i := range runeMatrix {
		runeMatrix[i] = make([]byte, cell.Width)
	}

	canvas := &Canvas{
		Matrix: runeMatrix,
	}

	canvas.drawer(cell, 0, 0)
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
