package parser

import "reflect"

func (c *Cell) parseStruct(desc reflect.Value) error {
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

	var (
		cellWidth []int = make([]int, desc.NumField())
		rowHeight []int = make([]int, 2)
	)

	structFields := reflect.VisibleFields(desc.Type())

	table.Cells[0] = map[int]*Cell{}
	table.Cells[1] = map[int]*Cell{}

	for index, structField := range structFields {
		headerCell := &Cell{
			ParentTable: table,
			Position: Coordinate{
				X: 0,
				Y: index,
			},
			Width:   0,
			Height:  0,
			Type:    0,
			Content: nil,
		}
		headerCell.parsePrimitive(structField.Name)

		table.Cells[0][index] = headerCell

		valueCell := &Cell{
			ParentTable: table,
			Position: Coordinate{
				X: 1,
				Y: index,
			},
			Width:   0,
			Height:  0,
			Type:    0,
			Content: nil,
		}
		valueCell.parse(desc.FieldByName(structField.Name))
		table.Cells[1][index] = valueCell

		cellWidth[index] = max(headerCell.Width, valueCell.Width)
		rowHeight[0] = max(rowHeight[0], headerCell.Height)
		rowHeight[1] = max(rowHeight[1], valueCell.Height)
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
