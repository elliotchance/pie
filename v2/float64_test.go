package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat64(t *testing.T) {
	assert.Equal(t, 123.0, pie.Float64(123))
	assert.Equal(t, 1.89, pie.Float64(1.89))
	assert.Equal(t, 1.89, pie.Float64("1.89"))
}
