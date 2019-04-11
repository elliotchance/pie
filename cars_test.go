package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var carsContainsTests = []struct {
	ss       Cars
	contains Car
	expected bool
}{
	{nil, Car{"a", "green"}, false},
	{nil, Car{}, false},
	{Cars{Car{"a", "green"}, Car{"b", "blue"}, Car{"c", "gray"}}, Car{"a", "green"}, true},
	{Cars{Car{"a", "green"}, Car{"b", "blue"}, Car{"c", "gray"}}, Car{"b", "blue"}, true},
	{Cars{Car{"a", "green"}, Car{"b", "blue"}, Car{"c", "gray"}}, Car{"c", "gray"}, true},
	{Cars{Car{"a", "green"}, Car{"b", "blue"}, Car{"c", "gray"}}, Car{"A", ""}, false},
	{Cars{Car{"a", "green"}, Car{"b", "blue"}, Car{"c", "gray"}}, Car{}, false},
	{Cars{Car{"a", "green"}, Car{"b", "blue"}, Car{"c", "gray"}}, Car{"d", ""}, false},
	{Cars{Car{"a", "green"}, Car{}, Car{"c", "gray"}}, Car{}, true},
}

func TestCars_Contains(t *testing.T) {
	for _, test := range carsContainsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Contains(test.contains))
		})
	}
}

var carsOnlyAndWithoutTests = []struct {
	ss                Cars
	condition         func(Car) bool
	expectedOnly      Cars
	expectedWithout   Cars
	expectedTransform Cars
}{
	{
		nil,
		func(s Car) bool {
			return s.Name == ""
		},
		nil,
		nil,
		nil,
	},
	{
		Cars{Car{"a", "green"}, Car{"b", "blue"}, Car{"c", "gray"}},
		func(s Car) bool {
			return s.Name != "b"
		},
		Cars{Car{"a", "green"}, Car{"c", "gray"}},
		Cars{Car{"b", "blue"}},
		Cars{Car{"A", "green"}, Car{"B", "blue"}, Car{"C", "gray"}},
	},
}

func TestCars_Only(t *testing.T) {
	for _, test := range carsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.expectedOnly, test.ss.Only(test.condition))
		})
	}
}

func TestCars_Without(t *testing.T) {
	for _, test := range carsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.expectedWithout, test.ss.Without(test.condition))
		})
	}
}

func TestCars_Transform(t *testing.T) {
	for _, test := range carsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.expectedTransform, test.ss.Transform(func(car Car) Car {
				car.Name = strings.ToUpper(car.Name)

				return car
			}))
		})
	}
}

var carsFirstAndLastTests = []struct {
	ss             Cars
	first, firstOr Car
	last, lastOr   Car
}{
	{
		nil,
		Car{},
		Car{"default1", "unknown"},
		Car{},
		Car{"default2", "unknown"},
	},
	{
		Cars{Car{"foo", "red"}},
		Car{"foo", "red"},
		Car{"foo", "red"},
		Car{"foo", "red"},
		Car{"foo", "red"},
	},
	{
		Cars{Car{"a", "green"}, Car{"b", "blue"}},
		Car{"a", "green"},
		Car{"a", "green"},
		Car{"b", "blue"},
		Car{"b", "blue"},
	},
	{
		Cars{Car{"a", "green"}, Car{"b", "blue"}, Car{"c", "gray"}},
		Car{"a", "green"},
		Car{"a", "green"},
		Car{"c", "gray"},
		Car{"c", "gray"},
	},
}

func TestCars_FirstOr(t *testing.T) {
	for _, test := range carsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.firstOr, test.ss.FirstOr(Car{"default1", "unknown"}))
		})
	}
}

func TestCars_LastOr(t *testing.T) {
	for _, test := range carsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.lastOr, test.ss.LastOr(Car{"default2", "unknown"}))
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
	ss       Cars
	min, max Car
	len      int
}{
	{
		nil,
		Car{},
		Car{},
		0,
	},
	{
		Cars{},
		Car{},
		Car{},
		0,
	},
	{
		Cars{Car{"foo", "red"}},
		Car{"foo", "red"},
		Car{"foo", "red"},
		1,
	},
	{
		Cars{Car{"bar", "yellow"}, Car{"Baz", "black"}, Car{"qux", "cyan"}, Car{"foo", "red"}},
		Car{"Baz", "black"},
		Car{"qux", "cyan"},
		4,
	},
}

func TestCars_Len(t *testing.T) {
	for _, test := range carsStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableCars(t, &test.ss)()
			assert.Equal(t, test.len, Cars(test.ss).Len())
		})
	}
}

var carsJSONTests = []struct {
	ss         Cars
	jsonString string
}{
	{
		nil,
		`[]`, // Instead of null.
	},
	{
		Cars{},
		`[]`,
	},
	{
		Cars{Car{"foo", "red"}},
		`[{"Name":"foo","Color":"red"}]`,
	},
	{
		Cars{Car{"bar", "yellow"}, Car{"Baz", "black"}, Car{"qux", "cyan"}, Car{"foo", "red"}},
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
	ss        Cars
	sorted    Cars
	reversed  Cars
	areSorted bool
}{
	{
		nil,
		nil,
		nil,
		true,
	},
	{
		Cars{},
		Cars{},
		Cars{},
		true,
	},
	{
		Cars{Car{"foo", "red"}},
		Cars{Car{"foo", "red"}},
		Cars{Car{"foo", "red"}},
		true,
	},
	{
		Cars{Car{"bar", "yellow"}, Car{"Baz", "black"}, Car{"foo", "red"}},
		Cars{Car{"Baz", "black"}, Car{"bar", "yellow"}, Car{"foo", "red"}},
		Cars{Car{"foo", "red"}, Car{"Baz", "black"}, Car{"bar", "yellow"}},
		false,
	},
	{
		Cars{Car{"bar", "yellow"}, Car{"Baz", "black"}, Car{"qux", "cyan"}, Car{"foo", "red"}},
		Cars{Car{"Baz", "black"}, Car{"bar", "yellow"}, Car{"foo", "red"}, Car{"qux", "cyan"}},
		Cars{Car{"foo", "red"}, Car{"qux", "cyan"}, Car{"Baz", "black"}, Car{"bar", "yellow"}},
		false,
	},
	{
		Cars{Car{"Baz", "black"}, Car{"bar", "yellow"}},
		Cars{Car{"Baz", "black"}, Car{"bar", "yellow"}},
		Cars{Car{"bar", "yellow"}, Car{"Baz", "black"}},
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
