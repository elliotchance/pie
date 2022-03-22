package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	for _, test := range sortTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.sorted, pie.Sort(test.ss))
		})
	}
}
