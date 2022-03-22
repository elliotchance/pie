package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sortTests = []struct {
	ss        []float64
	sorted    []float64
	reversed  []float64
	areSorted bool
}{
	{
		nil,
		nil,
		nil,
		true,
	},
	{
		[]float64{},
		[]float64{},
		[]float64{},
		true,
	},
	{
		[]float64{789},
		[]float64{789},
		[]float64{789},
		true,
	},
	{
		[]float64{12.789, -13.2, 789},
		[]float64{-13.2, 12.789, 789},
		[]float64{789, -13.2, 12.789},
		false,
	},
	{
		[]float64{12.789, -13.2, 1.234e6, 789},
		[]float64{-13.2, 12.789, 789, 1.234e6},
		[]float64{789, 1.234e6, -13.2, 12.789},
		false,
	},
	{
		[]float64{-13.2, 12.789},
		[]float64{-13.2, 12.789},
		[]float64{12.789, -13.2},
		true,
	},
}

func TestAreSorted(t *testing.T) {
	for _, test := range sortTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.areSorted, pie.AreSorted(test.ss))
		})
	}
}
