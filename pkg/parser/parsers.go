package parser

import (
	"fmt"
	"reflect"
)

func (c *Cell) parserPrimitive(in interface{}) error {
	// single cell representing the primitive value

	// it's a primitive so the type of the parent Cell whould be
	// PrimitiveCell which indicates there is a string in it's
	// contetnt field
	c.Type = PrimitiveCell
	// as this is a PrimitiveCell, we extract the string
	// reperesentation of it's value and assign it to it's
	// content field
	strRep := fmt.Sprint(in)
	c.Content = strRep
	// width of the cell will be the number of Unicode code points
	// in the string reperesntation of it
	// the +2 at the end is for borders on left and right sides
	c.Width = stringWidth(strRep)
	// height of the cell will be number of \n(newline characters)
	// in the string reperesntation of it
	// the +2 at the end is for borders on top and bottom sides
	c.Height = stringHeight(strRep)
	return nil
}
func (c *Cell) parserMap(in reflect.Value) error {
	// it's an array so the type of the parent Cell whould be
	// ComplexCell which indicates there is a Table in it's
	// contetnt field
	c.Type = MapCell
	// as this is a ComplexCell, we make a table with location
	// of the parent Cell's current cursor and add it to it's
	// content field
	table := newTable(len(in.MapKeys()), 2)
	c.Content = table

	var (
		rowWidth   []int = make([]int, 2)
		cellHeight []int = make([]int, len(in.MapKeys()))
		// sum of each rows maximum heights will be
		// the height of parent Cell
		parentHeight int
		parentWidth  int
	)

	mapKeys := in.MapKeys()

	for index, mapKey := range mapKeys {
		headerCell := c.newCell(index, 0)
		headerCell.parserPrimitive(mapKey)
		table.Cells[index][0] = headerCell

		valueCell := c.newCell(index, 1)
		valueCell.parse(in.MapIndex(mapKey))
		table.Cells[index][1] = valueCell

		cellHeight[index] = max(headerCell.BorderBoxHeight(), valueCell.BorderBoxHeight())
		rowWidth[0] = max(rowWidth[0], headerCell.BorderBoxWidth())
		rowWidth[1] = max(rowWidth[1], valueCell.BorderBoxWidth())

		parentHeight += max(headerCell.BorderBoxHeight(), valueCell.BorderBoxHeight())
		bbw := headerCell.BorderBoxWidth() + valueCell.BorderBoxWidth()
		parentWidth = max(parentWidth, bbw)
	}

	for rowIndex, row := range table.Cells {
		for colIndex, cell := range row {
			cell.SetBorderBoxHeight(cellHeight[rowIndex])
			cell.SetBorderBoxWidth(rowWidth[colIndex])
		}
	}

	c.SetBorderBoxWidth(parentWidth)
	c.SetBorderBoxHeight(parentHeight)

	return nil
}
func (c *Cell) parserStruct(in reflect.Value) error {

	in = extractInterface(in)

	// it's an array so the type of the parent Cell whould be
	// ComplexCell which indicates there is a Table in it's
	// contetnt field
	c.Type = StructCell
	// as this is a ComplexCell, we make a table with location
	// of the parent Cell's current cursor and add it to it's
	// content field
	table := newTable(2, in.NumField())
	c.Content = table

	var (
		cellWidth []int = make([]int, in.NumField())
		rowHeight []int = make([]int, 2)

		// sum of each rows maximum heights will be
		// the height of parent Cell
		parentHeight int
		parentWidth  int
	)

	structFields := reflect.VisibleFields(in.Type())

	for index, structField := range structFields {
		headerCell := c.newCell(0, index)
		headerCell.parserPrimitive(structField.Name)
		table.Cells[0][index] = headerCell

		valueCell := c.newCell(1, index)
		valueCell.parse(in.FieldByName(structField.Name))
		table.Cells[1][index] = valueCell

		cellWidth[index] = max(headerCell.BorderBoxWidth(), valueCell.BorderBoxWidth())
		rowHeight[0] = max(rowHeight[0], headerCell.BorderBoxHeight())
		rowHeight[1] = max(rowHeight[1], valueCell.BorderBoxHeight())

		parentHeight = max(parentHeight, headerCell.BorderBoxHeight()+valueCell.BorderBoxHeight())
		parentWidth += max(headerCell.BorderBoxWidth(), valueCell.BorderBoxWidth())
	}

	for rowIndex, row := range table.Cells {
		for colIndex, cell := range row {
			cell.SetBorderBoxHeight(rowHeight[rowIndex])
			cell.SetBorderBoxWidth(cellWidth[colIndex])
		}
	}

	c.SetBorderBoxWidth(parentWidth)
	c.SetBorderBoxHeight(parentHeight)

	return nil
}

func (c *Cell) parserArray(in reflect.Value) error {
	in = extractInterface(in)

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
	categorizedArray := [][]reflect.Value{}
	uncategorizedArray := []reflect.Value{}

	in = extractInterface(in)

	for i := 0; i < in.Len(); i += 1 {
		index := extractInterface(in.Index(i))

		kind := index.Type().Kind()
		if kind == reflect.Struct {
			assigned := false
			for i, c := range categorizedArray {
				if c[0].Type() == index.Type() {
					categorizedArray[i] = append(categorizedArray[i], index)
					assigned = true
					break
				}
			}
			if !assigned {
				categorizedArray = append(categorizedArray, []reflect.Value{index})
			}
		} else {
			uncategorizedArray = append(uncategorizedArray, index)
		}
	}

	// it's an array so the type of the parent Cell whould be
	// ComplexCell which indicates there is a Table in it's
	// contetnt field
	c.Type = ArrayCell
	// as this is a ComplexCell, we make a table with location
	// of the parent Cell's current cursor and add it to it's
	// content field
	table := newTable(len(categorizedArray)+len(uncategorizedArray), 2)
	c.Content = table

	var (
		indexCellBBW int
		valueCellBBW int
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
		indexCell := c.newCell(catIndex, 0)

		// parsing the indexCell itself, this is where
		// recursion happens
		// in this case .parse will call .primitiveParse
		// respectively as we are passing an int to it
		err := indexCell.parserPrimitive(catIndex + 1)
		if err != nil {
			return err
		}

		// after the indexCell has been parsed
		// we add it to it's appropirate location
		// in the table's Cells field, it whould go
		// i'th row, and first(0) col
		table.Cells[catIndex][0] = indexCell

		// the second col is the value Cell
		// which can be a PrimitiveCell or
		// a ComplexCell
		valueCell := c.newCell(catIndex, 1)
		// parsing the valueCell itself, this is where
		// recursion happens
		// in this case we don't know the value's type
		// to .parse could call any of the *parse methods
		err = valueCell.parserArrayOfStructs(category)
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
		tallerOnesHeight := max(indexCell.BorderBoxHeight(), valueCell.BorderBoxHeight())
		indexCell.SetBorderBoxHeight(tallerOnesHeight)
		valueCell.SetBorderBoxHeight(tallerOnesHeight)
		// index Cell in all rows should have the same width
		// so we calculate the maximum of them here and will
		// apply it later
		indexCellBBW = max(indexCellBBW, indexCell.BorderBoxWidth())
		// value Cell in all rows should have the same width
		// so we calculate the maximum of them here and will
		// apply it later
		valueCellBBW = max(valueCellBBW, valueCell.BorderBoxWidth())
		// parent Cell height will be the sum of all rows heights
		// which we calculated and applied with tallerOnesHeight
		parentHeight += tallerOnesHeight
	}

	for i, any := range uncategorizedArray {
		index := len(categorizedArray) + i
		// the first col is the index Cell
		// which is a PrimitiveCell
		indexCell := c.newCell(index, 0)
		// parsing the indexCell itself, this is where
		// recursion happens
		// in this case .parse will call .primitiveParse
		// respectively as we are passing an int to it
		err := indexCell.parserPrimitive(index + 1)
		if err != nil {
			return err
		}

		// after the indexCell has been parsed
		// we add it to it's appropirate location
		// in the table's Cells field, it whould go
		// i'th row, and first(0) col
		table.Cells[index][0] = indexCell

		// the second col is the value Cell
		// which can be a PrimitiveCell or
		// a ComplexCell
		valueCell := c.newCell(index, 1)

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
		tallerOnesHeight := max(indexCell.BorderBoxHeight(), valueCell.BorderBoxHeight())
		indexCell.SetBorderBoxHeight(tallerOnesHeight)
		valueCell.SetBorderBoxHeight(tallerOnesHeight)
		// index Cell in all rows should have the same width
		// so we calculate the maximum of them here and will
		// apply it later
		indexCellBBW = max(indexCellBBW, indexCell.BorderBoxWidth())
		// value Cell in all rows should have the same width
		// so we calculate the maximum of them here and will
		// apply it later
		valueCellBBW = max(valueCellBBW, valueCell.BorderBoxWidth())

		// parent Cell height will be the sum of all rows heights
		// which we calculated and applied with tallerOnesHeight
		parentHeight += tallerOnesHeight
	}

	// we go through each row to apply the width to
	// it's first and second Cell
	for row := range table.Cells {
		table.Cells[row][0].SetBorderBoxWidth(indexCellBBW)
		table.Cells[row][1].SetBorderBoxWidth(valueCellBBW)
	}

	// and finally we set the parent field to calculated height and width
	// the +2 at the end is for borders
	c.SetBorderBoxHeight(parentHeight)
	c.SetBorderBoxWidth(valueCellBBW + indexCellBBW)

	return nil
}

func (c *Cell) parserArrayOfStructs(in []reflect.Value) error {
	for i := range in {
		in[i] = extractInterface(in[i])
	}
	// first row will be keys
	// second row forward will be value
	// value itself can be complex or primitive

	// it's an array so the type of the parent Cell whould be
	// ComplexCell which indicates there is a Table in it's
	// contetnt field
	c.Type = ArrayOfStructsCell
	// as this is a ComplexCell, we make a table with location
	// of the parent Cell's current cursor and add it to it's
	// content field
	table := newTable(len(in)+1, in[0].NumField()+1)
	c.Content = table

	structFields := reflect.VisibleFields(in[0].Type())

	var (
		cellWidth []int = make([]int, len(structFields)+1)
		rowHeight []int = make([]int, len(in)+1)
	)

	hashCell := c.newCell(0, 0)
	hashCell.parserPrimitive("#")

	table.Cells[0][0] = hashCell

	cellWidth[0] = hashCell.BorderBoxWidth()
	rowHeight[0] = hashCell.BorderBoxHeight()

	for index, structField := range structFields {
		headerCell := c.newCell(0, index+1)
		headerCell.parserPrimitive(structField.Name)

		table.Cells[0][index+1] = headerCell

		cellWidth[index+1] = headerCell.BorderBoxWidth()
		rowHeight[0] = max(rowHeight[0], hashCell.BorderBoxHeight())
	}

	for rowIndex, row := range in {
		indexCell := c.newCell(rowIndex+1, 0)
		indexCell.parserPrimitive(rowIndex + 1)

		table.Cells[rowIndex+1][0] = indexCell

		cellWidth[0] = max(cellWidth[0], indexCell.BorderBoxWidth())
		rowHeight[rowIndex+1] = indexCell.BorderBoxHeight()

		for colIndex, structField := range structFields {
			valueCell := c.newCell(rowIndex+1, colIndex+1)
			valueCell.parse(row.FieldByName(structField.Name))

			table.Cells[rowIndex+1][colIndex+1] = valueCell

			cellWidth[colIndex+1] = max(cellWidth[colIndex+1], valueCell.BorderBoxWidth())
			rowHeight[rowIndex+1] = max(rowHeight[rowIndex+1], valueCell.BorderBoxHeight())
		}
	}

	var (
		// sum of each rows maximum heights will be
		// the height of parent Cell
		parentHeight int
		parentWidth  int
	)
	for rowIndex, row := range table.Cells {
		for colIndex, cell := range row {
			cell.SetBorderBoxHeight(rowHeight[rowIndex])
			parentHeight += rowHeight[rowIndex]
			cell.SetBorderBoxWidth(cellWidth[colIndex])
			if rowIndex == 0 {
				parentWidth += cellWidth[colIndex]
			}

		}
	}

	// the +2 at the end is for borders
	c.SetBorderBoxHeight(parentHeight)
	c.SetBorderBoxWidth(parentWidth)

	return nil
}
