package parser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDraw(t *testing.T) {
	primitiveCell := &Cell{
		ParentTable: nil,
		Position:    Coordinate{},
		Width:       0,
		Height:      0,
		Type:        0,
		Content:     nil,
	}
	err := primitiveCell.parse(reflect.ValueOf("test"))
	assert.Nil(t, err)
	require.NotNil(t, primitiveCell.Content)

	c := primitiveCell.Draw()

	assert.Equal(t, []byte("+----+"), c.Matrix[0])
	assert.Equal(t, []byte("|test|"), c.Matrix[1])
	assert.Equal(t, []byte("+----+"), c.Matrix[2])
}
