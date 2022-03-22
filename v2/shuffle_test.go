package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var shuffleTests = []struct {
	ss       []float64
	expected []float64
	source   rand.Source
}{
	{
		nil,
		nil,
		nil,
	},
	{
		nil,
		nil,
		rand.NewSource(0),
	},
	{
		[]float64{},
		[]float64{},
		rand.NewSource(0),
	},
	{
		[]float64{12.3, 2.34, 4.56},
		[]float64{2.34, 12.3, 4.56},
		rand.NewSource(0),
	},
	{
		[]float64{12.3},
		[]float64{12.3},
		rand.NewSource(0),
	},
}

func TestShuffle(t *testing.T) {
	for _, test := range shuffleTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Shuffle(test.ss, test.source))
		})
	}
}
