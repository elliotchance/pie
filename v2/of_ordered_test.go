package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
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
}
