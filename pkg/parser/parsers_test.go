package parser

import (
	"reflect"
	"testing"

	"github.com/AliRostami1/tabler/testdata"
	"github.com/stretchr/testify/assert"
)

func TestParserPrimitive(t *testing.T) {
	c := newCell(0, 0)
	err := c.parserPrimitive("#")
	assert.Nil(t, err)
	assert.Equal(t, PrimitiveCell, c.Type)
	assert.Equal(t, 1, c.Width)
	assert.Equal(t, 1, c.Height)
	assert.Equal(t, Border{1, 1, 1, 1}, c.Border)
	assert.Equal(t, 3, c.BorderBoxWidth())
	assert.Equal(t, 3, c.BorderBoxHeight())
	assert.Equal(t, Location{0, 0}, c.Location)

	content, ok := c.Content.(string)
	assert.True(t, ok)
	assert.Equal(t, "#", content)
}

func TestParserMap(t *testing.T) {
	c := newCell(0, 0)
	err := c.parserMap(reflect.ValueOf(testdata.MapOfIntTestData))
	assert.Nil(t, err)
	assert.Equal(t, MapCell, c.Type)
	assert.Equal(t, 12, c.Width)
	assert.Equal(t, 21, c.Height)
	assert.Equal(t, Border{0, 0, 0, 0}, c.Border)
	assert.Equal(t, 12, c.BorderBoxWidth())
	assert.Equal(t, 21, c.BorderBoxHeight())
	assert.Equal(t, Location{0, 0}, c.Location)

	content, ok := c.Content.(*Table)
	assert.True(t, ok)
	content.ForEach(func(rowIndex, colIndex int) {
		c := content.Cells[rowIndex][colIndex]
		t.Logf("row:%d, col:%d, cell:%+#v", rowIndex, colIndex, *c)

		testBorder(t, c, rowIndex, colIndex)

		if colIndex == 0 {
			assert.Equal(t, 7, c.Width)

		} else {
			assert.Equal(t, 2, c.Width)
		}
	})
}

func TestParserStruct(t *testing.T) {
	c := newCell(0, 0)
	err := c.parserStruct(reflect.ValueOf(testdata.StructTestData))
	assert.Nil(t, err)
	assert.Equal(t, StructCell, c.Type)
	assert.Equal(t, 27, c.Width)
	assert.Equal(t, 5, c.Height)
	assert.Equal(t, Border{0, 0, 0, 0}, c.Border)
	assert.Equal(t, 27, c.BorderBoxWidth())
	assert.Equal(t, 5, c.BorderBoxHeight())
	assert.Equal(t, Location{0, 0}, c.Location)

	content, ok := c.Content.(*Table)
	assert.True(t, ok)
	content.ForEach(func(rowIndex, colIndex int) {
		c := content.Cells[rowIndex][colIndex]
		t.Logf("row:%d, col:%d, cell:%+#v", rowIndex, colIndex, *c)
		testBorder(t, c, rowIndex, colIndex)

		if colIndex == 0 {
			assert.Equal(t, 11, c.Width)
			assert.Equal(t, 1, c.Height)
		} else {
			assert.Equal(t, 6, c.Width)
			assert.Equal(t, 1, c.Height)
		}
	})
}

func TestParserArrayOfStructs(t *testing.T) {
	c := newCell(0, 0)
	err := c.parserArrayOfStructs(testdata.ArrayOfStructsDataReflect)
	assert.Nil(t, err)
	assert.Equal(t, ArrayOfStructsCell, c.Type)
	assert.Equal(t, 31, c.Width)
	assert.Equal(t, 23, c.Height)
	assert.Equal(t, Border{0, 0, 0, 0}, c.Border)
	assert.Equal(t, 31, c.BorderBoxWidth())
	assert.Equal(t, 23, c.BorderBoxHeight())
	assert.Equal(t, Location{0, 0}, c.Location)

	content, ok := c.Content.(*Table)
	assert.True(t, ok)
	content.ForEach(func(rowIndex, colIndex int) {
		c := content.Cells[rowIndex][colIndex]
		t.Logf("row:%d, col:%d, cell:%+#v", rowIndex, colIndex, *c)
		testBorder(t, c, rowIndex, colIndex)

		switch colIndex {
		case 0:
			assert.Equal(t, 2, c.Width)
			assert.Equal(t, 1, c.Height)
		case 1:
			assert.Equal(t, 12, c.Width)
			assert.Equal(t, 1, c.Height)
		default:
			assert.Equal(t, 6, c.Width)
			assert.Equal(t, 1, c.Height)
		}
	})
}

func TestParserArray(t *testing.T) {
	c := newCell(0, 0)
	err := c.parserArray(reflect.ValueOf(testdata.ArrayTestData))
	assert.Nil(t, err)
	assert.Equal(t, ArrayCell, c.Type)
	assert.Equal(t, 35, c.Width)
	assert.Equal(t, 43, c.Height)
	assert.Equal(t, Border{0, 0, 0, 0}, c.Border)
	assert.Equal(t, 35, c.BorderBoxWidth())
	assert.Equal(t, 43, c.BorderBoxHeight())
	assert.Equal(t, Location{0, 0}, c.Location)

	content, ok := c.Content.(*Table)
	assert.True(t, ok)
	content.ForEach(func(rowIndex, colIndex int) {
		c := content.Cells[rowIndex][colIndex]
		t.Logf("row:%d, col:%d, cell:%+#v", rowIndex, colIndex, *c)
		// testBorder(t, c, rowIndex, colIndex)

	})
}

func testBorder(t *testing.T, c *Cell, rowIndex, colIndex int) {
	if rowIndex == 0 {
		assert.Equal(t, 1, c.Border.Top)
	} else {
		assert.Equal(t, 0, c.Border.Top)
	}
	if colIndex == 0 {
		assert.Equal(t, 1, c.Border.Left)
	} else {
		assert.Equal(t, 0, c.Border.Left)
	}
	assert.Equal(t, 1, c.Border.Bottom)
	assert.Equal(t, 1, c.Border.Right)
}
