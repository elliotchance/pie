package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastOr(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.lastOr, pie.LastOr(test.ss, 202))
		})
	}
}
