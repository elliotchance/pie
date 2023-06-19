package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var deleteTests = []struct {
	ss       []int
	idx      []int
	expected []int
}{
	// idx out of bounds
	{
		[]int{1, 2},
		[]int{-1},
		[]int{1, 2},
	},
	{
		[]int{1, 2},
		[]int{2},
		[]int{1, 2},
	},
	// remove from empty slice
	{
		[]int{},
		[]int{0},
		[]int{},
	},
	{
		[]int{1},
		[]int{0},
		[]int{},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]int{2},
		[]int{1, 2, 4, 5},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 3},
		[]int{1, 3, 5},
	},
	// mixed indices
	{
		[]int{1, 2, 3, 4, 5},
		[]int{1, -1, 5, 3},
		[]int{1, 3, 5},
	},
}

func TestDelete(t *testing.T) {
	for _, test := range deleteTests {

		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Delete(test.ss, test.idx...))
		})
	}
}
