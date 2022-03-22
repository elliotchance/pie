package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func stringShorter(a, b string) bool {
	return len(a) < len(b)
}

var sortByLengthTests = []struct {
	ss           []string
	sortedStable []string
}{
	{
		nil,
		nil,
	},
	{
		[]string{},
		[]string{},
	},
	{
		[]string{"foo"},
		[]string{"foo"},
	},
	{
		[]string{"aaa", "b", "cc"},
		[]string{"b", "cc", "aaa"},
	},
	{
		[]string{"zz", "aaa", "b", "cc"},
		[]string{"b", "zz", "cc", "aaa"},
	},
}

func TestSortStableUsing(t *testing.T) {
	less := stringShorter
	for _, test := range sortByLengthTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.sortedStable, pie.SortStableUsing(test.ss, less))
		})
	}
}
