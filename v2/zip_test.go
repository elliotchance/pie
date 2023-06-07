package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var zipTests = []struct {
	Ss1 []int
	Ss2 []float32
}{
	{
		[]int{},
		[]float32{},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]float32{},
	},
	{
		[]int{},
		[]float32{1.0, 2.0, 3.0, 4.0, 5.0},
	},
	// the same length
	{
		[]int{1, 2, 3, 4, 5},
		[]float32{1.0, 2.0, 3.0, 4.0, 5.0},
	},
	// Ss1 bigger
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		[]float32{1.0, 2.0, 3.0, 4.0, 5.0},
	},
	// Ss1 less
	{
		[]int{1, 2, 3},
		[]float32{1.0, 2.0, 3.0, 4.0, 5.0},
	},
}

func TestZip(t *testing.T) {

	for _, test := range zipTests {

		t.Run("", func(t *testing.T) {
			for i, pair := range pie.Zip(test.Ss1, test.Ss2) {
				var a int
				var b float32

				if i < len(test.Ss1) {
					a = test.Ss1[i]
				}
				if i < len(test.Ss2) {
					b = test.Ss2[i]
				}

				assert.Equal(t, pair.A, a)
				assert.Equal(t, pair.B, b)
			}
		})
	}
}
