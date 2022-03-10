package parser

import (
	"fmt"
	"reflect"
)

func Parse(in interface{}) (*Cell, error) {
	baseField := &Cell{
		ParentTable: nil,
		Position:    Coordinate{X: 0, Y: 0},
		Width:       0,
		Height:      0,
		Type:        0,
		Content:     nil,
	}

	err := baseField.parse(reflect.ValueOf(in))

	return baseField, err
}

func (c *Cell) parse(in reflect.Value) (err error) {
	switch in.Type().Kind() {
	case reflect.Array, reflect.Slice:
		// its an iteratable type
		err = c.parseArray(in)
	case reflect.Struct:
		// its a struct type
		err = c.parseStruct(in)
	case reflect.Map:
		// its a map type
		err = c.parseMap(in)
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		// its a primitive type
		err = c.parsePrimitive(in)
	default:
		return UnknownTypeErr{
			Type: reflect.TypeOf(in),
		}
	}
	return
}

type CellType int8

const (
	PrimitiveCell CellType = iota
	ComplexCell
)

func (c CellType) String() string {
	switch c {
	case PrimitiveCell:
		return "primitive"
	case ComplexCell:
		return "table"
	default:
		panic("this should not happen")
	}
}

type Coordinate struct {
	X int
	Y int
}

type Cell struct {
	ParentTable *Table
	Position    Coordinate
	Padding     int
	Width       int
	Height      int
	Type        CellType
	Content     interface{}
}

type Table struct {
	Location Coordinate
	Cells    map[int]map[int]*Cell
}

type UnknownTypeErr struct {
	Type reflect.Type
}

func (e UnknownTypeErr) Error() string {
	return fmt.Sprintf("value of type %s is unknown to the Parser", e.Type)
}

type WrongSetArg struct {
	CellType
	Arg reflect.Type
}

func (e WrongSetArg) Error() string {
	return fmt.Sprintf("argument of type %s is not assignable to field of type %s", e.Arg, e.CellType)
}

func max(firstNum, secondNum int) int {
	if firstNum > secondNum {
		return firstNum
	}
	return secondNum
}
