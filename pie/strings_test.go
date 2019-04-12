package pie

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var stringsContainsTests = []struct {
	ss       Strings
	contains string
	expected bool
}{
	{nil, "a", false},
	{nil, "", false},
	{Strings{"a", "b", "c"}, "a", true},
	{Strings{"a", "b", "c"}, "b", true},
	{Strings{"a", "b", "c"}, "c", true},
	{Strings{"a", "b", "c"}, "A", false},
	{Strings{"a", "b", "c"}, "", false},
	{Strings{"a", "b", "c"}, "d", false},
	{Strings{"a", "", "c"}, "", true},
}

func TestStrings_Contains(t *testing.T) {
	for _, test := range stringsContainsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Contains(test.contains))
		})
	}
}

var stringsSelectTests = []struct {
	ss                Strings
	condition         func(string) bool
	expectedSelect    Strings
	expectedUnselect  Strings
	expectedTransform Strings
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
		Strings{"a", "b", "c"},
		func(s string) bool {
			return s != "b"
		},
		Strings{"a", "c"},
		Strings{"b"},
		Strings{"A", "B", "C"},
	},
}

func TestStrings_Select(t *testing.T) {
	for _, test := range stringsSelectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expectedSelect, test.ss.Select(test.condition))
		})
	}
}

func TestStrings_Unselect(t *testing.T) {
	for _, test := range stringsSelectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expectedUnselect, test.ss.Unselect(test.condition))
		})
	}
}

func TestStrings_Transform(t *testing.T) {
	for _, test := range stringsSelectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expectedTransform, test.ss.Transform(strings.ToUpper))
		})
	}
}

var firstAndLastTests = []struct {
	ss             Strings
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
		Strings{"foo"},
		"foo",
		"foo",
		"foo",
		"foo",
	},
	{
		Strings{"a", "b"},
		"a",
		"a",
		"b",
		"b",
	},
	{
		Strings{"a", "b", "c"},
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
	ss       Strings
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
			assert.Equal(t, test.min, Strings(test.ss).Min())
		})
	}
}

func TestStrings_Max(t *testing.T) {
	for _, test := range stringsStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.max, Strings(test.ss).Max())
		})
	}
}

func TestStrings_Len(t *testing.T) {
	for _, test := range stringsStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.len, Strings(test.ss).Len())
		})
	}
}

var stringsJSONTests = []struct {
	ss         Strings
	jsonString string
}{
	{
		nil,
		`[]`, // Instead of null.
	},
	{
		Strings{},
		`[]`,
	},
	{
		Strings{"foo"},
		`["foo"]`,
	},
	{
		Strings{"bar", "Baz", "qux", "foo"},
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
	ss        Strings
	sorted    Strings
	reversed  Strings
	areSorted bool
}{
	{
		nil,
		nil,
		nil,
		true,
	},
	{
		Strings{},
		Strings{},
		Strings{},
		true,
	},
	{
		Strings{"foo"},
		Strings{"foo"},
		Strings{"foo"},
		true,
	},
	{
		Strings{"bar", "Baz", "foo"},
		Strings{"Baz", "bar", "foo"},
		Strings{"foo", "Baz", "bar"},
		false,
	},
	{
		Strings{"bar", "Baz", "qux", "foo"},
		Strings{"Baz", "bar", "foo", "qux"},
		Strings{"foo", "qux", "Baz", "bar"},
		false,
	},
	{
		Strings{"Baz", "bar"},
		Strings{"Baz", "bar"},
		Strings{"bar", "Baz"},
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
	ss        Strings
	unique    Strings
	areUnique bool
}{
	{
		nil,
		nil,
		true,
	},
	{
		Strings{},
		Strings{},
		true,
	},
	{
		Strings{"baz"},
		Strings{"baz"},
		true,
	},
	{
		Strings{"foo", "bar", "foo"},
		Strings{"bar", "foo"},
		false,
	},
	{
		Strings{"foo", "bar", "qux", "baz"},
		Strings{"bar", "baz", "foo", "qux"},
		true,
	},
}

func TestStrings_Unique(t *testing.T) {
	for _, test := range stringsUniqueTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()

			// We have to sort the unique slice because it is always returned in
			// random order.
			assert.Equal(t, test.unique, test.ss.Unique().Sort())
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

var carPointersToStringsTests = []struct {
	ss        Strings
	transform func(string) string
	expected  Strings
}{
	{
		nil,
		func(s string) string {
			return "foo"
		},
		nil,
	},
	{
		Strings{},
		func(s string) string {
			return fmt.Sprintf("%s!", s)
		},
		nil,
	},
	{
		Strings{"6.2", "7.2", "8.2"},
		func(s string) string {
			return fmt.Sprintf("%s!", s)
		},
		Strings{"6.2!", "7.2!", "8.2!"},
	},
}

func TestStrings_ToStrings(t *testing.T) {
	for _, test := range carPointersToStringsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.ToStrings(test.transform))
		})
	}
}
