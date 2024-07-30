package pie_test

import (
	"testing"

	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
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

	t.Run("unique_stable", func(t *testing.T) {
		names := pie.OfNumeric([]float64{-4.56, 1.23, -4.56}).
			UniqueStable().
			Result

		assert.Equal(t, []float64{-4.56, 1.23}, names)
	})
}
