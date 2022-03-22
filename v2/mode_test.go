package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMode(t *testing.T) {
	cmp := func(a, b []float64) bool {
		m := make(map[float64]struct{})
		for _, i := range a {
			m[i] = struct{}{}
		}

		for _, i := range b {
			if _, ok := m[i]; !ok {
				return false
			}
		}

		return true
	}

	for _, test := range statsTests {
		t.Run("", func(t *testing.T) {
			assert.True(t, cmp(test.mode, pie.Mode(test.ss)))
		})
	}
}
