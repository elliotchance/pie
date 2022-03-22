package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var reduceTests = []struct {
	ss       []float64
	expected float64
	reducer  func(a, b float64) float64
}{
	{
		[]float64{1, 2, 3},
		6,
		func(a, b float64) float64 { return a + b },
	},
	{
		[]float64{1, 2, 3},
		-4,
		func(a, b float64) float64 { return a - b },
	},
	{
		[]float64{},
		0,
		func(a, b float64) float64 { return a - b },
	},
	{
		[]float64{1},
		1,
		func(a, b float64) float64 { return a - b },
	},
}

func TestReduce(t *testing.T) {
	for _, test := range reduceTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Reduce(test.ss, test.reducer))
		})
	}
}
