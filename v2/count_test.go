package pie_test

import (
	"testing"

	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
)

var countTests = []struct {
	ss         []float64
	lookingFor float64
	expected   int
}{
	{nil, 1, 0},
	{[]float64{}, 1, 0},
	{[]float64{1, 2, 3}, 4, 0},
	{[]float64{1, 2, 3}, 1, 1},
	{[]float64{1, 2, 3}, 3, 1},
	{[]float64{1, 2, 2, 3, 2}, 2, 3},
	{[]float64{5, 5, 5}, 5, 3},
}

func TestCount(t *testing.T) {
	for _, test := range countTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Count(test.ss, test.lookingFor))
		})
	}
}
