package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 584.2727, pie.Abs(-584.2727))
	assert.Equal(t, 5, pie.Abs(-5))
	assert.Equal(t, 584.2727, pie.Abs(584.2727))
}
