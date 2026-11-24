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
	"SharedDuplicates": {
		[]float64{5, 5, 5},
		[]float64{5, 5, 6},
		[]float64{6},
		[]float64{5},
	},
	"EqualDuplicates": {
		[]float64{5, 5, 5},
		[]float64{5, 5, 5},
		nil,
		nil,
	},
	"ReorderedDuplicates": {
		[]float64{5, 6, 5, 6},
		[]float64{6, 6, 5, 5},
		nil,
		nil,
	},
	"MoreDuplicatesAdded": {
		[]float64{5, 5},
		[]float64{5, 5, 5, 5},
		[]float64{5, 5},
		nil,
	},
	"MoreDuplicatesRemoved": {
		[]float64{5, 5, 5, 5},
		[]float64{5, 5},
		nil,
		[]float64{5, 5},
	},
	"DuplicateSurplusOrder": {
		[]float64{1, 2, 1, 2, 1, 4, 2},
		[]float64{1, 2, 3, 1, 3, 5},
		[]float64{3, 3, 5},
		[]float64{2, 1, 4, 2},
	},
}

func TestDiff(t *testing.T) {
	for testName, test := range diffTests {
		t.Run(testName, func(t *testing.T) {
			ss1 := append([]float64(nil), test.ss1...)
			ss2 := append([]float64(nil), test.ss2...)
			added, removed := pie.Diff(test.ss1, test.ss2)
			assert.Equal(t, test.added, added)
			assert.Equal(t, test.removed, removed)
			assert.Equal(t, ss1, test.ss1)
			assert.Equal(t, ss2, test.ss2)
		})
	}
}
