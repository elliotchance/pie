package pie_test

import (
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddFloat64(t *testing.T) {
	assert.Equal(t, 8.3, pie.AddFloat64(3.2)(5.1))
}

func TestEqualFloat64(t *testing.T) {
	assert.False(t, pie.EqualFloat64(3.1)(5.1))
	assert.True(t, pie.EqualFloat64(5.1)(5.1))
}
