package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTop(t *testing.T) {
	for _, test := range topAndBottomTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.top, pie.Top(test.ss, test.n))
		})
	}
}
