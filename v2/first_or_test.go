package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var firstAndLastTests = []struct {
	ss      []float64
	firstOr float64
	lastOr  float64
}{
	{
		nil,
		102,
		202,
	},
	{
		[]float64{100},
		100,
		100,
	},
	{
		[]float64{1, 2},
		1,
		2,
	},
	{
		[]float64{1, 2, 3},
		1,
		3,
	},
}

func TestFirstOr(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.firstOr, pie.FirstOr(test.ss, 102))
		})
	}
}
