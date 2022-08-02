package pie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var chunkTests = []struct {
	ss          []int
	chunkLength int
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

func TestChunk(t *testing.T) {
	for _, test := range chunkTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, Chunk(test.ss, test.chunkLength))
		})
	}
}
