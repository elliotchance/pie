package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	assert.Equal(t, []float64(nil), pie.Insert([]float64(nil), 0))
	assert.Equal(t, []float64{2.0, 1.0}, pie.Insert([]float64{1.0}, 0, 2.0))
	assert.Equal(t, []float64{1.0, 2.0}, pie.Insert([]float64{1.0}, 1, 2.0))
	assert.Equal(t, []float64{1.0, 2.0, 3.3}, pie.Insert([]float64{1.0, 3.3}, 1, 2.0))
}
