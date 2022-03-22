package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat64s(t *testing.T) {
	assert.Equal(t, []float64(nil), pie.Float64s([]float64(nil)))

	assert.Equal(t,
		[]float64{92.384, 823.324, 453},
		pie.Float64s([]float64{92.384, 823.324, 453}))
}
