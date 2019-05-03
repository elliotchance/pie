package pie

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/elliotchance/testify-stats/assert"
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

var intsFilterTests = []struct {
	ss                Ints
	condition         func(int) bool
	expectedFilter    Ints
	expectedFilterNot Ints
	expectedMap       Ints
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

func TestInts_Filter(t *testing.T) {
	for _, test := range intsFilterTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedFilter, test.ss.Filter(test.condition))
		})
	}
}

func TestInts_FilterNot(t *testing.T) {
	for _, test := range intsFilterTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedFilterNot, test.ss.FilterNot(test.condition))
		})
	}
}

func TestInts_Map(t *testing.T) {
	for _, test := range intsFilterTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedMap, test.ss.Map(func(i int) int {
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
		Ints{12, -13, 14e6, 789},
		Ints{-13, 12, 789, 14e6},
		Ints{789, 14e6, -13, 12},
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

var intsUniqueTests = []struct {
	ss        Ints
	unique    Ints
	areUnique bool
}{
	{
		nil,
		nil,
		true,
	},
	{
		Ints{},
		Ints{},
		true,
	},
	{
		Ints{789},
		Ints{789},
		true,
	},
	{
		Ints{12, -13, 12},
		Ints{-13, 12},
		false,
	},
	{
		Ints{12, -13, 14e6, 789},
		Ints{-13, 12, 789, 14e6},
		true,
	},
}

func TestInts_Unique(t *testing.T) {
	for _, test := range intsUniqueTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()

			// We have to sort the unique slice because it is always returned in
			// random order.
			assert.Equal(t, test.unique, test.ss.Unique().Sort())
		})
	}
}

func TestInts_AreUnique(t *testing.T) {
	for _, test := range intsUniqueTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.areUnique, test.ss.AreUnique())
		})
	}
}

var intsToStringsTests = []struct {
	ss        Ints
	transform func(int) string
	expected  Strings
}{
	{
		nil,
		func(s int) string {
			return "foo"
		},
		nil,
	},
	{
		Ints{},
		func(s int) string {
			return fmt.Sprintf("%d!", s)
		},
		nil,
	},
	{
		Ints{6, 7, 8},
		func(s int) string {
			return fmt.Sprintf("%d!", s)
		},
		Strings{"6!", "7!", "8!"},
	},
}

func TestInts_ToStrings(t *testing.T) {
	for _, test := range intsToStringsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.ToStrings(test.transform))
		})
	}
}

func TestInts_Append(t *testing.T) {
	assert.Equal(t,
		Ints{}.Append(),
		Ints{},
	)

	assert.Equal(t,
		Ints{}.Append(1),
		Ints{1},
	)

	assert.Equal(t,
		Ints{}.Append(1, 2),
		Ints{1, 2},
	)

	assert.Equal(t,
		Ints{1}.Append(2),
		Ints{1, 2},
	)

	assert.Equal(t,
		Ints{1}.Append(2, 5),
		Ints{1, 2, 5},
	)
}

func TestInts_Extend(t *testing.T) {
	assert.Equal(t,
		Ints{}.Extend(),
		Ints{},
	)

	assert.Equal(t,
		Ints{}.Extend([]int{1}),
		Ints{1},
	)

	assert.Equal(t,
		Ints{}.Extend([]int{1}, []int{2}),
		Ints{1, 2},
	)

	assert.Equal(t,
		Ints{1}.Extend([]int{2}),
		Ints{1, 2},
	)

	assert.Equal(t,
		Ints{1}.Extend([]int{2, 5}),
		Ints{1, 2, 5},
	)
}

func TestInts_All(t *testing.T) {
	assert.True(t,
		Ints{}.All(func(value int) bool {
			return false
		}),
	)

	assert.True(t,
		Ints{}.All(func(value int) bool {
			return false
		}),
	)

	// None
	assert.False(t,
		Ints{12, 4}.All(func(value int) bool {
			return value == 1
		}),
	)

	// Some
	assert.False(t,
		Ints{12, 4}.All(func(value int) bool {
			return value == 12
		}),
	)

	// All
	assert.True(t,
		Ints{12, 4}.All(func(value int) bool {
			return value > 0
		}),
	)
}

func TestInts_Any(t *testing.T) {
	assert.False(t,
		Ints{}.Any(func(value int) bool {
			return true
		}),
	)

	assert.False(t,
		Ints{}.Any(func(value int) bool {
			return true
		}),
	)

	// None
	assert.False(t,
		Ints{12, 4}.Any(func(value int) bool {
			return value == 1
		}),
	)

	// Some
	assert.True(t,
		Ints{12, 4}.Any(func(value int) bool {
			return value == 12
		}),
	)

	// All
	assert.True(t,
		Ints{12, 4}.Any(func(value int) bool {
			return value > 0
		}),
	)
}

var intsShuffleTests = []struct {
	ss       Ints
	expected Ints
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
		Ints{},
		Ints{},
		rand.NewSource(0),
	},
	{
		Ints{1, 2, 4},
		Ints{2, 1, 4},
		rand.NewSource(0),
	},
	{
		Ints{12},
		Ints{12},
		rand.NewSource(0),
	},
}

func TestInts_Shuffle(t *testing.T) {
	for _, test := range intsShuffleTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Shuffle(test.source))
		})
	}
}

var intsTopAndBottomTests = []struct {
	ss     Ints
	n      int
	top    Ints
	bottom Ints
}{
	{
		nil,
		1,
		nil,
		nil,
	},
	{
		Ints{},
		1,
		nil,
		nil,
	},
	{
		Ints{1, 2},
		1,
		Ints{1},
		Ints{2},
	},
	{
		Ints{1, 2},
		3,
		Ints{1, 2},
		Ints{2, 1},
	},
	{
		Ints{1, 2},
		0,
		nil,
		nil,
	},
	{
		Ints{1, 2},
		-1,
		nil,
		nil,
	},
}

func TestInts_Top(t *testing.T) {
	for _, test := range intsTopAndBottomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.top, test.ss.Top(test.n))
		})
	}
}

func TestInts_Bottom(t *testing.T) {
	for _, test := range intsTopAndBottomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.bottom, test.ss.Bottom(test.n))
		})
	}
}

func TestInts_Median(t *testing.T) {
	assert.Equal(t, 0, Ints{}.Median())
	assert.Equal(t, 12, Ints{12}.Median())
	assert.Equal(t, 8, Ints{12, 4}.Median())
	assert.Equal(t, 4, Ints{2, 12, 4}.Median())
}

func TestInts_Each(t *testing.T) {
	var values []int

	values = []int{}
	Ints{}.Each(func(value int) {
		values = append(values, value)
	})
	assert.Equal(t, []int{}, values)

	values = []int{}
	Ints{435, 8923}.Each(func(value int) {
		values = append(values, value)
	})
	assert.Equal(t, []int{435, 8923}, values)
}

var intsRandomTests = []struct {
	ss       Ints
	expected int
	source   rand.Source
}{
	{
		nil,
		0,
		nil,
	},
	{
		nil,
		0,
		rand.NewSource(0),
	},
	{
		Ints{},
		0,
		rand.NewSource(0),
	},
	{
		Ints{1, 2, 4},
		4,
		rand.NewSource(1),
	},
	{
		Ints{1, 2, 4},
		1,
		rand.NewSource(0),
	},
	{
		Ints{12},
		12,
		rand.NewSource(0),
	},
}

func TestInts_Random(t *testing.T) {
	for _, test := range intsRandomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Random(test.source))
		})
	}
}

func TestInts_Abs(t *testing.T) {
	assert.Equal(t, Ints{1, 5, 7}, Ints{-1, 5, -7}.Abs())
	assert.Equal(t, Ints{689845, 688969, 220373, 89437, 308836}, Ints{-689845, -688969, -220373, -89437, 308836}.Abs())
	assert.Equal(t, Ints{1, 2}, Ints{1, 2}.Abs())
}
