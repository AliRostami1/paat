package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeLength(t *testing.T) {
	s := distributeLength(100, []int{10, 20, 20})
	assert.Equal(t, []int{20, 40, 40}, s)
}

func TestCalcBorder(t *testing.T) {
	b := Border{1, 1, 1, 1}
	rb := calcBorder(b, 0, 0)
	assert.Equal(t, rb, Border{1, 1, 1, 1})

	b = Border{0, 0, 1, 1}
	rb = calcBorder(b, 0, 0)
	assert.Equal(t, rb, Border{0, 0, 1, 1})

	b = Border{1, 1, 1, 1}
	rb = calcBorder(b, 2, 0)
	assert.Equal(t, rb, Border{1, 0, 1, 1})

	b = Border{1, 1, 1, 1}
	rb = calcBorder(b, 0, 2)
	assert.Equal(t, rb, Border{0, 1, 1, 1})

	b = Border{1, 1, 1, 1}
	rb = calcBorder(b, 2, 2)
	assert.Equal(t, rb, Border{0, 0, 1, 1})
}
