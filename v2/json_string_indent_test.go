package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJSONStringIndent(t *testing.T) {
	for _, test := range jsonIndentTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.jsonString, pie.JSONStringIndent(test.ss, "", "  "))
		})
	}
}
