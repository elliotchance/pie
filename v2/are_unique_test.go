package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var uniqueTests = []struct {
	ss        []float64
	unique    []float64
	areUnique bool
}{
	{
		nil,
		nil,
		true,
	},
	{
		[]float64{},
		[]float64{},
		true,
	},
	{
		[]float64{789},
		[]float64{789},
		true,
	},
	{
		[]float64{12.789, -13.2, 12.789},
		[]float64{-13.2, 12.789},
		false,
	},
	{
		[]float64{12.789, -13.2, 1.234e6, 789},
		[]float64{-13.2, 12.789, 789, 1.234e6},
		true,
	},
}

func TestAreUnique(t *testing.T) {
	for _, test := range uniqueTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.areUnique, pie.AreUnique(test.ss))
		})
	}
}
