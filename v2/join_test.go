package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoin(t *testing.T) {
	assert.Equal(t, "", pie.Join([]float64(nil), "a"))
	assert.Equal(t, "", pie.Join([]float64{}, "a"))
	var f1, f2 float64 = 0.1, 20000000000000000
	assert.Equal(t, "0.1;2e+16", pie.Join([]float64{f1, f2}, ";"))
}
