package pie_test

import (
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var stringsContainsTests = []struct {
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
}

func TestStrings_Contains(t *testing.T) {
	for _, test := range stringsContainsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Contains(test.contains))
		})
	}
}

var stringsOnlyAndWithoutTests = []struct {
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
	for _, test := range stringsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expectedOnly, test.ss.Only(test.condition))
		})
	}
}

func TestStrings_Without(t *testing.T) {
	for _, test := range stringsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expectedWithout, test.ss.Without(test.condition))
		})
	}
}

func TestStrings_Transform(t *testing.T) {
	for _, test := range stringsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expectedTransform, test.ss.Transform(strings.ToUpper))
		})
	}
}

var firstAndLastTests = []struct {
	ss             pie.Strings
	first, firstOr string
	last, lastOr   string
}{
	{
		nil,
		"",
		"default1",
		"",
		"default2",
	},
	{
		pie.Strings{"foo"},
		"foo",
		"foo",
		"foo",
		"foo",
	},
	{
		pie.Strings{"a", "b"},
		"a",
		"a",
		"b",
		"b",
	},
	{
		pie.Strings{"a", "b", "c"},
		"a",
		"a",
		"c",
		"c",
	},
}

func TestStrings_FirstOr(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.firstOr, test.ss.FirstOr("default1"))
		})
	}
}

func TestStrings_LastOr(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.lastOr, test.ss.LastOr("default2"))
		})
	}
}

func TestStrings_First(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.first, test.ss.First())
		})
	}
}

func TestStrings_Last(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.last, test.ss.Last())
		})
	}
}

var stringsStatsTests = []struct {
	ss       pie.Strings
	min, max string
	len      int
}{
	{
		nil,
		"",
		"",
		0,
	},
	{
		[]string{},
		"",
		"",
		0,
	},
	{
		[]string{"foo"},
		"foo",
		"foo",
		1,
	},
	{
		[]string{"bar", "Baz", "qux", "foo"},
		"Baz",
		"qux",
		4,
	},
}

func TestStrings_Min(t *testing.T) {
	for _, test := range stringsStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.min, pie.Strings(test.ss).Min())
		})
	}
}

func TestStrings_Max(t *testing.T) {
	for _, test := range stringsStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.max, pie.Strings(test.ss).Max())
		})
	}
}

func TestStrings_Len(t *testing.T) {
	for _, test := range stringsStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.len, pie.Strings(test.ss).Len())
		})
	}
}

var stringsJSONTests = []struct {
	ss         pie.Strings
	jsonString string
}{
	{
		nil,
		`[]`, // Instead of null.
	},
	{
		pie.Strings{},
		`[]`,
	},
	{
		pie.Strings{"foo"},
		`["foo"]`,
	},
	{
		pie.Strings{"bar", "Baz", "qux", "foo"},
		`["bar","Baz","qux","foo"]`,
	},
}

func TestStrings_JSONString(t *testing.T) {
	for _, test := range stringsJSONTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.jsonString, test.ss.JSONString())
		})
	}
}

var stringsSortTests = []struct {
	ss        pie.Strings
	sorted    pie.Strings
	reversed  pie.Strings
	areSorted bool
}{
	{
		nil,
		nil,
		nil,
		true,
	},
	{
		pie.Strings{},
		pie.Strings{},
		pie.Strings{},
		true,
	},
	{
		pie.Strings{"foo"},
		pie.Strings{"foo"},
		pie.Strings{"foo"},
		true,
	},
	{
		pie.Strings{"bar", "Baz", "foo"},
		pie.Strings{"Baz", "bar", "foo"},
		pie.Strings{"foo", "Baz", "bar"},
		false,
	},
	{
		pie.Strings{"bar", "Baz", "qux", "foo"},
		pie.Strings{"Baz", "bar", "foo", "qux"},
		pie.Strings{"foo", "qux", "Baz", "bar"},
		false,
	},
	{
		pie.Strings{"Baz", "bar"},
		pie.Strings{"Baz", "bar"},
		pie.Strings{"bar", "Baz"},
		true,
	},
}

func TestStrings_Sort(t *testing.T) {
	for _, test := range stringsSortTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.sorted, test.ss.Sort())
		})
	}
}

func TestStrings_Reverse(t *testing.T) {
	for _, test := range stringsSortTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.reversed, test.ss.Reverse())
		})
	}
}

func TestStrings_AreSorted(t *testing.T) {
	for _, test := range stringsSortTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.areSorted, test.ss.AreSorted())
		})
	}
}

var stringsUniqueTests = []struct {
	ss        pie.Strings
	unique    pie.Strings
	areUnique bool
}{
	{
		nil,
		nil,
		true,
	},
	{
		pie.Strings{},
		pie.Strings{},
		true,
	},
	{
		pie.Strings{"foo"},
		pie.Strings{"foo"},
		true,
	},
	{
		pie.Strings{"bar", "Baz", "foo"},
		pie.Strings{"Baz", "bar", "foo"},
		true,
	},
	{
		pie.Strings{"bar", "Baz", "qux", "bar"},
		pie.Strings{"Baz", "bar", "qux"},
		false,
	},
}

func TestStrings_Unique(t *testing.T) {
	for _, test := range stringsUniqueTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.unique, test.ss.Unique())
		})
	}
}

func TestStrings_AreUnique(t *testing.T) {
	for _, test := range stringsUniqueTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.areUnique, test.ss.AreUnique())
		})
	}
}
