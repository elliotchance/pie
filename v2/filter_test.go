package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var selectTests = []struct {
	ss                []float64
	condition         func(float64) bool
	expectedFilter    []float64
	expectedFilterNot []float64
	expectedMap       []float64
}{
	{
		nil,
		func(s float64) bool {
			return s == 5
		},
		nil,
		nil,
		nil,
	},
	{
		[]float64{1, 2, 3},
		func(s float64) bool {
			return s != 2
		},
		[]float64{1, 3},
		[]float64{2},
		[]float64{6.2, 7.2, 8.2},
	},
}

func TestFilter(t *testing.T) {
	for _, test := range selectTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedFilter, pie.Filter(test.ss, test.condition))
		})
	}
}
