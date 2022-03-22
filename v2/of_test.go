package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
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
}
