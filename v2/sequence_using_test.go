package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSequenceUsing(t *testing.T) {
	for _, test := range sequenceAndSequenceUsingTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.SequenceUsing(test.ss,
				func(i int) float64 { return float64(i) }, test.params...))
		})
	}
}
