package pie_test

import (
	"strings"
	"testing"

	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	t.Run("chaining", func(t *testing.T) {
		name := pie.Of([]string{"Bob", "Sally", "John", "Jane"}).
			FilterNot(func(name string) bool {
				return strings.HasPrefix(name, "J")
			}).
			Map(strings.ToUpper).
			LastOr("")

		assert.Equal(t, "SALLY", name)
	})

	t.Run("result", func(t *testing.T) {
		names := pie.Of([]string{"Bob", "Sally", "John", "Jane"}).
			FilterNot(func(name string) bool {
				return strings.HasPrefix(name, "J")
			}).
			Result

		assert.Equal(t, []string{"Bob", "Sally"}, names)
	})

	t.Run("delete", func(t *testing.T) {
		names := pie.Of([]string{"Bob", "Sally", "John", "Jane"}).
			Delete(2, 3).
			Result

		assert.Equal(t, []string{"Bob", "Sally"}, names)
	})

	t.Run("rotate", func(t *testing.T) {
		names := pie.Of([]string{"Bob", "Sally", "John", "Jane"}).
			Rotate(1).
			Result

		assert.Equal(t, []string{"Jane", "Bob", "Sally", "John"}, names)
	})
}
