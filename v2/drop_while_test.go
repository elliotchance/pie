package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dropWhileTests = []struct {
	ss        []float64
	f         func(s float64) bool
	dropWhile []float64
}{
	{
		ss:        nil,
		f:         func(s float64) bool { return s == 0.1 },
		dropWhile: []float64{},
	},
	{
		ss:        []float64{2.1, 2.1, 2.1, 7.2, 8.1},
		f:         func(s float64) bool { return s == 2.1 },
		dropWhile: []float64{7.2, 8.1},
	},
	{
		ss:        []float64{2.1, 4.1, 5.1, 7.2, 8.1},
		f:         func(s float64) bool { return s == 0.2 },
		dropWhile: []float64{2.1, 4.1, 5.1, 7.2, 8.1},
	},
	{
		ss:        []float64{2.1, 2.1, 2.1, 2.1, 2.1},
		f:         func(s float64) bool { return s == 2.1 },
		dropWhile: []float64{},
	},
}

func TestDropWhile(t *testing.T) {
	for _, test := range dropWhileTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.dropWhile, pie.DropWhile(test.ss, test.f))
		})
	}
}
