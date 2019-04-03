package pie_test

import (
	"fmt"
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"testing"
)

var float64ComparisonTests = []struct {
	a, b                     float64
	expectedEqual            bool
	expectedNotEqual         bool
	expectedGreaterThan      bool
	expectedGreaterThanEqual bool
	expectedLessThan         bool
	expectedLessThanEqual    bool
}{
	{200.5, 100.1, false, true, true, true, false, false},
	{100.1, 100.1, true, false, false, true, false, true},
	{200.5, 50.7, false, true, true, true, false, false},
	{100.1, 200.5, false, true, false, false, true, true},
}

func TestEqualFloat64(t *testing.T) {
	for _, test := range float64ComparisonTests {
		t.Run(fmt.Sprintf("%f %f", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedEqual, pie.EqualFloat64(test.b)(test.a))
		})
	}
}

func TestNotEqualFloat64(t *testing.T) {
	for _, test := range float64ComparisonTests {
		t.Run(fmt.Sprintf("%f %f", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedNotEqual, pie.NotEqualFloat64(test.b)(test.a))
		})
	}
}

func TestGreaterThanFloat64(t *testing.T) {
	for _, test := range float64ComparisonTests {
		t.Run(fmt.Sprintf("%f %f", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedGreaterThan, pie.GreaterThanFloat64(test.b)(test.a))
		})
	}
}

func TestGreaterThanEqualFloat64(t *testing.T) {
	for _, test := range float64ComparisonTests {
		t.Run(fmt.Sprintf("%f %f", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedGreaterThanEqual, pie.GreaterThanEqualFloat64(test.b)(test.a))
		})
	}
}

func TestLessThanFloat64(t *testing.T) {
	for _, test := range float64ComparisonTests {
		t.Run(fmt.Sprintf("%f %f", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedLessThan, pie.LessThanFloat64(test.b)(test.a))
		})
	}
}

func TestLessThanEqualFloat64(t *testing.T) {
	for _, test := range float64ComparisonTests {
		t.Run(fmt.Sprintf("%f %f", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedLessThanEqual, pie.LessThanEqualFloat64(test.b)(test.a))
		})
	}
}
