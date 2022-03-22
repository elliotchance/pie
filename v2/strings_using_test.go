package pie_test

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var stringsUsingTests = []struct {
	ss        []float64
	transform func(float64) string
	expected  []string
}{
	{
		nil,
		func(s float64) string {
			return "foo"
		},
		nil,
	},
	{
		[]float64{},
		func(s float64) string {
			return fmt.Sprintf("%f!", s)
		},
		[]string{},
	},
	{
		[]float64{6.2, 7.2, 8.2},
		func(s float64) string {
			return fmt.Sprintf("%.2f!", s)
		},
		[]string{"6.20!", "7.20!", "8.20!"},
	},
}

func TestStringsUsing(t *testing.T) {
	for _, test := range stringsUsingTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.StringsUsing(test.ss, test.transform))
		})
	}
}
