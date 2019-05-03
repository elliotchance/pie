package pie

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var carsContainsTests = []struct {
	ss       cars
	contains car
	expected bool
}{
	{nil, car{"a", "green"}, false},
	{nil, car{}, false},
	{cars{car{"a", "green"}, car{"b", "blue"}, car{"c", "gray"}}, car{"a", "green"}, true},
	{cars{car{"a", "green"}, car{"b", "blue"}, car{"c", "gray"}}, car{"b", "blue"}, true},
	{cars{car{"a", "green"}, car{"b", "blue"}, car{"c", "gray"}}, car{"c", "gray"}, true},
	{cars{car{"a", "green"}, car{"b", "blue"}, car{"c", "gray"}}, car{"A", ""}, false},
	{cars{car{"a", "green"}, car{"b", "blue"}, car{"c", "gray"}}, car{}, false},
	{cars{car{"a", "green"}, car{"b", "blue"}, car{"c", "gray"}}, car{"d", ""}, false},
	{cars{car{"a", "green"}, car{}, car{"c", "gray"}}, car{}, true},
}

func TestCars_Contains(t *testing.T) {
	for _, test := range carsContainsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Contains(test.contains))
		})
	}
}

var carsSelectTests = []struct {
	ss                cars
	condition         func(car) bool
	expectedSelect    cars
	expectedUnselect  cars
	expectedTransform cars
}{
	{
		nil,
		func(s car) bool {
			return s.Name == ""
		},
		nil,
		nil,
		nil,
	},
	{
		cars{car{"a", "green"}, car{"b", "blue"}, car{"c", "gray"}},
		func(s car) bool {
			return s.Name != "b"
		},
		cars{car{"a", "green"}, car{"c", "gray"}},
		cars{car{"b", "blue"}},
		cars{car{"A", "green"}, car{"B", "blue"}, car{"C", "gray"}},
	},
}

func TestCars_Select(t *testing.T) {
	for _, test := range carsSelectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.expectedSelect, test.ss.Select(test.condition))
		})
	}
}

func TestCars_Unselect(t *testing.T) {
	for _, test := range carsSelectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.expectedUnselect, test.ss.Unselect(test.condition))
		})
	}
}

func TestCars_Transform(t *testing.T) {
	for _, test := range carsSelectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.expectedTransform, test.ss.Transform(func(car car) car {
				car.Name = strings.ToUpper(car.Name)

				return car
			}))
		})
	}
}

var carsFirstAndLastTests = []struct {
	ss             cars
	first, firstOr car
	last, lastOr   car
}{
	{
		nil,
		car{},
		car{"default1", "unknown"},
		car{},
		car{"default2", "unknown"},
	},
	{
		cars{car{"foo", "red"}},
		car{"foo", "red"},
		car{"foo", "red"},
		car{"foo", "red"},
		car{"foo", "red"},
	},
	{
		cars{car{"a", "green"}, car{"b", "blue"}},
		car{"a", "green"},
		car{"a", "green"},
		car{"b", "blue"},
		car{"b", "blue"},
	},
	{
		cars{car{"a", "green"}, car{"b", "blue"}, car{"c", "gray"}},
		car{"a", "green"},
		car{"a", "green"},
		car{"c", "gray"},
		car{"c", "gray"},
	},
}

func TestCars_FirstOr(t *testing.T) {
	for _, test := range carsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.firstOr, test.ss.FirstOr(car{"default1", "unknown"}))
		})
	}
}

func TestCars_LastOr(t *testing.T) {
	for _, test := range carsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.lastOr, test.ss.LastOr(car{"default2", "unknown"}))
		})
	}
}

func TestCars_First(t *testing.T) {
	for _, test := range carsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.first, test.ss.First())
		})
	}
}

func TestCars_Last(t *testing.T) {
	for _, test := range carsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.last, test.ss.Last())
		})
	}
}

var carsStatsTests = []struct {
	ss       cars
	min, max car
	len      int
}{
	{
		nil,
		car{},
		car{},
		0,
	},
	{
		cars{},
		car{},
		car{},
		0,
	},
	{
		cars{car{"foo", "red"}},
		car{"foo", "red"},
		car{"foo", "red"},
		1,
	},
	{
		cars{car{"bar", "yellow"}, car{"Baz", "black"}, car{"qux", "cyan"}, car{"foo", "red"}},
		car{"Baz", "black"},
		car{"qux", "cyan"},
		4,
	},
}

func TestCars_Len(t *testing.T) {
	for _, test := range carsStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.len, cars(test.ss).Len())
		})
	}
}

var carsJSONTests = []struct {
	ss         cars
	jsonString string
}{
	{
		nil,
		`[]`, // Instead of null.
	},
	{
		cars{},
		`[]`,
	},
	{
		cars{car{"foo", "red"}},
		`[{"Name":"foo","Color":"red"}]`,
	},
	{
		cars{car{"bar", "yellow"}, car{"Baz", "black"}, car{"qux", "cyan"}, car{"foo", "red"}},
		`[{"Name":"bar","Color":"yellow"},{"Name":"Baz","Color":"black"},{"Name":"qux","Color":"cyan"},{"Name":"foo","Color":"red"}]`,
	},
}

func TestCars_JSONString(t *testing.T) {
	for _, test := range carsJSONTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.jsonString, test.ss.JSONString())
		})
	}
}

var carsSortTests = []struct {
	ss        cars
	sorted    cars
	reversed  cars
	areSorted bool
}{
	{
		nil,
		nil,
		nil,
		true,
	},
	{
		cars{},
		cars{},
		cars{},
		true,
	},
	{
		cars{car{"foo", "red"}},
		cars{car{"foo", "red"}},
		cars{car{"foo", "red"}},
		true,
	},
	{
		cars{car{"bar", "yellow"}, car{"Baz", "black"}, car{"foo", "red"}},
		cars{car{"Baz", "black"}, car{"bar", "yellow"}, car{"foo", "red"}},
		cars{car{"foo", "red"}, car{"Baz", "black"}, car{"bar", "yellow"}},
		false,
	},
	{
		cars{car{"bar", "yellow"}, car{"Baz", "black"}, car{"qux", "cyan"}, car{"foo", "red"}},
		cars{car{"Baz", "black"}, car{"bar", "yellow"}, car{"foo", "red"}, car{"qux", "cyan"}},
		cars{car{"foo", "red"}, car{"qux", "cyan"}, car{"Baz", "black"}, car{"bar", "yellow"}},
		false,
	},
	{
		cars{car{"Baz", "black"}, car{"bar", "yellow"}},
		cars{car{"Baz", "black"}, car{"bar", "yellow"}},
		cars{car{"bar", "yellow"}, car{"Baz", "black"}},
		true,
	},
}

func TestCars_Reverse(t *testing.T) {
	for _, test := range carsSortTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.reversed, test.ss.Reverse())
		})
	}
}

var carsToStringsTests = []struct {
	ss        cars
	transform func(car) string
	expected  Strings
}{
	{
		nil,
		func(s car) string {
			return "foo"
		},
		nil,
	},
	{
		cars{},
		func(s car) string {
			return fmt.Sprintf("%s!", s.Name)
		},
		nil,
	},
	{
		cars{car{"bar", "yellow"}, car{"Baz", "black"}, car{"foo", "red"}},
		func(s car) string {
			return fmt.Sprintf("%s!", s.Color)
		},
		Strings{"yellow!", "black!", "red!"},
	},
}

func TestCars_ToStrings(t *testing.T) {
	for _, test := range carsToStringsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.ToStrings(test.transform))
		})
	}
}

func TestCars_Append(t *testing.T) {
	assert.Equal(t,
		cars{}.Append(),
		cars{},
	)

	assert.Equal(t,
		cars{}.Append(car{"bar", "yellow"}),
		cars{car{"bar", "yellow"}},
	)

	assert.Equal(t,
		cars{}.Append(car{"bar", "yellow"}, car{"Baz", "black"}),
		cars{car{"bar", "yellow"}, car{"Baz", "black"}},
	)

	assert.Equal(t,
		cars{car{"bar", "yellow"}}.Append(car{"Baz", "black"}),
		cars{car{"bar", "yellow"}, car{"Baz", "black"}},
	)

	assert.Equal(t,
		cars{car{"bar", "yellow"}}.Append(car{"Baz", "black"}, car{"foo", "red"}),
		cars{car{"bar", "yellow"}, car{"Baz", "black"}, car{"foo", "red"}},
	)
}

func TestCars_Extend(t *testing.T) {
	assert.Equal(t,
		cars{}.Extend(),
		cars{},
	)

	assert.Equal(t,
		cars{}.Extend(cars{car{"bar", "yellow"}}),
		cars{car{"bar", "yellow"}},
	)

	assert.Equal(t,
		cars{}.Extend(cars{car{"bar", "yellow"}}, cars{car{"Baz", "black"}}),
		cars{car{"bar", "yellow"}, car{"Baz", "black"}},
	)

	assert.Equal(t,
		cars{car{"bar", "yellow"}}.Extend(cars{car{"Baz", "black"}}),
		cars{car{"bar", "yellow"}, car{"Baz", "black"}},
	)

	assert.Equal(t,
		cars{car{"bar", "yellow"}}.Extend(cars{car{"Baz", "black"}, car{"foo", "red"}}),
		cars{car{"bar", "yellow"}, car{"Baz", "black"}, car{"foo", "red"}},
	)
}

func TestCars_All(t *testing.T) {
	assert.True(t,
		(cars)(nil).All(func(value car) bool {
			return false
		}),
	)

	assert.True(t,
		(cars)(nil).All(func(value car) bool {
			return false
		}),
	)

	// None
	assert.False(t,
		cars{car{"bar", "yellow"}, car{"Baz", "black"}}.All(func(value car) bool {
			return value.Color == "green"
		}),
	)

	// Some
	assert.False(t,
		cars{car{"bar", "yellow"}, car{"Baz", "black"}}.All(func(value car) bool {
			return value.Color == "yellow"
		}),
	)

	// All
	assert.True(t,
		cars{car{"bar", "yellow"}, car{"Baz", "black"}}.All(func(value car) bool {
			return len(value.Name) > 0
		}),
	)
}

func TestCars_Any(t *testing.T) {
	assert.False(t,
		(cars)(nil).Any(func(value car) bool {
			return true
		}),
	)

	assert.False(t,
		(cars)(nil).Any(func(value car) bool {
			return true
		}),
	)

	// None
	assert.False(t,
		cars{car{"bar", "yellow"}, car{"Baz", "black"}}.Any(func(value car) bool {
			return value.Color == "green"
		}),
	)

	// Some
	assert.True(t,
		cars{car{"bar", "yellow"}, car{"Baz", "black"}}.Any(func(value car) bool {
			return value.Color == "yellow"
		}),
	)

	// All
	assert.True(t,
		cars{car{"bar", "yellow"}, car{"Baz", "black"}}.Any(func(value car) bool {
			return len(value.Name) > 0
		}),
	)
}

var carsShuffleTests = []struct {
	ss       cars
	expected cars
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
		cars{},
		cars{},
		rand.NewSource(0),
	},
	{
		cars{car{"bar", "yellow"}, car{"Baz", "black"}, car{"foo", "red"}},
		cars{car{"Baz", "black"}, car{"bar", "yellow"}, car{"foo", "red"}},
		rand.NewSource(0),
	},
	{
		cars{car{"bar", "yellow"}},
		cars{car{"bar", "yellow"}},
		rand.NewSource(0),
	},
}

func TestCars_Shuffle(t *testing.T) {
	for _, test := range carsShuffleTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Shuffle(test.source))
		})
	}
}

var carsTopAndBottomTests = []struct {
	ss     cars
	n      int
	top    cars
	bottom cars
}{
	{
		nil,
		1,
		nil,
		nil,
	},
	{
		cars{},
		1,
		nil,
		nil,
	},
	{
		cars{car{"bar", "yellow"}, car{"Baz", "black"}},
		1,
		cars{car{"bar", "yellow"}},
		cars{car{"Baz", "black"}},
	},
	{
		cars{car{"bar", "yellow"}, car{"Baz", "black"}},
		3,
		cars{car{"bar", "yellow"}, car{"Baz", "black"}},
		cars{car{"Baz", "black"}, car{"bar", "yellow"}},
	},
	{
		cars{car{"bar", "yellow"}, car{"Baz", "black"}},
		0,
		nil,
		nil,
	},
	{
		cars{car{"bar", "yellow"}, car{"Baz", "black"}},
		-1,
		nil,
		nil,
	},
}

func TestCars_Top(t *testing.T) {
	for _, test := range carsTopAndBottomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.top, test.ss.Top(test.n))
		})
	}
}

func TestCars_Bottom(t *testing.T) {
	for _, test := range carsTopAndBottomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.bottom, test.ss.Bottom(test.n))
		})
	}
}

func TestCars_Each(t *testing.T) {
	var names []string

	names = []string{}
	cars{}.Each(func(car car) {
		names = append(names, car.Name)
	})
	assert.Equal(t, []string{}, names)

	names = []string{}
	cars{car{"bar", "yellow"}, car{"Baz", "black"}}.Each(func(car car) {
		names = append(names, car.Name)
	})
	assert.Equal(t, []string{"bar", "Baz"}, names)
}

var carsRandomTests = []struct {
	ss       cars
	expected car
	source   rand.Source
}{
	{
		nil,
		car{},
		nil,
	},
	{
		nil,
		car{},
		rand.NewSource(0),
	},
	{
		cars{},
		car{},
		rand.NewSource(0),
	},
	{
		cars{car{"bar", "yellow"}, car{"Baz", "black"}, car{"foo", "red"}},
		car{"bar", "yellow"},
		rand.NewSource(0),
	},
	{
		cars{car{"bar", "yellow"}, car{"Baz", "black"}, car{"foo", "red"}},
		car{"foo", "red"},
		rand.NewSource(1),
	},
	{
		cars{car{"bar", "yellow"}},
		car{"bar", "yellow"},
		rand.NewSource(0),
	},
}

func TestCars_Random(t *testing.T) {
	for _, test := range carsRandomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Random(test.source))
		})
	}
}
