package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var shiftAndUnshiftTests = []struct {
	ss      []float64
	shifted float64
	shift   []float64
	params  []float64
	unshift []float64
}{
	{
		nil,
		0,
		nil,
		nil,
		[]float64{},
	},
	{
		nil,
		0,
		nil,
		[]float64{},
		[]float64{},
	},
	{
		nil,
		0,
		nil,
		[]float64{1.23, 2.34},
		[]float64{1.23, 2.34},
	},
	{
		[]float64{},
		0,
		nil,
		nil,
		[]float64{},
	},
	{
		[]float64{},
		0,
		nil,
		[]float64{},
		[]float64{},
	},
	{
		[]float64{},
		0,
		nil,
		[]float64{1.23, 2.34},
		[]float64{1.23, 2.34},
	},
	{
		[]float64{1.23},
		1.23,
		nil,
		[]float64{2.34},
		[]float64{2.34, 1.23},
	},
	{
		[]float64{1.23, 2.34},
		1.23,
		[]float64{2.34},
		[]float64{3.45},
		[]float64{3.45, 1.23, 2.34},
	},
	{
		[]float64{1.23, 2.34},
		1.23,
		[]float64{2.34},
		[]float64{3.45, 4.56},
		[]float64{3.45, 4.56, 1.23, 2.34},
	},
}

func TestShiftAndUnshift(t *testing.T) {
	for _, test := range shiftAndUnshiftTests {
		t.Run("", func(t *testing.T) {
			shifted, shift := pie.Shift(test.ss)
			assert.Equal(t, test.shifted, shifted)
			assert.Equal(t, test.shift, shift)
			assert.Equal(t, test.unshift, pie.Unshift(test.ss, test.params...))
		})
	}
}
