package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	assert.False(t,
		pie.Any([]float64{}, func(value float64) bool {
			return true
		}),
	)

	// None
	assert.False(t,
		pie.Any([]float64{12.3, 4.56}, func(value float64) bool {
			return value == 1
		}),
	)

	// Some
	assert.True(t,
		pie.Any([]float64{12.3, 4.56}, func(value float64) bool {
			return value == 12.3
		}),
	)

	// All
	assert.True(t,
		pie.Any([]float64{12.3, 4.56}, func(value float64) bool {
			return value > 0
		}),
	)
}
