package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var equalsTests = []struct {
	ss       []float64
	rhs      []float64
	expected bool
}{
	{nil, nil, true},
	{[]float64{}, []float64{}, true},
	{nil, []float64{}, true},
	{[]float64{1.0, 2.0}, []float64{1.0, 2.0}, true},
	{[]float64{1.0, 2.0}, []float64{1.0, 5.0}, false},
	{[]float64{1.0, 2.0}, []float64{1.0}, false},
	{[]float64{1.0}, []float64{2.0}, false},
	{[]float64{1.0}, nil, false},
}

func TestEquals(t *testing.T) {
	for _, test := range equalsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Equals(test.ss, test.rhs))
		})
	}
}
