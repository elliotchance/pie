package pie_test

import (
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrings_Contains(t *testing.T) {
	for _, test := range []struct {
		ss       pie.Strings
		contains string
		expected bool
	}{
		{nil, "a", false},
		{nil, "", false},
		{pie.Strings{"a", "b", "c"}, "a", true},
		{pie.Strings{"a", "b", "c"}, "b", true},
		{pie.Strings{"a", "b", "c"}, "c", true},
		{pie.Strings{"a", "b", "c"}, "A", false},
		{pie.Strings{"a", "b", "c"}, "", false},
		{pie.Strings{"a", "b", "c"}, "d", false},
		{pie.Strings{"a", "", "c"}, "", true},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, test.ss.Contains(test.contains))
		})
	}
}

var onlyAndWithoutTests = []struct {
	ss                pie.Strings
	condition         func(string) bool
	expectedOnly      pie.Strings
	expectedWithout   pie.Strings
	expectedTransform pie.Strings
}{
	{
		nil,
		func(s string) bool {
			return s == ""
		},
		nil,
		nil,
		nil,
	},
	{
		pie.Strings{"a", "b", "c"},
		func(s string) bool {
			return s != "b"
		},
		pie.Strings{"a", "c"},
		pie.Strings{"b"},
		pie.Strings{"A", "B", "C"},
	},
}

func TestStrings_Only(t *testing.T) {
	for _, test := range onlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedOnly, test.ss.Only(test.condition))
		})
	}
}

func TestStrings_Without(t *testing.T) {
	for _, test := range onlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedWithout, test.ss.Without(test.condition))
		})
	}
}

func TestStrings_Transform(t *testing.T) {
	for _, test := range onlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedTransform, test.ss.Transform(pie.ToUpper()))
		})
	}
}

var firstAndLastTests = []struct {
	ss          pie.Strings
	first, last string
}{
	{
		nil,
		"default",
		"default",
	},
	{
		pie.Strings{"foo"},
		"foo",
		"foo",
	},
	{
		pie.Strings{"a", "b"},
		"a",
		"b",
	},
	{
		pie.Strings{"a", "b", "c"},
		"a",
		"c",
	},
}

func TestStrings_First(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.first, test.ss.First("default"))
		})
	}
}

func TestStrings_Last(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.last, test.ss.Last("default"))
		})
	}
}
