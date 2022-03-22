package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var randomTests = []struct {
	ss       []float64
	expected float64
	source   rand.Source
}{
	{
		nil,
		0.0,
		nil,
	},
	{
		nil,
		0.0,
		rand.NewSource(0),
	},
	{
		[]float64{},
		0.0,
		rand.NewSource(0),
	},
	{
		[]float64{12.3, 2.34, 4.56},
		12.3,
		rand.NewSource(0),
	},
	{
		[]float64{12.3, 2.34, 4.56},
		4.56,
		rand.NewSource(1),
	},
	{
		[]float64{12.3},
		12.3,
		rand.NewSource(0),
	},
}

func TestRandom(t *testing.T) {
	for _, test := range randomTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Random(test.ss, test.source))
		})
	}
}
