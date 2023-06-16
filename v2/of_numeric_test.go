package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOfONumeric(t *testing.T) {
	t.Run("chaining", func(t *testing.T) {
		total := pie.OfNumeric([]float64{123, 456}).
			Sum()

		assert.Equal(t, 579.0, total)
	})

	t.Run("result", func(t *testing.T) {
		names := pie.OfNumeric([]float64{1.23, 4.56}).
			Filter(func(x float64) bool {
				return x < 4
			}).
			Result

		assert.Equal(t, []float64{1.23}, names)
	})

	t.Run("delete", func(t *testing.T) {
		names := pie.OfNumeric([]float64{1.23, 4.56}).
			Delete(1).
			Result

		assert.Equal(t, []float64{1.23}, names)
	})
}
