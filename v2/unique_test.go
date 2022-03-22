package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnique(t *testing.T) {
	for _, test := range uniqueTests {
		t.Run("", func(t *testing.T) {
			// We have to sort the unique slice because it is always returned in
			// random order.
			assert.Equal(t, test.unique, pie.Sort(pie.Unique(test.ss)))
		})
	}
}
