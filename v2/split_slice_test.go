package pie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var splitSliceTests = []struct {
	ss          []int
	splitLength int
	expected    [][]int
}{
	{nil, 1, [][]int{}},
	{nil, 0, [][]int{}},
	{[]int{1, 2, 3}, 4, [][]int{{1, 2, 3}}},
	{[]int{1, 2, 3}, 3, [][]int{{1, 2, 3}}},
	{[]int{1, 2, 3}, 2, [][]int{{1, 2}, {3}}},
	{[]int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
	{[]int{1, 2, 3}, 0, [][]int{}},
	{[]int{1, 2, 3}, -1, [][]int{}},
}

func TestSplitSlice(t *testing.T) {
	for _, test := range splitSliceTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, SplitSlice(test.ss, test.splitLength))
		})
	}
}
