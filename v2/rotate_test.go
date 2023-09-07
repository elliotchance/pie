package pie_test

import (
	"testing"

	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
)

var rotateTests = []struct {
	ss      []float64
	rotated []float64
	n       int
}{
	{
		nil,
		nil,
		0,
	},
	{
		[]float64{1.23, 2.34},
		[]float64{1.23, 2.34},
		0,
	},
	{
		[]float64{},
		[]float64{},
		0,
	},
	{
		[]float64{},
		[]float64{},
		3,
	},
	{
		[]float64{1.23, 2.34},
		[]float64{2.34, 1.23},
		1,
	},
	{
		[]float64{1.23, 2.34},
		[]float64{2.34, 1.23},
		-1,
	},
	{
		[]float64{1.23},
		[]float64{1.23},
		1000,
	},
	{
		[]float64{1.23, 2.34, 3.45},
		[]float64{1.23, 2.34, 3.45},
		3,
	},
	{
		[]float64{1.23, 2.34, 3.45},
		[]float64{1.23, 2.34, 3.45},
		-3,
	},
	{
		[]float64{1.23, 2.34, 3.45},
		[]float64{1.23, 2.34, 3.45},
		6,
	},
	{
		[]float64{1.23, 2.34, 3.45},
		[]float64{1.23, 2.34, 3.45},
		-6,
	},
	{
		[]float64{1.23, 2.34, 3.45},
		[]float64{2.34, 3.45, 1.23},
		-1,
	},
	{
		[]float64{1.23, 2.34, 3.45},
		[]float64{3.45, 1.23, 2.34},
		1,
	},
}

func TestRotate(t *testing.T) {
	for _, test := range rotateTests {
		t.Run("", func(t *testing.T) {
			rotated := pie.Rotate(test.ss, test.n)
			assert.Equal(t, test.rotated, rotated)
		})
	}
}
