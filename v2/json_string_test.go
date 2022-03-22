package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJSONString(t *testing.T) {
	for _, test := range jsonTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.jsonString, pie.JSONString(test.ss))
		})
	}
}
