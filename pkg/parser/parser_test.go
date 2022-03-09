package parser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestInterface struct {
	firstField  string
	secondField int
	thirdField  float64
}

var testDatas = [...]TestInterface{
	{
		firstField:  "1",
		secondField: 1,
		thirdField:  2.1,
	},
	{
		firstField:  "2",
		secondField: 2,
		thirdField:  2.2,
	},
	{
		firstField:  "3",
		secondField: 3,
		thirdField:  2.3,
	},
	{
		firstField:  "4",
		secondField: 4,
		thirdField:  2.4,
	},
	{
		firstField:  "5",
		secondField: 5,
		thirdField:  2.5,
	},
	{
		firstField:  "6",
		secondField: 6,
		thirdField:  2.6,
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

	assert.Equal(t, 1, primitiveCell.Width)
	assert.Equal(t, 1, primitiveCell.Height)

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
		assert.Equal(t, 1, primitiveCell1.Width)
		assert.Equal(t, 1, primitiveCell1.Height)

		primitiveCell2 := &Cell{
			ParentTable: nil,
			Position:    Coordinate{},
			Width:       0,
			Height:      0,
			Type:        0,
			Content:     nil,
		}
		err = primitiveCell2.parsePrimitive(testData.secondField)
		assert.Nil(t, err)
		require.NotNil(t, primitiveCell2.Content)

		strContent, ok = primitiveCell2.Content.(string)
		assert.True(t, ok)
		assert.Equal(t, fmt.Sprint(testData.secondField), strContent)

		assert.Equal(t, PrimitiveCell, primitiveCell2.Type)
		assert.Equal(t, 1, primitiveCell2.Width)
		assert.Equal(t, 1, primitiveCell2.Height)

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
		assert.Equal(t, 3, primitiveCell3.Width)
		assert.Equal(t, 1, primitiveCell3.Height)
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
		}
	}

	assert.Equal(t, "firstField", table.Cells[0][0].Content)
	assert.Equal(t, "secondField", table.Cells[0][1].Content)
	assert.Equal(t, "thirdField", table.Cells[0][2].Content)
	assert.Equal(t, "1", table.Cells[1][0].Content)
	assert.Equal(t, "1", table.Cells[1][1].Content)
	assert.Equal(t, "2.1", table.Cells[1][2].Content)
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
		}
	}
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

	assert.Equal(t, 2, len(table.Cells[0]))

	assert.Equal(t, 1, len(table.Cells))

	for rowIndex, row := range table.Cells {
		assert.Equal(t, 2, len(row))
		for colIndex, col := range row {
			t.Logf("main table: row:%d col:%d = %v", rowIndex, colIndex, *col)
			assert.NotNil(t, col)
		}
	}

	nestedTable, ok := table.Cells[0][1].Content.(*Table)
	assert.True(t, ok)
	require.NotNil(t, nestedTable)

	assert.Equal(t, 6, len(nestedTable.Cells))

	for rowIndex, row := range nestedTable.Cells {
		assert.Equal(t, 4, len(row))
		for colIndex, col := range row {
			t.Logf("nested table: row:%d col:%d = %v", rowIndex, colIndex, *col)
			assert.NotNil(t, col)
		}
	}
}
