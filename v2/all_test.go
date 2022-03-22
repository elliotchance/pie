package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	assert.True(t,
		pie.All([]float64{}, func(value float64) bool {
			return false
		}),
	)

	// None
	assert.False(t,
		pie.All([]float64{12.3, 4.56}, func(value float64) bool {
			return value == 1
		}),
	)

	// Some
	assert.False(t,
		pie.All([]float64{12.3, 4.56}, func(value float64) bool {
			return value == 12.3
		}),
	)

	// All
	assert.True(t,
		pie.All([]float64{12.3, 4.56}, func(value float64) bool {
			return value > 0
		}),
	)
}
