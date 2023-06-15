package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var deleteTests = []struct {
	ss       []int
	idx      int
	expected []int
}{
	{
		[]int{1, 2},
		-1,
		[]int{1, 2},
	},
	{
		[]int{1, 2},
		2,
		[]int{1, 2},
	},
	{
		[]int{},
		0,
		[]int{},
	},
	{
		[]int{1},
		0,
		[]int{},
	},
	{
		[]int{1, 2, 3, 4, 5},
		2,
		[]int{1, 2, 4, 5},
	},
}

func TestDelete(t *testing.T) {
	for _, test := range deleteTests {

		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Delete(test.ss, test.idx))
		})
	}
}
