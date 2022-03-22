package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var diffTests = map[string]struct {
	ss1     []float64
	ss2     []float64
	added   []float64
	removed []float64
}{
	"BothEmpty": {
		nil,
		nil,
		nil,
		nil,
	},
	"OnlyRemovedUnique": {
		[]float64{4334.5435, 879.123},
		nil,
		nil,
		[]float64{4334.5435, 879.123},
	},
	"OnlyRemovedDuplicates": {
		[]float64{4334.5435, 92.384, 4334.5435},
		nil,
		nil,
		[]float64{4334.5435, 92.384, 4334.5435},
	},
	"OnlyAddedUnique": {
		nil,
		[]float64{879.123, 92.384},
		[]float64{879.123, 92.384},
		nil,
	},
	"OnlyAddedDuplicates": {
		nil,
		[]float64{879.123, 92.384, 92.384, 879.123},
		[]float64{879.123, 92.384, 92.384, 879.123},
		nil,
	},
	"AddedAndRemovedUnique": {
		[]float64{4334.5435, 879.123, 92.384, 823.324},
		[]float64{92.384, 823.324, 453, 3.345},
		[]float64{453, 3.345},
		[]float64{4334.5435, 879.123},
	},
	"AddedAndRemovedDuplicates": {
		[]float64{4334.5435, 879.123, 92.384, 92.384, 823.324},
		[]float64{92.384, 823.324, 453, 823.324, 3.345},
		[]float64{453, 823.324, 3.345},
		[]float64{4334.5435, 879.123, 92.384},
	},
}

func TestDiff(t *testing.T) {
	for testName, test := range diffTests {
		t.Run(testName, func(t *testing.T) {
			added, removed := pie.Diff(test.ss1, test.ss2)
			assert.Equal(t, test.added, added)
			assert.Equal(t, test.removed, removed)
		})
	}
}
