package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedian(t *testing.T) {
	assert.Equal(t, 0.0, pie.Median([]float64{}))
	assert.Equal(t, 12.3, pie.Median([]float64{12.3}))
	assert.Equal(t, 8.4, pie.Median([]float64{12.3, 4.5}))
	assert.Equal(t, 4.5, pie.Median([]float64{2.1, 12.3, 4.5}))
}
