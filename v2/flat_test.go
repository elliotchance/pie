package pie_test

import (
	"testing"

	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
)

var flatTests = []struct {
	ss   [][]float64
	flat []float64
}{
	{
		nil,
		nil,
	},
	{
		[][]float64{{100}},
		[]float64{100},
	},
	{
		[][]float64{{100}, {101, 102}, {102, 103}},
		[]float64{100, 101, 102, 102, 103},
	},
	{
		[][]float64{nil, {101, 102}, {}},
		[]float64{101, 102},
	},
	{
		[][]float64{nil, nil},
		nil,
	},
}

func TestFlat(t *testing.T) {
	for _, test := range flatTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.flat, pie.Flat(test.ss))
		})
	}
}
