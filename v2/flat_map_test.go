package pie_test

import (
	"testing"

	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
)

var flatMapTests = []struct {
	ss       []float64
	fn       func(float64) []float64
	expected []float64
}{
	{
		nil,
		nil,
		nil,
	},
	{
		[]float64{100, 103},
		func(i float64) []float64 { return []float64{i + 42, 42} },
		[]float64{142, 42, 145, 42},
	},
}

func TestFlatMap(t *testing.T) {
	for _, test := range flatMapTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.FlatMap(test.ss, test.fn))
		})
	}
}
