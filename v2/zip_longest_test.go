package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZipLongest(t *testing.T) {
	for _, test := range zipTests {

		t.Run("", func(t *testing.T) {
			c := pie.ZipLongest(test.ss1, test.ss2)

			assert.Equal(t, c, test.expectedLong)
		})
	}
}
