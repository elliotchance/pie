package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var findFirstUsingTests = []struct {
	ss         []float64
	expression func(value float64) bool
	expected   int
}{
	{
		nil,
		func(value float64) bool { return value == 1.5 },
		-1,
	},
	{
		[]float64{},
		func(value float64) bool { return value == 0.1 },
		-1,
	},
	{
		[]float64{0.0, 1.5, 3.2},
		func(value float64) bool { return value == 9.99 },
		-1,
	},
	{
		[]float64{5.4, 6.98, 4.987, 33.04},
		func(value float64) bool { return value == 33.04 },
		3,
	},
	{
		[]float64{9.0, 0.11, 150.44, 33.04},
		func(value float64) bool { return value == 0.11 },
		1,
	},
}

func TestFindFirstUsing(t *testing.T) {
	for _, test := range findFirstUsingTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.FindFirstUsing(test.ss, test.expression))
		})
	}
}
