package pie_test

import (
	"testing"

	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
)

func TestOfOrdered(t *testing.T) {
	t.Run("chaining", func(t *testing.T) {
		name := pie.OfOrdered([]string{"Bob", "Sally", "John", "Jane"}).
			Join("+")

		assert.Equal(t, "Bob+Sally+John+Jane", name)
	})

	t.Run("result", func(t *testing.T) {
		names := pie.OfOrdered([]string{"Bob", "Sally", "John", "Jane"}).
			Filter(func(s string) bool {
				return len(s) < 4
			}).
			Result

		assert.Equal(t, []string{"Bob"}, names)
	})

	t.Run("delete", func(t *testing.T) {
		names := pie.Of([]string{"Bob", "Sally", "John", "Jane"}).
			Delete(2, 3).
			Result

		assert.Equal(t, []string{"Bob", "Sally"}, names)
	})

	t.Run("unique_stable", func(t *testing.T) {
		names := pie.OfNumeric([]float64{-4.56, 1.23, -4.56}).
			UniqueStable().
			Result

		assert.Equal(t, []float64{-4.56, 1.23}, names)
	})
}
