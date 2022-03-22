package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroup(t *testing.T) {
	assert.Equal(t, map[float64]int{}, pie.Group([]float64{}))

	assert.Equal(t, map[float64]int{
		1: 1,
	}, pie.Group([]float64{1}))

	assert.Equal(t, map[float64]int{
		1: 1,
		2: 2,
		3: 3,
	}, pie.Group([]float64{1, 2, 2, 3, 3, 3}))
}
