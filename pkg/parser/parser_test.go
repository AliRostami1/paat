package parser

import (
	"fmt"
	"reflect"
	"testing"

	testdata "github.com/AliRostami1/tabler/testdata/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestInterface struct {
	firstField string
	seconField int // i know its second but for simpilicity in testing Cell.Width we are going with secon
	thirdField float64
}

var testDatas = [...]TestInterface{
	{
		firstField: "field1",
		seconField: 1,
		thirdField: 2.1,
	},
	{
		firstField: "field2",
		seconField: 2,
		thirdField: 2.2,
	},
	{
		firstField: "field3",
		seconField: 3,
		thirdField: 2.3,
	},
	{
		firstField: "field4",
		seconField: 4,
		thirdField: 2.4,
	},
	{
		firstField: "field5",
		seconField: 5,
		thirdField: 2.5,
	},
	{
		firstField: "field6",
		seconField: 6,
		thirdField: 2.6,
	},
}

var testDataReflects = [...]reflect.Value{
	reflect.ValueOf(testDatas[0]),
	reflect.ValueOf(testDatas[1]),
	reflect.ValueOf(testDatas[2]),
	reflect.ValueOf(testDatas[3]),
	reflect.ValueOf(testDatas[4]),
	reflect.ValueOf(testDatas[5]),
}

func TestParser(t *testing.T) {
	cell, err := Parse(testdata.ComplexTestData)
	assert.Nil(t, err)
	require.NotNil(t, cell)

	assert.Equal(t, ComplexCell, cell.Type)
	mainTable, ok := cell.Content.(*Table)
	assert.True(t, ok)
	assert.NotNil(t, mainTable)

	assert.Equal(t, 2, len(mainTable.Cells))
	assert.Equal(t, 7, len(mainTable.Cells[0]))
	assert.Equal(t, 7, len(mainTable.Cells[1]))

	logTable(t, mainTable, 0)
	t.FailNow()
}

func TestCellParser(t *testing.T) {
	primitiveCell := &Cell{
		ParentTable: nil,
		Position:    Coordinate{},
		Width:       0,
		Height:      0,
		Type:        0,
		Content:     nil,
	}
	err := primitiveCell.parse(reflect.ValueOf("1"))
	assert.Nil(t, err)
	require.NotNil(t, primitiveCell.Content)

	strContent, ok := primitiveCell.Content.(string)
	assert.True(t, ok)
	assert.Equal(t, "1", strContent)

	assert.Equal(t, 3, primitiveCell.Width)
	assert.Equal(t, 3, primitiveCell.Height)

	// ======================================

	complexCell := &Cell{
		ParentTable: nil,
		Position:    Coordinate{},
		Width:       0,
		Height:      0,
		Type:        0,
		Content:     nil,
	}

	err = complexCell.parse(testDataReflects[0])
	assert.Nil(t, err)
}

func TestPrimitiveParser(t *testing.T) {
	for _, testData := range testDatas {
		primitiveCell1 := &Cell{
			ParentTable: nil,
			Position:    Coordinate{},
			Width:       0,
			Height:      0,
			Type:        0,
			Content:     nil,
		}
		err := primitiveCell1.parsePrimitive(testData.firstField)
		assert.Nil(t, err)
		require.NotNil(t, primitiveCell1.Content)

		strContent, ok := primitiveCell1.Content.(string)
		assert.True(t, ok)
		assert.Equal(t, testData.firstField, strContent)

		assert.Equal(t, PrimitiveCell, primitiveCell1.Type)
		assert.Equal(t, 8, primitiveCell1.Width)
		assert.Equal(t, 3, primitiveCell1.Height)

		primitiveCell2 := &Cell{
			ParentTable: nil,
			Position:    Coordinate{},
			Width:       0,
			Height:      0,
			Type:        0,
			Content:     nil,
		}
		err = primitiveCell2.parsePrimitive(testData.seconField)
		assert.Nil(t, err)
		require.NotNil(t, primitiveCell2.Content)

		strContent, ok = primitiveCell2.Content.(string)
		assert.True(t, ok)
		assert.Equal(t, fmt.Sprint(testData.seconField), strContent)

		assert.Equal(t, PrimitiveCell, primitiveCell2.Type)
		assert.Equal(t, 3, primitiveCell2.Width)
		assert.Equal(t, 3, primitiveCell2.Height)

		primitiveCell3 := &Cell{
			ParentTable: nil,
			Position:    Coordinate{},
			Width:       0,
			Height:      0,
			Type:        0,
			Content:     nil,
		}
		err = primitiveCell3.parsePrimitive(testData.thirdField)
		assert.Nil(t, err)
		require.NotNil(t, primitiveCell3.Content)

		strContent, ok = primitiveCell3.Content.(string)
		assert.True(t, ok)
		assert.Equal(t, fmt.Sprint(testData.thirdField), strContent)

		assert.Equal(t, PrimitiveCell, primitiveCell3.Type)
		assert.Equal(t, 5, primitiveCell3.Width)
		assert.Equal(t, 3, primitiveCell3.Height)
	}

}

func TestStructParser(t *testing.T) {
	complexCell := &Cell{
		ParentTable: nil,
		Position:    Coordinate{},
		Width:       0,
		Height:      0,
		Type:        0,
		Content:     nil,
	}
	err := complexCell.parseStruct(testDataReflects[0])
	assert.Nil(t, err)
	require.NotNil(t, complexCell.Content)

	assert.Equal(t, ComplexCell, complexCell.Type)

	table, ok := complexCell.Content.(*Table)
	assert.True(t, ok)
	require.NotNil(t, table)

	assert.Equal(t, 2, len(table.Cells))
	assert.Equal(t, 3, len(table.Cells[0]))
	assert.Equal(t, 3, len(table.Cells[1]))

	for rowIndex, row := range table.Cells {
		for colIndex, col := range row {
			t.Logf("row:%d col:%d = %v", rowIndex, colIndex, *col)
			assert.NotNil(t, col)
			assert.Equal(t, 12, col.Width)
			assert.Equal(t, 3, col.Height)
		}
	}

	assert.Equal(t, "firstField", table.Cells[0][0].Content) // width 12
	assert.Equal(t, "seconField", table.Cells[0][1].Content) // width 12
	assert.Equal(t, "thirdField", table.Cells[0][2].Content) // width 12
	assert.Equal(t, "field1", table.Cells[1][0].Content)     // width 3
	assert.Equal(t, "1", table.Cells[1][1].Content)          // width 3
	assert.Equal(t, "2.1", table.Cells[1][2].Content)        // width 5

	assert.Equal(t, 38, complexCell.Width)
	assert.Equal(t, 8, complexCell.Height)
}

func TestArrayOfStructsParser(t *testing.T) {
	complexCell := &Cell{
		ParentTable: nil,
		Position:    Coordinate{},
		Width:       0,
		Height:      0,
		Type:        0,
		Content:     nil,
	}

	err := complexCell.parseArrayOfStructs(testDataReflects[:])
	assert.Nil(t, err)
	require.NotNil(t, complexCell.Content)

	assert.Equal(t, ComplexCell, complexCell.Type)

	table, ok := complexCell.Content.(*Table)
	assert.True(t, ok)
	require.NotNil(t, table)

	assert.Equal(t, 7, len(table.Cells))

	for rowIndex, row := range table.Cells {
		assert.Equal(t, 4, len(row))
		for colIndex, col := range row {
			t.Logf("row:%d col:%d = %v", rowIndex, colIndex, *col)
			assert.NotNil(t, col)
			if col.Position.Y != 0 {
				assert.Equal(t, 12, col.Width)
				assert.Equal(t, 3, col.Height)
			} else {
				assert.Equal(t, 3, col.Width)
				assert.Equal(t, 3, col.Height)
			}
		}
	}

	assert.Equal(t, 41, complexCell.Width)  // 3 * 12 + 1 * 3 + 2
	assert.Equal(t, 23, complexCell.Height) // 7 * 3 + 2
}

func TestArrayParser(t *testing.T) {
	complexCell := &Cell{
		ParentTable: nil,
		Position:    Coordinate{},
		Width:       0,
		Height:      0,
		Type:        0,
		Content:     nil,
	}
	err := complexCell.parseArray(reflect.ValueOf(testDatas))
	assert.Nil(t, err)
	require.NotNil(t, complexCell.Content)

	assert.Equal(t, ComplexCell, complexCell.Type)

	table, ok := complexCell.Content.(*Table)
	assert.True(t, ok)
	require.NotNil(t, table)

	assert.Equal(t, 1, len(table.Cells))    // 1 row
	assert.Equal(t, 2, len(table.Cells[0])) // 2 cols

	for rowIndex, row := range table.Cells {
		assert.Equal(t, 2, len(row))
		for colIndex, col := range row {
			t.Logf("main table: row:%d col:%d = %v", rowIndex, colIndex, *col)
			assert.NotNil(t, col)
			if colIndex != 0 {
				assert.Equal(t, 41, col.Width)
				assert.Equal(t, 23, col.Height)
			} else {
				assert.Equal(t, 3, col.Width)
				assert.Equal(t, 23, col.Height)
			}
		}
	}

	nestedTable, ok := table.Cells[0][1].Content.(*Table)
	assert.True(t, ok)
	require.NotNil(t, nestedTable)

	assert.Equal(t, 7, len(nestedTable.Cells))

	for rowIndex, row := range nestedTable.Cells {
		assert.Equal(t, 4, len(row))
		for colIndex, col := range row {
			t.Logf("nested table: row:%d col:%d = %v", rowIndex, colIndex, *col)
			assert.NotNil(t, col)
			if colIndex != 0 {
				assert.Equal(t, 12, col.Width)
				assert.Equal(t, 3, col.Height)
			} else {
				assert.Equal(t, 3, col.Width)
				assert.Equal(t, 3, col.Height)
			}
		}
	}

	assert.Equal(t, 46, complexCell.Width)
	assert.Equal(t, 25, complexCell.Height)
}

func TestMax(t *testing.T) {
	assert.Equal(t, 30, max(30, 20))
	assert.Equal(t, 30, max(20, 30))
	assert.Equal(t, 1000, max(10, 1000))
	assert.Equal(t, -1, max(-1, -100))
	assert.Equal(t, -1, max(-1, -1))
}

func logTable(t *testing.T, table *Table, nestingIndex int) {
	for rowIndex := 0; rowIndex < len(table.Cells); rowIndex += 1 {
		for colIndex := 0; colIndex < len(table.Cells[rowIndex]); colIndex += 1 {
			cell := table.Cells[rowIndex][colIndex]
			t.Logf("nestingIndex:%d row:%d col:%d = {width: %d, height: %d, Content: %v}", nestingIndex, rowIndex, colIndex, cell.Width, cell.Height, cell.Content)
			if cell.Type == ComplexCell {
				nestedTable, ok := cell.Content.(*Table)
				assert.True(t, ok)
				logTable(t, nestedTable, nestingIndex+1)
			}
		}
	}
}
