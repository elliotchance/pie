package pie

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/elliotchance/testify-stats/assert"
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

var stringsFilterTests = []struct {
	ss                Strings
	condition         func(string) bool
	expectedFilter    Strings
	expectedFilterNot Strings
	expectedMap       Strings
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

func TestStrings_Filter(t *testing.T) {
	for _, test := range stringsFilterTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expectedFilter, test.ss.Filter(test.condition))
		})
	}
}

func TestStrings_FilterNot(t *testing.T) {
	for _, test := range stringsFilterTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expectedFilterNot, test.ss.FilterNot(test.condition))
		})
	}
}

func TestStrings_Map(t *testing.T) {
	for _, test := range stringsFilterTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expectedMap, test.ss.Map(strings.ToUpper))
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

func TestStrings_Join(t *testing.T) {
	assert.Equal(t, "", Strings(nil).Join("a"))
	assert.Equal(t, "", Strings{}.Join("a"))
	assert.Equal(t, "foo--bar", Strings{"foo", "", "bar"}.Join("-"))
}

func TestStrings_Append(t *testing.T) {
	assert.Equal(t,
		Strings{}.Append(),
		Strings{},
	)

	assert.Equal(t,
		Strings{}.Append("bar"),
		Strings{"bar"},
	)

	assert.Equal(t,
		Strings{}.Append("bar", "Baz"),
		Strings{"bar", "Baz"},
	)

	assert.Equal(t,
		Strings{"bar"}.Append("Baz"),
		Strings{"bar", "Baz"},
	)

	assert.Equal(t,
		Strings{"bar"}.Append("Baz", "foo"),
		Strings{"bar", "Baz", "foo"},
	)
}

func TestStrings_Extend(t *testing.T) {
	assert.Equal(t,
		Strings{}.Extend(),
		Strings{},
	)

	assert.Equal(t,
		Strings{}.Extend([]string{"bar"}),
		Strings{"bar"},
	)

	assert.Equal(t,
		Strings{}.Extend([]string{"bar"}, []string{"Baz"}),
		Strings{"bar", "Baz"},
	)

	assert.Equal(t,
		Strings{"bar"}.Extend([]string{"Baz"}),
		Strings{"bar", "Baz"},
	)

	assert.Equal(t,
		Strings{"bar"}.Extend([]string{"Baz", "foo"}),
		Strings{"bar", "Baz", "foo"},
	)
}

func TestStrings_All(t *testing.T) {
	assert.True(t,
		Strings{}.All(func(value string) bool {
			return false
		}),
	)

	assert.True(t,
		Strings{}.All(func(value string) bool {
			return false
		}),
	)

	// None
	assert.False(t,
		Strings{"foo", "bar"}.All(func(value string) bool {
			return value == "baz"
		}),
	)

	// Some
	assert.False(t,
		Strings{"foo", "bar"}.All(func(value string) bool {
			return value == "foo"
		}),
	)

	// All
	assert.True(t,
		Strings{"foo", "bar"}.All(func(value string) bool {
			return len(value) > 0
		}),
	)
}

func TestStrings_Any(t *testing.T) {
	assert.False(t,
		Strings{}.Any(func(value string) bool {
			return true
		}),
	)

	assert.False(t,
		Strings{}.Any(func(value string) bool {
			return true
		}),
	)

	// None
	assert.False(t,
		Strings{"foo", "bar"}.Any(func(value string) bool {
			return value == "baz"
		}),
	)

	// Some
	assert.True(t,
		Strings{"foo", "bar"}.Any(func(value string) bool {
			return value == "foo"
		}),
	)

	// All
	assert.True(t,
		Strings{"foo", "bar"}.Any(func(value string) bool {
			return len(value) > 0
		}),
	)
}

var stringsShuffleTests = []struct {
	ss       Strings
	expected Strings
	source   rand.Source
}{
	{
		nil,
		nil,
		nil,
	},
	{
		nil,
		nil,
		rand.NewSource(0),
	},
	{
		Strings{},
		Strings{},
		rand.NewSource(0),
	},
	{
		Strings{"foo", "bar", "baz"},
		Strings{"bar", "foo", "baz"},
		rand.NewSource(0),
	},
	{
		Strings{"foo"},
		Strings{"foo"},
		rand.NewSource(0),
	},
}

func TestStrings_Shuffle(t *testing.T) {
	for _, test := range stringsShuffleTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Shuffle(test.source))
		})
	}
}

var stringsTopAndBottomTests = []struct {
	ss     Strings
	n      int
	top    Strings
	bottom Strings
}{
	{
		nil,
		1,
		nil,
		nil,
	},
	{
		Strings{},
		1,
		nil,
		nil,
	},
	{
		Strings{"foo", "bar"},
		1,
		Strings{"foo"},
		Strings{"bar"},
	},
	{
		Strings{"foo", "bar"},
		3,
		Strings{"foo", "bar"},
		Strings{"bar", "foo"},
	},
	{
		Strings{"foo", "bar"},
		0,
		nil,
		nil,
	},
	{
		Strings{"foo", "bar"},
		-1,
		nil,
		nil,
	},
}

func TestStrings_Top(t *testing.T) {
	for _, test := range stringsTopAndBottomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.top, test.ss.Top(test.n))
		})
	}
}

func TestStrings_Bottom(t *testing.T) {
	for _, test := range stringsTopAndBottomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.bottom, test.ss.Bottom(test.n))
		})
	}
}

func TestStrings_Each(t *testing.T) {
	var values []string

	values = []string{}
	Strings{}.Each(func(value string) {
		values = append(values, value)
	})
	assert.Equal(t, []string{}, values)

	values = []string{}
	Strings{"baz", "qux"}.Each(func(value string) {
		values = append(values, value)
	})
	assert.Equal(t, []string{"baz", "qux"}, values)
}

var stringsRandomTests = []struct {
	ss       Strings
	expected string
	source   rand.Source
}{
	{
		nil,
		"",
		nil,
	},
	{
		nil,
		"",
		rand.NewSource(0),
	},
	{
		Strings{},
		"",
		rand.NewSource(0),
	},
	{
		Strings{"foo", "bar", "baz"},
		"baz",
		rand.NewSource(1),
	},
	{
		Strings{"foo", "bar", "baz"},
		"foo",
		rand.NewSource(0),
	},
	{
		Strings{"foo"},
		"foo",
		rand.NewSource(0),
	},
}

func TestStrings_Random(t *testing.T) {
	for _, test := range stringsRandomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Random(test.source))
		})
	}
}

var stringsSendTests = []struct {
	ss            Strings
	recieveDelay  time.Duration
	canceledDelay time.Duration
	expected      Strings
}{
	{
		nil,
		0,
		0,
		nil,
	},
	{
		Strings{"foo", "bar"},
		0,
		0,
		Strings{"foo", "bar"},
	},
	{
		Strings{"foo", "bar"},
		time.Millisecond * 30,
		time.Millisecond * 10,
		Strings{"foo"},
	},
	{
		Strings{"foo", "bar"},
		time.Millisecond * 3,
		time.Millisecond * 10,
		Strings{"foo", "bar"},
	},
}

func TestStrings_Send(t *testing.T) {
	for _, test := range stringsSendTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableStrings(t, &test.ss)()
			ch := make(chan string)
			actual := getStringsFromChan(ch, test.recieveDelay)
			ctx := createContextByDelay(test.canceledDelay)

			actualSended := test.ss.Send(ctx, ch)
			close(ch)

			assert.Equal(t, test.expected, actualSended)
			assert.Equal(t, test.expected, actual())
		})
	}
}
