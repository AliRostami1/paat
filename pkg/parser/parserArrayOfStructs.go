package parser

import "reflect"

func (c *Cell) parseArrayOfStructs(desc []reflect.Value) error {
	// first row will be keys
	// second row forward will be value
	// value itself can be complex or primitive

	// it's an array so the type of the parent Cell whould be
	// ComplexCell which indicates there is a Table in it's
	// contetnt field
	c.Type = ComplexCell
	// as this is a ComplexCell, we make a table with location
	// of the parent Cell's current cursor and add it to it's
	// content field
	table := &Table{
		Location: Coordinate{X: c.Position.X, Y: c.Position.Y},
		Cells:    map[int]map[int]*Cell{},
	}
	c.Content = table

	structFields := reflect.VisibleFields(desc[0].Type())

	var (
		cellWidth []int = make([]int, len(structFields)+1)
		rowHeight []int = make([]int, len(desc)+1)
	)

	hashCell := &Cell{
		ParentTable: table,
		Position: Coordinate{
			X: 0,
			Y: 0,
		},
		Width:   0,
		Height:  0,
		Type:    0,
		Content: nil,
	}
	hashCell.parsePrimitive("#")

	table.Cells[0] = map[int]*Cell{}
	table.Cells[0][0] = hashCell

	cellWidth[0] = hashCell.Width
	rowHeight[0] = hashCell.Height

	for index, structField := range structFields {
		headerCell := &Cell{
			ParentTable: table,
			Position: Coordinate{
				X: 0,
				Y: index + 1,
			},
			Width:   0,
			Height:  0,
			Type:    0,
			Content: nil,
		}
		headerCell.parsePrimitive(structField.Name)

		table.Cells[0][index+1] = headerCell

		cellWidth[index+1] = headerCell.Width
		rowHeight[0] = max(rowHeight[0], hashCell.Height)
	}

	for rowIndex, row := range desc {
		indexCell := &Cell{
			ParentTable: table,
			Position: Coordinate{
				X: rowIndex + 1,
				Y: 0,
			},
			Width:   0,
			Height:  0,
			Type:    0,
			Content: nil,
		}
		indexCell.parsePrimitive(rowIndex + 1)

		table.Cells[rowIndex+1] = map[int]*Cell{}
		table.Cells[rowIndex+1][0] = indexCell

		cellWidth[0] = max(cellWidth[0], indexCell.Width)
		rowHeight[rowIndex+1] = indexCell.Height

		for colIndex, structField := range structFields {
			headerCell := &Cell{
				ParentTable: table,
				Position: Coordinate{
					X: rowIndex + 1,
					Y: colIndex + 1,
				},
				Width:   0,
				Height:  0,
				Type:    0,
				Content: nil,
			}
			headerCell.parse(row.FieldByName(structField.Name))

			table.Cells[rowIndex+1][colIndex+1] = headerCell

			cellWidth[colIndex+1] = max(cellWidth[colIndex+1], indexCell.Width)
			rowHeight[rowIndex+1] = max(rowHeight[rowIndex+1], headerCell.Height)
		}
	}

	var (
		// sum of each rows maximum heights will be
		// the height of parent Cell
		parentHeight int
		parentWidth  int
	)
	for rowIndex, row := range table.Cells {
		parentHeight += rowHeight[rowIndex]
		for colIndex, cell := range row {
			if rowIndex == 0 {
				parentWidth += cellWidth[colIndex]
			}

			cell.Height = rowHeight[rowIndex]
			cell.Width = cellWidth[colIndex]
		}
	}

	// the +2 at the end is for borders
	c.Height = parentHeight + 2
	c.Width = parentWidth + 2

	return nil
}
