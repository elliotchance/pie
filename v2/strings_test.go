package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrings(t *testing.T) {
	assert.Equal(t, []string{}, pie.Strings([]float64{}))

	assert.Equal(t,
		[]string{"92.384", "823.324", "453"},
		pie.Strings([]float64{92.384, 823.324, 453}))
}
