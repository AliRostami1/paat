package parser

import (
	"fmt"
	"reflect"
)

func Parse(in interface{}) (*Cell, error) {
	baseField := &Cell{}

	err := baseField.parse(reflect.ValueOf(in))

	return baseField, err
}

func (c *Cell) parse(in reflect.Value) (err error) {
	switch in.Type().Kind() {
	case reflect.Interface:
		err = c.parse(extractInterface(in))
	case reflect.Array, reflect.Slice:
		// its an iteratable type
		err = c.parserArray(in)
	case reflect.Struct:
		// its a struct type
		err = c.parserStruct(in)
	case reflect.Map:
		// its a map type
		err = c.parserMap(in)
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		// its a primitive type
		err = c.parserPrimitive(in)
	default:
		return UnknownTypeErr{
			Type: reflect.TypeOf(in),
		}
	}
	return
}

type CellType uint8

const (
	PrimitiveCell CellType = iota
	MapCell
	StructCell
	ArrayCell
	ArrayOfStructsCell
)

func (c CellType) String() string {
	switch c {
	case PrimitiveCell:
		return "primitive"
	case MapCell:
		return "map"
	case StructCell:
		return "struct"
	case ArrayCell:
		return "array"
	case ArrayOfStructsCell:
		return "array-of-structs"
	default:
		panic("this should not happen")
	}
}

type Border struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

type Location struct {
	Row int
	Col int
}

type Cell struct {
	Location
	Type    CellType
	Width   int
	Height  int
	Border  Border
	Content interface{}
}

func newCell(row, col int) *Cell {
	return &Cell{
		Location: Location{
			Row: row,
			Col: col,
		},
	}
}

func (c *Cell) BorderBoxWidth() int {
	return c.Border.Left + c.Width + c.Border.Right
}

func (c *Cell) BorderBoxHeight() int {
	return c.Border.Top + c.Height + c.Border.Bottom
}

func (c *Cell) SetWidth(w int) {
	c.Width = w

	if c.Type == MapCell || c.Type == StructCell || c.Type == ArrayCell || c.Type == ArrayOfStructsCell {
		content := c.Content.(*Table)
		content.ForEach(func(rowIndex, colIndex int) {
			currentWidths := make([]int, len(content.Cells[rowIndex]))

			for i, cell := range content.Cells[rowIndex] {
				currentWidths[i] = cell.BorderBoxWidth()
			}

			dl := distributeLength(w, currentWidths)
			content.Cells[rowIndex][colIndex].SetBorderBoxWidth(dl[colIndex])
		})
	}
}

func (c *Cell) SetBorderBoxWidth(bbw int) {
	w := bbw - c.Border.Left - c.Border.Right
	c.SetWidth(w)
}

func (c *Cell) SetHeight(h int) {
	c.Height = h

	// if c.Type == PrimitiveCell {
	// 	str := c.Content.(string)
	// }

	if c.Type == MapCell || c.Type == StructCell || c.Type == ArrayCell || c.Type == ArrayOfStructsCell {
		content := c.Content.(*Table)

		currentHeights := make([]int, len(content.Cells))

		for i, row := range content.Cells {
			currentHeights[i] = row[0].BorderBoxHeight()
		}

		dl := distributeLength(h, currentHeights)
		content.ForEach(func(rowIndex, colIndex int) {
			content.Cells[rowIndex][colIndex].SetBorderBoxHeight(dl[rowIndex])
		})
	}
}

func (c *Cell) SetBorderBoxHeight(bbh int) {
	pureHeight := bbh - c.Border.Top - c.Border.Bottom
	c.SetHeight(pureHeight)
}

type Table struct {
	Cells [][]*Cell
}

func (c *Table) ForEach(f func(rowIndex, colIndex int)) {
	for rowIndex := range c.Cells {
		for colIndex := range c.Cells[rowIndex] {
			f(rowIndex, colIndex)
		}
	}
}

func newTable(rows, cols int) *Table {
	c := make([][]*Cell, rows)
	for i := range c {
		c[i] = make([]*Cell, cols)
	}

	return &Table{
		Cells: c,
	}
}

type UnknownTypeErr struct {
	Type reflect.Type
}

func (e UnknownTypeErr) Error() string {
	return fmt.Sprintf("value of type %s is unknown to the Parser", e.Type)
}
