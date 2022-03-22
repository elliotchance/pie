package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var statsTests = []struct {
	ss                     []float64
	min, max, sum, product float64
	len                    int
	mode                   []float64
	average                float64
}{
	{
		nil,
		0,
		0,
		0,
		0,
		0,
		nil,
		0,
	},
	{
		[]float64{},
		0,
		0,
		0,
		0,
		0,
		[]float64{},
		0,
	},
	{
		[]float64{1.5},
		1.5,
		1.5,
		1.5,
		1.5,
		1,
		[]float64{1.5},
		1.5,
	},
	{
		[]float64{2.2, 3.1, 5.1, 1.9},
		1.9,
		5.1,
		12.3,
		66.0858,
		4,
		[]float64{2.2, 3.1, 5.1, 1.9},
		3.075,
	},
}

func TestAverage(t *testing.T) {
	for _, test := range statsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.average, pie.Average(test.ss))
		})
	}
}
