package pie_test

import (
	"testing"

	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
)

var uniqueStableTests = []struct {
	ss           []float64
	uniqueStable []float64
}{
	{
		nil,
		nil,
	},
	{
		[]float64{},
		[]float64{},
	},
	{
		[]float64{789},
		[]float64{789},
	},
	{
		[]float64{12.789, -13.2, 12.789},
		[]float64{12.789, -13.2},
	},
	{
		[]float64{12.789, -13.2, 1.234e6, 789},
		[]float64{12.789, -13.2, 1.234e6, 789},
	},
}

func TestUniqueStable(t *testing.T) {
	for _, test := range uniqueStableTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.uniqueStable, pie.UniqueStable(test.ss))
		})
	}
}
