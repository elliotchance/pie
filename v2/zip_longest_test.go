package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZipLongest(t *testing.T) {
	for _, test := range zipTests {

		t.Run("", func(t *testing.T) {
			for i, pair := range pie.ZipLongest(test.ss1, test.ss2) {
				var a int
				var b float32

				if i < len(test.ss1) {
					a = test.ss1[i]
				}
				if i < len(test.ss2) {
					b = test.ss2[i]
				}

				assert.Equal(t, pair.A, a)
				assert.Equal(t, pair.B, b)
			}
		})
	}
}
