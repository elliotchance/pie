package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	assert.Equal(t, "123", pie.String(123))
	assert.Equal(t, "1.89", pie.String(1.89))
	assert.Equal(t, "1.89", pie.String("1.89"))
}
