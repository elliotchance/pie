package pie

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var carPointerA = &car{"a", "green"}
var carPointerB = &car{"b", "blue"}
var carPointerC = &car{"c", "gray"}
var carPointerEmpty = &car{}

var carPointersContainsTests = []struct {
	ss       carPointers
	contains *car
	expected bool
}{
	{nil, carPointerA, false},
	{nil, carPointerEmpty, false},
	{nil, nil, false},
	{carPointers{carPointerA, carPointerB, carPointerC}, carPointerA, true},
	{carPointers{carPointerA, carPointerB, carPointerC}, carPointerB, true},
	{carPointers{carPointerA, carPointerB, carPointerC}, carPointerC, true},
	{carPointers{carPointerA, carPointerB, carPointerC}, &car{"a", "green"}, false},
	{carPointers{carPointerA, carPointerB, carPointerC}, &car{"A", ""}, false},
	{carPointers{carPointerA, carPointerB, carPointerC}, &car{}, false},
	{carPointers{carPointerA, carPointerB, carPointerC}, &car{"d", ""}, false},
	{carPointers{carPointerA, carPointerEmpty, carPointerC}, carPointerEmpty, true},
	{carPointers{carPointerA, nil, carPointerC}, nil, true},
	{carPointers{carPointerA, carPointerEmpty, carPointerC}, nil, false},
}

func TestCarPointers_Contains(t *testing.T) {
	for _, test := range carPointersContainsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Contains(test.contains))
		})
	}
}

var carPointersSelectTests = []struct {
	ss                carPointers
	condition         func(*car) bool
	expectedSelect    carPointers
	expectedUnselect  carPointers
	expectedTransform carPointers
}{
	{
		nil,
		func(s *car) bool {
			return s.Name == ""
		},
		nil,
		nil,
		nil,
	},
	{
		carPointers{carPointerA, carPointerB, carPointerC},
		func(s *car) bool {
			return s.Name != "b"
		},
		carPointers{carPointerA, carPointerC},
		carPointers{carPointerB},
		carPointers{&car{"A", "green"}, &car{"B", "blue"}, &car{"C", "gray"}},
	},
}

func TestCarPointers_Select(t *testing.T) {
	for _, test := range carPointersSelectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.expectedSelect, test.ss.Select(test.condition))
		})
	}
}

func TestCarPointers_Unselect(t *testing.T) {
	for _, test := range carPointersSelectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.expectedUnselect, test.ss.Unselect(test.condition))
		})
	}
}

func TestCarPointers_Transform(t *testing.T) {
	for _, test := range carPointersSelectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.expectedTransform, test.ss.Transform(func(c *car) *car {
				return &car{
					Name:  strings.ToUpper(c.Name),
					Color: c.Color,
				}
			}))
		})
	}
}

var carPointersFirstAndLastTests = []struct {
	ss             carPointers
	first, firstOr *car
	last, lastOr   *car
}{
	{
		nil,
		&car{},
		&car{"default1", "unknown"},
		&car{},
		&car{"default2", "unknown"},
	},
	{
		carPointers{&car{"foo", "red"}},
		&car{"foo", "red"},
		&car{"foo", "red"},
		&car{"foo", "red"},
		&car{"foo", "red"},
	},
	{
		carPointers{carPointerA, carPointerB},
		carPointerA,
		carPointerA,
		carPointerB,
		carPointerB,
	},
	{
		carPointers{carPointerA, carPointerB, carPointerC},
		carPointerA,
		carPointerA,
		carPointerC,
		carPointerC,
	},
}

func TestCarPointers_FirstOr(t *testing.T) {
	for _, test := range carPointersFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.firstOr, test.ss.FirstOr(&car{"default1", "unknown"}))
		})
	}
}

func TestCarPointers_LastOr(t *testing.T) {
	for _, test := range carPointersFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.lastOr, test.ss.LastOr(&car{"default2", "unknown"}))
		})
	}
}

func TestCarPointers_First(t *testing.T) {
	for _, test := range carPointersFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.first, test.ss.First())
		})
	}
}

func TestCarPointers_Last(t *testing.T) {
	for _, test := range carPointersFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.last, test.ss.Last())
		})
	}
}

var carPointersStatsTests = []struct {
	ss       carPointers
	min, max *car
	len      int
}{
	{
		nil,
		&car{},
		&car{},
		0,
	},
	{
		carPointers{},
		&car{},
		&car{},
		0,
	},
	{
		carPointers{&car{"foo", "red"}},
		&car{"foo", "red"},
		&car{"foo", "red"},
		1,
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}, &car{"qux", "cyan"}, &car{"foo", "red"}},
		&car{"Baz", "black"},
		&car{"qux", "cyan"},
		4,
	},
}

func TestCarPointers_Len(t *testing.T) {
	for _, test := range carPointersStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.len, test.ss.Len())
		})
	}
}

var carPointersJSONTests = []struct {
	ss         carPointers
	jsonString string
}{
	{
		nil,
		`[]`, // Instead of null.
	},
	{
		carPointers{},
		`[]`,
	},
	{
		carPointers{&car{"foo", "red"}},
		`[{"Name":"foo","Color":"red"}]`,
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}, &car{"qux", "cyan"}, &car{"foo", "red"}},
		`[{"Name":"bar","Color":"yellow"},{"Name":"Baz","Color":"black"},{"Name":"qux","Color":"cyan"},{"Name":"foo","Color":"red"}]`,
	},
}

func TestCarPointers_JSONString(t *testing.T) {
	for _, test := range carPointersJSONTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.jsonString, test.ss.JSONString())
		})
	}
}

var carPointersSortTests = []struct {
	ss        carPointers
	sorted    carPointers
	reversed  carPointers
	areSorted bool
}{
	{
		nil,
		nil,
		nil,
		true,
	},
	{
		carPointers{},
		carPointers{},
		carPointers{},
		true,
	},
	{
		carPointers{&car{"foo", "red"}},
		carPointers{&car{"foo", "red"}},
		carPointers{&car{"foo", "red"}},
		true,
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}, &car{"foo", "red"}},
		carPointers{&car{"Baz", "black"}, &car{"bar", "yellow"}, &car{"foo", "red"}},
		carPointers{&car{"foo", "red"}, &car{"Baz", "black"}, &car{"bar", "yellow"}},
		false,
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}, &car{"qux", "cyan"}, &car{"foo", "red"}},
		carPointers{&car{"Baz", "black"}, &car{"bar", "yellow"}, &car{"foo", "red"}, &car{"qux", "cyan"}},
		carPointers{&car{"foo", "red"}, &car{"qux", "cyan"}, &car{"Baz", "black"}, &car{"bar", "yellow"}},
		false,
	},
	{
		carPointers{&car{"Baz", "black"}, &car{"bar", "yellow"}},
		carPointers{&car{"Baz", "black"}, &car{"bar", "yellow"}},
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}},
		true,
	},
}

func TestCarPointers_Reverse(t *testing.T) {
	for _, test := range carPointersSortTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.reversed, test.ss.Reverse())
		})
	}
}

var stringsToStringsTests = []struct {
	ss        carPointers
	transform func(*car) string
	expected  Strings
}{
	{
		nil,
		func(s *car) string {
			return "foo"
		},
		nil,
	},
	{
		carPointers{},
		func(s *car) string {
			return fmt.Sprintf("%s!", s.Name)
		},
		nil,
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}, &car{"foo", "red"}},
		func(s *car) string {
			return fmt.Sprintf("%s!", s.Color)
		},
		Strings{"yellow!", "black!", "red!"},
	},
}

func TestCarPointers_ToStrings(t *testing.T) {
	for _, test := range stringsToStringsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.ToStrings(test.transform))
		})
	}
}

func TestCarPointers_Append(t *testing.T) {
	assert.Equal(t,
		(carPointers)(nil).Append(),
		(carPointers)(nil),
	)

	assert.Equal(t,
		(carPointers)(nil).Append(&car{"bar", "yellow"}),
		carPointers{&car{"bar", "yellow"}},
	)

	assert.Equal(t,
		(carPointers)(nil).Append(&car{"bar", "yellow"}, &car{"Baz", "black"}),
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}},
	)

	assert.Equal(t,
		carPointers{&car{"bar", "yellow"}}.Append(&car{"Baz", "black"}),
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}},
	)

	assert.Equal(t,
		carPointers{&car{"bar", "yellow"}}.Append(&car{"Baz", "black"}, &car{"foo", "red"}),
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}, &car{"foo", "red"}},
	)
}

func TestCarPointers_Extend(t *testing.T) {
	assert.Equal(t,
		(carPointers)(nil).Extend(),
		(carPointers)(nil),
	)

	assert.Equal(t,
		(carPointers)(nil).Extend(carPointers{&car{"bar", "yellow"}}),
		carPointers{&car{"bar", "yellow"}},
	)

	assert.Equal(t,
		(carPointers)(nil).Extend(carPointers{&car{"bar", "yellow"}}, carPointers{&car{"Baz", "black"}}),
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}},
	)

	assert.Equal(t,
		carPointers{&car{"bar", "yellow"}}.Extend(carPointers{&car{"Baz", "black"}}),
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}},
	)

	assert.Equal(t,
		carPointers{&car{"bar", "yellow"}}.Extend(carPointers{&car{"Baz", "black"}, &car{"foo", "red"}}),
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}, &car{"foo", "red"}},
	)
}

func TestCarPointers_All(t *testing.T) {
	assert.True(t,
		(carPointers)(nil).All(func(value *car) bool {
			return false
		}),
	)

	assert.True(t,
		(carPointers)(nil).All(func(value *car) bool {
			return false
		}),
	)

	// None
	assert.False(t,
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}}.All(func(value *car) bool {
			return value.Color == "green"
		}),
	)

	// Some
	assert.False(t,
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}}.All(func(value *car) bool {
			return value.Color == "yellow"
		}),
	)

	// All
	assert.True(t,
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}}.All(func(value *car) bool {
			return len(value.Name) > 0
		}),
	)
}

func TestCarPointers_Any(t *testing.T) {
	assert.False(t,
		(carPointers)(nil).Any(func(value *car) bool {
			return true
		}),
	)

	assert.False(t,
		(carPointers)(nil).Any(func(value *car) bool {
			return true
		}),
	)

	// None
	assert.False(t,
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}}.Any(func(value *car) bool {
			return value.Color == "green"
		}),
	)

	// Some
	assert.True(t,
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}}.Any(func(value *car) bool {
			return value.Color == "yellow"
		}),
	)

	// All
	assert.True(t,
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}}.Any(func(value *car) bool {
			return len(value.Name) > 0
		}),
	)
}

var carPointersShuffleTests = []struct {
	ss       carPointers
	expected carPointers
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
		carPointers{},
		carPointers{},
		rand.NewSource(0),
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}, &car{"foo", "red"}},
		carPointers{&car{"Baz", "black"}, &car{"bar", "yellow"}, &car{"foo", "red"}},
		rand.NewSource(0),
	},
	{
		carPointers{&car{"bar", "yellow"}},
		carPointers{&car{"bar", "yellow"}},
		rand.NewSource(0),
	},
}

func TestCarPointers_Shuffle(t *testing.T) {
	for _, test := range carPointersShuffleTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Shuffle(test.source))
		})
	}
}

var carPointersTopAndBottomTests = []struct {
	ss     carPointers
	n      int
	top    carPointers
	bottom carPointers
}{
	{
		nil,
		1,
		nil,
		nil,
	},
	{
		carPointers{},
		1,
		nil,
		nil,
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}},
		1,
		carPointers{&car{"bar", "yellow"}},
		carPointers{&car{"Baz", "black"}},
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}},
		3,
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}},
		carPointers{&car{"Baz", "black"}, &car{"bar", "yellow"}},
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}},
		0,
		nil,
		nil,
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}},
		-1,
		nil,
		nil,
	},
}

func TestCarPointers_Top(t *testing.T) {
	for _, test := range carPointersTopAndBottomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.top, test.ss.Top(test.n))
		})
	}
}

func TestCarPointers_Bottom(t *testing.T) {
	for _, test := range carPointersTopAndBottomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.bottom, test.ss.Bottom(test.n))
		})
	}
}

func TestCarPointers_Each(t *testing.T) {
	var names []string

	names = []string{}
	carPointers(nil).Each(func(car *car) {
		names = append(names, car.Name)
	})
	assert.Equal(t, []string{}, names)

	names = []string{}
	carPointers{}.Each(func(car *car) {
		names = append(names, car.Name)
	})
	assert.Equal(t, []string{}, names)

	names = []string{}
	carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}}.Each(func(car *car) {
		names = append(names, car.Name)
	})
	assert.Equal(t, []string{"bar", "Baz"}, names)
}

var carPointersRandomTests = []struct {
	ss       carPointers
	expected *car
	source   rand.Source
}{
	{
		nil,
		&car{},
		nil,
	},
	{
		nil,
		&car{},
		rand.NewSource(0),
	},
	{
		carPointers{},
		&car{},
		rand.NewSource(0),
	},
	{
		carPointers{&car{"bar", "yellow"}, &car{"Baz", "black"}, &car{"foo", "red"}},
		&car{"foo", "red"},
		rand.NewSource(0),
	},
	{
		carPointers{&car{"bar", "yellow"}},
		&car{"bar", "yellow"},
		rand.NewSource(0),
	},
}

func TestCarPointers_Random(t *testing.T) {
	for _, test := range carPointersRandomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCarPointers(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Random(test.source))
		})
	}
}
