package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var intsContainsTests = []struct {
	ss       Ints
	contains int
	expected bool
}{
	{nil, 1, false},
	{Ints{1, 2, 3}, 1, true},
	{Ints{1, 2, 3}, 2, true},
	{Ints{1, 2, 3}, 3, true},
	{Ints{1, 2, 3}, 4, false},
	{Ints{1, 2, 3}, 5, false},
	{Ints{1, 2, 3}, 6, false},
	{Ints{1, 5, 3}, 5, true},
}

func TestInts_Contains(t *testing.T) {
	for _, test := range intsContainsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, test.ss.Contains(test.contains))
		})
	}
}

var intsOnlyAndWithoutTests = []struct {
	ss                Ints
	condition         func(int) bool
	expectedOnly      Ints
	expectedWithout   Ints
	expectedTransform Ints
}{
	{
		nil,
		func(s int) bool {
			return s == 5
		},
		nil,
		nil,
		nil,
	},
	{
		Ints{1, 2, 3},
		func(s int) bool {
			return s != 2
		},
		Ints{1, 3},
		Ints{2},
		Ints{6, 7, 8},
	},
}

func TestInts_Only(t *testing.T) {
	for _, test := range intsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedOnly, test.ss.Only(test.condition))
		})
	}
}

func TestInts_Without(t *testing.T) {
	for _, test := range intsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedWithout, test.ss.Without(test.condition))
		})
	}
}

func TestInts_Transform(t *testing.T) {
	for _, test := range intsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedTransform, test.ss.Transform(func(i int) int {
				return i + 5
			}))
		})
	}
}

var intsFirstAndLastTests = []struct {
	ss             Ints
	first, firstOr int
	last, lastOr   int
}{
	{
		nil,
		0,
		102,
		0,
		202,
	},
	{
		Ints{100},
		100,
		100,
		100,
		100,
	},
	{
		Ints{1, 2},
		1,
		1,
		2,
		2,
	},
	{
		Ints{1, 2, 3},
		1,
		1,
		3,
		3,
	},
}

func TestInts_FirstOr(t *testing.T) {
	for _, test := range intsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.firstOr, test.ss.FirstOr(102))
		})
	}
}

func TestInts_LastOr(t *testing.T) {
	for _, test := range intsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.lastOr, test.ss.LastOr(202))
		})
	}
}

func TestInts_First(t *testing.T) {
	for _, test := range intsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.first, test.ss.First())
		})
	}
}

func TestInts_Last(t *testing.T) {
	for _, test := range intsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.last, test.ss.Last())
		})
	}
}

var intsStatsTests = []struct {
	ss                 []int
	min, max, sum, len int
	average            float64
}{
	{
		nil,
		0,
		0,
		0,
		0,
		0,
	},
	{
		[]int{},
		0,
		0,
		0,
		0,
		0,
	},
	{
		[]int{1},
		1,
		1,
		1,
		1,
		1,
	},
	{
		[]int{2, 3, 5, 1},
		1,
		5,
		11,
		4,
		2.75,
	},
}

func TestInts_Min(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.min, Ints(test.ss).Min())
		})
	}
}

func TestInts_Max(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.max, Ints(test.ss).Max())
		})
	}
}

func TestInts_Sum(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.sum, Ints(test.ss).Sum())
		})
	}
}

func TestInts_Len(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.len, Ints(test.ss).Len())
		})
	}
}

func TestInts_Average(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.average, Ints(test.ss).Average())
		})
	}
}

var intsJSONTests = []struct {
	ss         Ints
	jsonString string
}{
	{
		nil,
		`[]`, // Instead of null.
	},
	{
		Ints{},
		`[]`,
	},
	{
		Ints{12},
		`[12]`,
	},
	{
		Ints{23, -2, 3424, 12},
		`[23,-2,3424,12]`,
	},
}

func TestInts_JSONString(t *testing.T) {
	for _, test := range intsJSONTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.jsonString, test.ss.JSONString())
		})
	}
}

var intsSortTests = []struct {
	ss        Ints
	sorted    Ints
	reversed  Ints
	areSorted bool
}{
	{
		nil,
		nil,
		nil,
		true,
	},
	{
		Ints{},
		Ints{},
		Ints{},
		true,
	},
	{
		Ints{789},
		Ints{789},
		Ints{789},
		true,
	},
	{
		Ints{12, -13, 789},
		Ints{-13, 12, 789},
		Ints{789, -13, 12},
		false,
	},
	{
		Ints{12, -13, 1.234e6, 789},
		Ints{-13, 12, 789, 1.234e6},
		Ints{789, 1.234e6, -13, 12},
		false,
	},
	{
		Ints{-13, 12},
		Ints{-13, 12},
		Ints{12, -13},
		true,
	},
}

func TestInts_Sort(t *testing.T) {
	for _, test := range intsSortTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.sorted, test.ss.Sort())
		})
	}
}

func TestInts_Reverse(t *testing.T) {
	for _, test := range intsSortTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.reversed, test.ss.Reverse())
		})
	}
}

func TestInts_AreSorted(t *testing.T) {
	for _, test := range intsSortTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.areSorted, test.ss.AreSorted())
		})
	}
}
