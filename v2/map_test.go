package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	for _, test := range selectTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedMap, pie.Map(test.ss, func(a float64) float64 {
				return a + 5.2
			}))
		})
	}
}
