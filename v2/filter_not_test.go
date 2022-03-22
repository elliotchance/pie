package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterNot(t *testing.T) {
	for _, test := range selectTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedFilterNot, pie.FilterNot(test.ss, test.condition))
		})
	}
}
