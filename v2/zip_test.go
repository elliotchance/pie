package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var zipTests = []struct {
	ss1           []int
	ss2           []float32
	expectedShort []pie.Zipped[int, float32]
	expectedLong  []pie.Zipped[int, float32]
}{
	{
		[]int{},
		[]float32{},
		[]pie.Zipped[int, float32]{},
		[]pie.Zipped[int, float32]{},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]float32{},
		[]pie.Zipped[int, float32]{},
		[]pie.Zipped[int, float32]{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}},
	},
	{
		[]int{},
		[]float32{1.0, 2.0, 3.0, 4.0, 5.0},
		[]pie.Zipped[int, float32]{},
		[]pie.Zipped[int, float32]{{0, 1.0}, {0, 2.0}, {0, 3.0}, {0, 4.0}, {0, 5.0}},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]float32{1.0, 2.0, 3.0, 4.0, 5.0},
		[]pie.Zipped[int, float32]{{1, 1.0}, {2, 2.0}, {3, 3.0}, {4, 4.0}, {5, 5.0}},
		[]pie.Zipped[int, float32]{{1, 1.0}, {2, 2.0}, {3, 3.0}, {4, 4.0}, {5, 5.0}},
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		[]float32{1.0, 2.0, 3.0, 4.0, 5.0},
		[]pie.Zipped[int, float32]{{1, 1.0}, {2, 2.0}, {3, 3.0}, {4, 4.0}, {5, 5.0}},
		[]pie.Zipped[int, float32]{{1, 1.0}, {2, 2.0}, {3, 3.0}, {4, 4.0}, {5, 5.0}, {6, 0}, {7, 0}, {8, 0}},
	},
	{
		[]int{1, 2, 3},
		[]float32{1.0, 2.0, 3.0, 4.0, 5.0},
		[]pie.Zipped[int, float32]{{1, 1.0}, {2, 2.0}, {3, 3.0}},
		[]pie.Zipped[int, float32]{{1, 1.0}, {2, 2.0}, {3, 3.0}, {0, 4.0}, {0, 5.0}},
	},
}

func TestZip(t *testing.T) {
	for _, test := range zipTests {

		t.Run("", func(t *testing.T) {
			c := pie.Zip(test.ss1, test.ss2)

			assert.Equal(t, c, test.expectedShort)
		})
	}
}
