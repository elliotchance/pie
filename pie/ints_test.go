package pie

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

var intsSelectTests = []struct {
	ss                Ints
	condition         func(int) bool
	expectedSelect    Ints
	expectedUnselect  Ints
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

func TestInts_Select(t *testing.T) {
	for _, test := range intsSelectTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedSelect, test.ss.Select(test.condition))
		})
	}
}

func TestInts_Unselect(t *testing.T) {
	for _, test := range intsSelectTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedUnselect, test.ss.Unselect(test.condition))
		})
	}
}

func TestInts_Transform(t *testing.T) {
	for _, test := range intsSelectTests {
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

var intsTopTests = []struct {
	ss  Ints
	top Ints
	n   int
}{
	{
		nil,
		nil,
		1,
	},
	{
		Ints{},
		nil,
		1,
	},
	{
		Ints{1, 2},
		Ints{1},
		1,
	},
	{
		Ints{1, 2},
		Ints{1, 2},
		3,
	},
	{
		Ints{1, 2},
		nil,
		0,
	},
	{
		Ints{1, 2},
		nil,
		-1,
	},
}

func TestInts_Top(t *testing.T) {
	for _, test := range intsTopTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.top, test.ss.Top(test.n))
		})
	}
}

var intsBottomTests = []struct {
	ss     Ints
	bottom Ints
	n      int
}{
	{
		nil,
		nil,
		1,
	},
	{
		Ints{},
		nil,
		1,
	},
	{
		Ints{1, 2},
		Ints{2},
		1,
	},
	{
		Ints{1, 2},
		Ints{2, 1},
		3,
	},
	{
		Ints{1, 2},
		nil,
		0,
	},
	{
		Ints{1, 2},
		nil,
		-1,
	},
}

func TestInts_Bottom(t *testing.T) {
	for _, test := range intsBottomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.bottom, test.ss.Bottom(test.n))
		})
	}
}
