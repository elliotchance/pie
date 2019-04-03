package pie_test

import (
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddInt(t *testing.T) {
	assert.Equal(t, 8, pie.AddInt(3)(5))
}
