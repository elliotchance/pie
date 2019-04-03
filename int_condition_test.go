package pie_test

import (
	"fmt"
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"testing"
)

var intComparisonTests = []struct {
	a, b                     int
	expectedEqual            bool
	expectedNotEqual         bool
	expectedGreaterThan      bool
	expectedGreaterThanEqual bool
	expectedLessThan         bool
	expectedLessThanEqual    bool
}{
	{200, 100, false, true, true, true, false, false},
	{100, 100, true, false, false, true, false, true},
	{200, 50, false, true, true, true, false, false},
	{100, 200, false, true, false, false, true, true},
}

func TestEqualInt(t *testing.T) {
	for _, test := range intComparisonTests {
		t.Run(fmt.Sprintf("%d %d", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedEqual, pie.EqualInt(test.b)(test.a))
		})
	}
}

func TestNotEqualInt(t *testing.T) {
	for _, test := range intComparisonTests {
		t.Run(fmt.Sprintf("%d %d", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedNotEqual, pie.NotEqualInt(test.b)(test.a))
		})
	}
}

func TestGreaterThanInt(t *testing.T) {
	for _, test := range intComparisonTests {
		t.Run(fmt.Sprintf("%d %d", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedGreaterThan, pie.GreaterThanInt(test.b)(test.a))
		})
	}
}

func TestGreaterThanEqualInt(t *testing.T) {
	for _, test := range intComparisonTests {
		t.Run(fmt.Sprintf("%d %d", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedGreaterThanEqual, pie.GreaterThanEqualInt(test.b)(test.a))
		})
	}
}

func TestLessThanInt(t *testing.T) {
	for _, test := range intComparisonTests {
		t.Run(fmt.Sprintf("%d %d", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedLessThan, pie.LessThanInt(test.b)(test.a))
		})
	}
}

func TestLessThanEqualInt(t *testing.T) {
	for _, test := range intComparisonTests {
		t.Run(fmt.Sprintf("%d %d", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedLessThanEqual, pie.LessThanEqualInt(test.b)(test.a))
		})
	}
}
