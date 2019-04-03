package pie_test

import (
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddInt(t *testing.T) {
	assert.Equal(t, 8, pie.AddInt(3)(5))
}

func TestEqualInt(t *testing.T) {
	assert.False(t, pie.EqualInt(3)(5))
	assert.True(t, pie.EqualInt(5)(5))
}
