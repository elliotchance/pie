package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPop(t *testing.T) {
	numbers := []float64{42.0, 4.2}

	assert.Equal(t, 42.0, *pie.Pop(&numbers))
	assert.Equal(t, []float64{4.2}, numbers)

	assert.Equal(t, 4.2, *pie.Pop(&numbers))
	assert.Equal(t, []float64{}, numbers)
}
