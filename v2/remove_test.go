package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var removeTests = []struct {
	ss        []int
	idx       int
	expected  []int
	succeeded bool
}{
	{
		[]int{},
		-1,
		[]int{},
		false,
	},
	{
		[]int{},
		0,
		[]int{},
		false,
	},
	{
		[]int{1, 2, 3, 4, 5},
		2,
		[]int{1, 2, 4, 5},
		true,
	},
}

func TestRemove(t *testing.T) {
	for _, test := range removeTests {

		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.succeeded, pie.Remove(&test.ss, test.idx))
			assert.Equal(t, test.ss, test.expected)
		})
	}
}
