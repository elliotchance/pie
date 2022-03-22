package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var containsTests = []struct {
	ss       []float64
	contains float64
	expected bool
}{
	{nil, 1, false},
	{[]float64{1, 2, 3}, 1, true},
	{[]float64{1, 2, 3}, 2, true},
	{[]float64{1, 2, 3}, 3, true},
	{[]float64{1, 2, 3}, 4, false},
	{[]float64{1, 2, 3}, 5, false},
	{[]float64{1, 2, 3}, 6, false},
	{[]float64{1, 5, 3}, 5, true},
}

func TestContains(t *testing.T) {
	for _, test := range containsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Contains(test.ss, test.contains))
		})
	}
}
