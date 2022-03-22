package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct(t *testing.T) {
	for _, test := range statsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.product, pie.Product(test.ss))
		})
	}
}
