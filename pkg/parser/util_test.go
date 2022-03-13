package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeLength(t *testing.T) {
	s := distributeLength(100, []int{10, 20, 20})
	assert.Equal(t, []int{20, 40, 40}, s)
	// s = distributeLength(101, 3)
	// assert.Equal(t, []int{34, 34, 33}, s)
	// s = distributeLength(102, 3)
	// assert.Equal(t, []int{34, 34, 34}, s)
}
