package parser

import "reflect"

type category struct {
	kind  reflect.Type
	value []reflect.Value
}

func (c *Cell) parseArray(desc reflect.Value) error {
	// first col will be indexes
	// second col will be values
	// values themselves can be
	// PrimitiveCell or ComplexCell
	// indexes are PrimitiveCell

	// here we divide our array to many groups
	// one is the structs that have the same type
	// will be categorized by their type in the
	// arrayOfStructsMap.
	// the rest that arent structs will be saved
	// in the arrayOfNonStructs.
	categorizedArray := []category{}
	uncategorizedArray := []reflect.Value{}
	for i := 0; i < desc.Len(); i += 1 {
		index := desc.Index(i)
		if index.Type().Kind() == reflect.Struct {
			assigned := false
			for i, c := range categorizedArray {
				if c.kind == index.Type() {
					categorizedArray[i].value = append(categorizedArray[i].value, index)
					assigned = true
					break
				}
			}
			if !assigned {
				categorizedArray = append(categorizedArray, category{
					kind:  index.Type(),
					value: []reflect.Value{index},
				})
			}
		} else {
			uncategorizedArray = append(uncategorizedArray, index)
		}
	}

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
		indexCellWidth int
		valueCellWidth int
		// sum of each rows maximum heights will be
		// the height of parent Cell
		parentHeight int
	)

	// and go through it's indexes one by one
	// for each field we add it's index number
	// as the first col and it's value as the
	// second col
	//
	// for example array {"value1", "value2", "value3"}
	// will eventually be something like this:
	//
	// +----+---------------------------+
	// | 1	|		value1				|
	// +----+---------------------------+
	// | 2	|		value2				|
	// +----+---------------------------+
	// | 3	|		value3				|
	// +----+---------------------------+
	//
	// note that values themselves can be other tables
	for catIndex, category := range categorizedArray {
		// the first col is the index Cell
		// which is a PrimitiveCell
		indexCell := &Cell{
			ParentTable: table,
			Position:    Coordinate{X: catIndex, Y: 0},
			Width:       0,
			Height:      0,
			Type:        0,
			Content:     nil,
		}
		// parsing the indexCell itself, this is where
		// recursion happens
		// in this case .parse will call .primitiveParse
		// respectively as we are passing an int to it
		err := indexCell.parsePrimitive(catIndex + 1)
		if err != nil {
			return err
		}

		// after the indexCell has been parsed
		// we add it to it's appropirate location
		// in the table's Cells field, it whould go
		// i'th row, and first(0) col
		table.Cells[catIndex] = map[int]*Cell{}
		table.Cells[catIndex][0] = indexCell

		// the second col is the value Cell
		// which can be a PrimitiveCell or
		// a ComplexCell
		valueCell := &Cell{
			ParentTable: table,
			Position:    Coordinate{X: catIndex, Y: 1},
			Width:       0,
			Height:      0,
			Type:        0,
			Content:     nil,
		}

		// parsing the valueCell itself, this is where
		// recursion happens
		// in this case we don't know the value's type
		// to .parse could call any of the *parse methods
		err = valueCell.parseArrayOfStructs(category.value)
		if err != nil {
			return err
		}

		// after the valueCell has been parsed
		// we add it to it's appropirate location
		// in the table.Cells field, it whould go
		// i'th row, and second(1) col
		table.Cells[catIndex][1] = valueCell

		// now that we have both Cells parsed
		// we have some width and height calibration to do
		//
		// height of 2 Cells in the same row should be equal
		// to the bigger one
		tallerOnesHeight := max(indexCell.Height, valueCell.Height)
		indexCell.Height = tallerOnesHeight
		valueCell.Height = tallerOnesHeight
		// index Cell in all rows should have the same width
		// so we calculate the maximum of them here and will
		// apply it later
		indexCellWidth = max(indexCellWidth, indexCell.Width)
		// value Cell in all rows should have the same width
		// so we calculate the maximum of them here and will
		// apply it later
		valueCellWidth = max(valueCellWidth, valueCell.Width)
		// parent Cell height will be the sum of all rows heights
		// which we calculated and applied with tallerOnesHeight
		parentHeight += tallerOnesHeight
	}

	for i, any := range uncategorizedArray {
		index := len(categorizedArray) + i
		// the first col is the index Cell
		// which is a PrimitiveCell
		indexCell := &Cell{
			ParentTable: table,
			Position:    Coordinate{X: i, Y: 0},
			Width:       0,
			Height:      0,
			Type:        0,
			Content:     nil,
		}
		// parsing the indexCell itself, this is where
		// recursion happens
		// in this case .parse will call .primitiveParse
		// respectively as we are passing an int to it
		err := indexCell.parsePrimitive(index + 1)
		if err != nil {
			return err
		}

		// after the indexCell has been parsed
		// we add it to it's appropirate location
		// in the table's Cells field, it whould go
		// i'th row, and first(0) col
		table.Cells[index] = map[int]*Cell{}
		table.Cells[index][0] = indexCell

		// the second col is the value Cell
		// which can be a PrimitiveCell or
		// a ComplexCell
		valueCell := &Cell{
			ParentTable: table,
			Position:    Coordinate{X: index, Y: 1},
			Width:       0,
			Height:      0,
			Type:        0,
			Content:     nil,
		}

		// parsing the valueCell itself, this is where
		// recursion happens
		// in this case we don't know the value's type
		// to .parse could call any of the *parse methods
		err = valueCell.parse(any)
		if err != nil {
			return err
		}

		// after the valueCell has been parsed
		// we add it to it's appropirate location
		// in the table.Cells field, it whould go
		// i'th row, and second(1) col
		table.Cells[index][1] = valueCell

		// now that we have both Cells parsed
		// we have some width and height calibration to do
		//
		// height of 2 Cells in the same row should be equal
		// to the bigger one
		tallerOnesHeight := max(indexCell.Height, valueCell.Height)
		indexCell.Height = tallerOnesHeight
		valueCell.Height = tallerOnesHeight
		// index Cell in all rows should have the same width
		// so we calculate the maximum of them here and will
		// apply it later
		indexCellWidth = max(indexCellWidth, indexCell.Width)
		// value Cell in all rows should have the same width
		// so we calculate the maximum of them here and will
		// apply it later
		valueCellWidth = max(valueCellWidth, valueCell.Width)
		// parent Cell height will be the sum of all rows heights
		// which we calculated and applied with tallerOnesHeight
		parentHeight += tallerOnesHeight
	}

	// we go through each row to apply the width to
	// it's first and second Cell
	for _, t := range table.Cells {
		t[0].Width = indexCellWidth
		t[1].Width = valueCellWidth
	}

	// and finally we set the parent field to calculated height and width
	// the +2 at the end is for borders
	c.Height = parentHeight + 2
	c.Width = indexCellWidth + valueCellWidth + 2

	return nil
}
