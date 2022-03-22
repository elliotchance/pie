package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var subSliceTests = []struct {
	ss       []float64
	start    int
	end      int
	subSlice []float64
}{
	{
		nil,
		1,
		1,
		nil,
	},
	{
		nil,
		1,
		2,
		[]float64{0},
	},
	{
		[]float64{},
		1,
		1,
		nil,
	},
	{
		[]float64{},
		1,
		2,
		[]float64{0},
	},
	{
		[]float64{1.23, 2.34},
		-1,
		-1,
		nil,
	},
	{
		[]float64{1.23, 2.34},
		-1,
		1,
		nil,
	},
	{
		[]float64{1.23, 2.34},
		1,
		-1,
		nil,
	},
	{
		[]float64{1.23, 2.34},
		2,
		0,
		nil,
	},

	{
		[]float64{1.23, 2.34},
		1,
		1,
		nil,
	},
	{
		[]float64{1.23, 2.34},
		1,
		2,
		[]float64{2.34},
	},
	{
		[]float64{1.23, 2.34},
		1,
		3,
		[]float64{2.34, 0},
	},
	{
		[]float64{1.23, 2.34},
		2,
		2,
		nil,
	},
	{
		[]float64{1.23, 2.34},
		2,
		3,
		[]float64{0},
	},
	{
		[]float64{1.23, 2.34, 0},
		2,
		3,
		[]float64{0},
	},
}

func TestSubSlice(t *testing.T) {
	for _, test := range subSliceTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.subSlice, pie.SubSlice(test.ss, test.start, test.end))
		})
	}
}
