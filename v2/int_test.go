package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(t *testing.T) {
	assert.Equal(t, 123, pie.Int(123))
	assert.Equal(t, 1, pie.Int(1.89))
	assert.Equal(t, 1, pie.Int("1.89"))
}
