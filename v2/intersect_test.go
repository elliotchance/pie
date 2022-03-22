package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var intersectTests = []struct {
	ss       []float64
	params   [][]float64
	expected []float64
}{
	{
		nil,
		nil,
		nil,
	},
	{
		[]float64{1.2, 3.2},
		nil,
		nil,
	},
	{
		nil,
		[][]float64{{1.2, 3.2, 5.5}, {5.5, 1.2}},
		nil,
	},
	{
		[]float64{1.2, 3.2},
		[][]float64{{1.2}, {3.2}},
		nil,
	},
	{
		[]float64{1.2, 3.2},
		[][]float64{{1.2}},
		[]float64{1.2},
	},
	{
		[]float64{1.2, 3.2},
		[][]float64{{1.2, 3.2, 5.5}, {5.5, 1.2}},
		[]float64{1.2},
	},
}

func TestIntersect(t *testing.T) {
	for _, test := range intersectTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Intersect(test.ss, test.params...))
		})
	}
}
