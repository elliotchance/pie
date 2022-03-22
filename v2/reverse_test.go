package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	for _, test := range sortTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.reversed, pie.Reverse(test.ss))
		})
	}
}
