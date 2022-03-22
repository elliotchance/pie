package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInts(t *testing.T) {
	assert.Equal(t, []int(nil), pie.Ints([]int(nil)))

	assert.Equal(t,
		[]int{92, 823, 453},
		pie.Ints([]float64{92.384, 823.324, 453}))
}
