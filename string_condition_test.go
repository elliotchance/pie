package pie_test

import (
	"fmt"
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrefix(t *testing.T) {
	for _, test := range []struct {
		s        string
		prefix   string
		expected bool
	}{
		{"", "a", false},
		{"", "", true},
		{"a", "a", true},
		{"ab", "a", true},
		{"a", "b", false},
		{"a", "A", false},
		{"ab", "b", false},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Prefix(test.prefix)(test.s))
		})
	}
}

func TestSuffix(t *testing.T) {
	for _, test := range []struct {
		s        string
		suffix   string
		expected bool
	}{
		{"", "a", false},
		{"", "", true},
		{"a", "a", true},
		{"ab", "a", false},
		{"a", "b", false},
		{"a", "A", false},
		{"ab", "b", true},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Suffix(test.suffix)(test.s))
		})
	}
}

var stringComparisonTests = []struct {
	a, b                     string
	expectedEqual            bool
	expectedNotEqual         bool
	expectedGreaterThan      bool
	expectedGreaterThanEqual bool
	expectedLessThan         bool
	expectedLessThanEqual    bool
}{
	{"foo", "bar", false, true, true, true, false, false},
	{"bar", "bar", true, false, false, true, false, true},
	{"foo", "Foo", false, true, true, true, false, false},
	{"bar", "foo", false, true, false, false, true, true},
}

func TestEqualString(t *testing.T) {
	for _, test := range stringComparisonTests {
		t.Run(fmt.Sprintf("%s %s", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedEqual, pie.EqualString(test.b)(test.a))
		})
	}
}

func TestNotEqualString(t *testing.T) {
	for _, test := range stringComparisonTests {
		t.Run(fmt.Sprintf("%s %s", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedNotEqual, pie.NotEqualString(test.b)(test.a))
		})
	}
}

func TestGreaterThanString(t *testing.T) {
	for _, test := range stringComparisonTests {
		t.Run(fmt.Sprintf("%s %s", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedGreaterThan, pie.GreaterThanString(test.b)(test.a))
		})
	}
}

func TestGreaterThanEqualString(t *testing.T) {
	for _, test := range stringComparisonTests {
		t.Run(fmt.Sprintf("%s %s", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedGreaterThanEqual, pie.GreaterThanEqualString(test.b)(test.a))
		})
	}
}

func TestLessThanString(t *testing.T) {
	for _, test := range stringComparisonTests {
		t.Run(fmt.Sprintf("%s %s", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedLessThan, pie.LessThanString(test.b)(test.a))
		})
	}
}

func TestLessThanEqualString(t *testing.T) {
	for _, test := range stringComparisonTests {
		t.Run(fmt.Sprintf("%s %s", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expectedLessThanEqual, pie.LessThanEqualString(test.b)(test.a))
		})
	}
}
