package pie

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
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
