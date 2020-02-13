package pie

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

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
	ss                          []int
	min, max, sum, product, len int
	mode                        Ints
	average                     float64
}{
	{
		nil,
		0,
		0,
		0,
		0,
		0,
		nil,
		0,
	},
	{
		[]int{},
		0,
		0,
		0,
		0,
		0,
		Ints{},
		0,
	},
	{
		[]int{1},
		1,
		1,
		1,
		1,
		1,
		Ints{1},
		1,
	},
	{
		[]int{2, 3, 5, 1},
		1,
		5,
		11,
		30,
		4,
		Ints{2, 3, 5, 1},
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

func TestInts_Mode(t *testing.T) {
	cmp := func(a, b Ints) bool {
		m := make(map[int]struct{})
		for _, i := range a {
			m[i] = struct{}{}
		}
		for _, i := range b {
			if _, ok := m[i]; !ok {
				return false
			}
		}
		return true
	}
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			//assert.Equal(t, test.mode, Ints(test.ss).Mode())
			assert.True(t, cmp(test.mode, Ints(test.ss).Mode()))
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

func TestInts_Product(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.product, Ints(test.ss).Product())
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

func TestInts_JSONBytes(t *testing.T) {
	for _, test := range intsJSONTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, []byte(test.jsonString), test.ss.JSONBytes())
		})
	}
}

var intsJSONIndentTests = []struct {
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
		`[
  12
]`,
	},
	{
		Ints{23, -2, 3424, 12},
		`[
  23,
  -2,
  3424,
  12
]`,
	},
}

func TestInts_JSONStringIndent(t *testing.T) {
	for _, test := range intsJSONIndentTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.jsonString, test.ss.JSONStringIndent("", "  "))
		})
	}
}

func TestInts_JSONBytesIndent(t *testing.T) {
	for _, test := range intsJSONIndentTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, []byte(test.jsonString), test.ss.JSONBytesIndent("", "  "))
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

var intsStringsUsingTests = []struct {
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

func TestInts_StringsUsing(t *testing.T) {
	for _, test := range intsStringsUsingTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.StringsUsing(test.transform))
		})
	}
}

func TestInts_Append(t *testing.T) {
	assert.Equal(t,
		len(Ints{}.Append()),
		0,
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

var intsReduceTests = []struct {
	ss       Ints
	expected int
	reducer  func(a, b int) int
}{
	{
		Ints{1, 2, 3},
		6,
		func(a, b int) int { return a + b },
	},
	{
		Ints{1, 2, 3},
		-4,
		func(a, b int) int { return a - b },
	},
	{
		Ints{},
		0,
		func(a, b int) int { return a - b },
	},
	{
		Ints{1},
		1,
		func(a, b int) int { return a - b },
	},
}

func TestInts_Reduce(t *testing.T) {
	for _, test := range intsReduceTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, test.ss.Reduce(test.reducer))
		})
	}
}

func TestInts_Abs(t *testing.T) {
	assert.Equal(t, Ints{1, 5, 7}, Ints{-1, 5, -7}.Abs())
	assert.Equal(t, Ints{689845, 688969, 220373, 89437, 308836}, Ints{-689845, -688969, -220373, -89437, 308836}.Abs())
	assert.Equal(t, Ints{1, 2}, Ints{1, 2}.Abs())
}

func TestInts_AbsLarge(t *testing.T) {
	// int64: prevent compile error (constant overflow) on 32-bit architectures.
	a64 := []int64{3, -4, 1234567890123456789, -987654321098765432}
	var a Ints
	for _, v := range a64 {
		a = a.Append(int(v))
	}
	b := a.Abs()
	for i := range a {
		v, absv := a[i], b[i]
		if absv != v && absv != -v {
			t.Errorf("Incorrect result for Abs(%d): %d", v, absv)
		}
	}
}

var intsSendTests = []struct {
	ss            Ints
	recieveDelay  time.Duration
	canceledDelay time.Duration
	expected      Ints
}{
	{
		nil,
		0,
		0,
		nil,
	},
	{
		Ints{1, 3},
		0,
		0,
		Ints{1, 3},
	},
	{
		Ints{1, 3},
		time.Millisecond * 30,
		time.Millisecond * 10,
		Ints{1},
	},
	{
		Ints{1, 3},
		time.Millisecond * 3,
		time.Millisecond * 10,
		Ints{1, 3},
	},
}

func TestInts_Send(t *testing.T) {
	for _, test := range intsSendTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			ch := make(chan int)
			actual := getIntsFromChan(ch, test.recieveDelay)
			ctx := createContextByDelay(test.canceledDelay)

			actualSended := test.ss.Send(ctx, ch)
			close(ch)

			assert.Equal(t, test.expected, actualSended)
			assert.Equal(t, test.expected, actual())
		})
	}
}

var intsIntersectTests = []struct {
	ss       Ints
	params   []Ints
	expected Ints
}{
	{
		nil,
		nil,
		nil,
	},
	{
		Ints{1, 3},
		nil,
		nil,
	},
	{
		nil,
		[]Ints{{1, 3, 5}, {5, 1}},
		nil,
	},
	{
		Ints{1, 3},
		[]Ints{{1}, {3}},
		nil,
	},
	{
		Ints{1, 3},
		[]Ints{{1}},
		Ints{1},
	},
	{
		Ints{1, 3},
		[]Ints{{1, 3, 5}, {5, 1}},
		Ints{1},
	},
}

func TestInts_Intersect(t *testing.T) {
	for _, test := range intsIntersectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Intersect(test.params...))
		})
	}
}

var intsDiffTests = map[string]struct {
	ss1     Ints
	ss2     Ints
	added   Ints
	removed Ints
}{
	"BothEmpty": {
		nil,
		nil,
		nil,
		nil,
	},
	"OnlyRemovedUnique": {
		Ints{1, 3},
		nil,
		nil,
		Ints{1, 3},
	},
	"OnlyRemovedDuplicates": {
		Ints{1, 3, 1},
		nil,
		nil,
		Ints{1, 3, 1},
	},
	"OnlyAddedUnique": {
		nil,
		Ints{2, 3},
		Ints{2, 3},
		nil,
	},
	"OnlyAddedDuplicates": {
		nil,
		Ints{2, 3, 3, 2},
		Ints{2, 3, 3, 2},
		nil,
	},
	"AddedAndRemovedUnique": {
		Ints{1, 2, 3, 4},
		Ints{3, 4, 5, 6},
		Ints{5, 6},
		Ints{1, 2},
	},
	"AddedAndRemovedDuplicates": {
		Ints{1, 2, 3, 3, 4},
		Ints{3, 4, 5, 4, 6},
		Ints{5, 4, 6},
		Ints{1, 2, 3},
	},
}

func TestInts_Diff(t *testing.T) {
	for testName, test := range intsDiffTests {
		t.Run(testName, func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss1)()
			defer assertImmutableInts(t, &test.ss2)()

			added, removed := test.ss1.Diff(test.ss2)
			assert.Equal(t, test.added, added)
			assert.Equal(t, test.removed, removed)
		})
	}
}

var intsSequenceAndSequenceUsingTests = []struct {
	ss       Ints
	params   []int
	expected Ints
}{
	// n
	{
		nil,
		nil,
		nil,
	},
	{
		nil,
		[]int{-1},
		nil,
	},
	{
		nil,
		[]int{0},
		nil,
	},
	{
		nil,
		[]int{3},
		Ints{0, 1, 2},
	},
	{
		Ints{},
		[]int{3},
		Ints{0, 1, 2},
	},
	// range
	{
		nil,
		[]int{2, 2},
		nil,
	},
	{
		Ints{},
		[]int{3, 2},
		nil,
	},
	{
		nil,
		[]int{0, 3},
		Ints{0, 1, 2},
	},
	{
		Ints{},
		[]int{3, 6},
		Ints{3, 4, 5},
	},
	{
		Ints{},
		[]int{-5, 0},
		Ints{-5, -4, -3, -2, -1},
	},
	{
		Ints{},
		[]int{-5, -10},
		nil,
	},
	// range with step
	{
		nil,
		[]int{3, 3, 1},
		nil,
	},
	{
		Ints{},
		[]int{3, 6, 2},
		Ints{3, 5},
	},
	{
		Ints{},
		[]int{3, 7, 2},
		Ints{3, 5},
	},
	{
		Ints{},
		[]int{-10, -6, 1},
		Ints{-10, -9, -8, -7},
	},
	{
		Ints{},
		[]int{-6, -10, -1},
		Ints{-6, -7, -8, -9},
	},
	{
		Ints{},
		[]int{-6, -10, 1},
		nil,
	},
}

func TestInts_Sequence(t *testing.T) {
	for _, test := range intsSequenceAndSequenceUsingTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Sequence(test.params...))
		})
	}
}

func TestInts_SequenceUsing(t *testing.T) {
	for _, test := range intsSequenceAndSequenceUsingTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.SequenceUsing(func(i int) int { return i }, test.params...))
		})
	}
}

func TestInts_Strings(t *testing.T) {
	assert.Equal(t, Strings(nil), Ints{}.Strings())

	assert.Equal(t,
		Strings{"92", "823", "453"},
		Ints{92, 823, 453}.Strings())
}

func TestInts_Ints(t *testing.T) {
	assert.Equal(t, Ints(nil), Ints{}.Ints())

	assert.Equal(t,
		Ints{92, 823, 453},
		Ints{92, 823, 453}.Ints())
}

func TestInts_Float64s(t *testing.T) {
	assert.Equal(t, Float64s(nil), Ints{}.Float64s())

	assert.Equal(t,
		Float64s{92, 823, 453},
		Ints{92, 823, 453}.Float64s())
}

var intsDropTopTests = []struct {
	ss      Ints
	n       int
	dropTop Ints
}{
	{
		nil,
		1,
		nil,
	},
	{
		Ints{},
		1,
		nil,
	},
	{
		Ints{1, 2},
		-1,
		nil,
	},
	{
		Ints{1, 2},
		0,
		Ints{1, 2},
	},

	{
		Ints{1, 2},
		1,
		Ints{2},
	},
	{
		Ints{1, 2},
		2,
		nil,
	},
	{
		Ints{1, 2},
		3,
		nil,
	},
}

func TestInts_DropTop(t *testing.T) {
	for _, test := range intsDropTopTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.dropTop, test.ss.DropTop(test.n))
		})
	}
}

var intsSubSliceTests = []struct {
	ss       Ints
	start    int
	end      int
	subSlice Ints
}{
	{
		nil,
		1,
		1,
		nil,
	},
	{
		nil,
		1,
		2,
		Ints{0},
	},
	{
		Ints{},
		1,
		1,
		nil,
	},
	{
		Ints{},
		1,
		2,
		Ints{0},
	},
	{
		Ints{1, 2},
		-1,
		-1,
		nil,
	},
	{
		Ints{1, 2},
		-1,
		1,
		nil,
	},
	{
		Ints{1, 2},
		1,
		-1,
		nil,
	},
	{
		Ints{1, 2},
		2,
		0,
		nil,
	},

	{
		Ints{1, 2},
		1,
		1,
		nil,
	},
	{
		Ints{1, 2},
		1,
		2,
		Ints{2},
	},
	{
		Ints{1, 2},
		1,
		3,
		Ints{2, 0},
	},
	{
		Ints{1, 2},
		2,
		2,
		nil,
	},
	{
		Ints{1, 2},
		2,
		3,
		Ints{0},
	},
	{
		Ints{1, 2, 0},
		2,
		3,
		Ints{0},
	},
}

func TestInts_SubSlice(t *testing.T) {
	for _, test := range intsSubSliceTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.subSlice, test.ss.SubSlice(test.start, test.end))
		})
	}
}

var intsFindFirstUsingTests = []struct {
	ss         Ints
	expression func(value int) bool
	expected   int
}{
	{
		nil,
		func(value int) bool { return value == 10 },
		-1,
	},
	{
		Ints{},
		func(value int) bool { return value == 150 },
		-1,
	},
	{
		Ints{10, 15},
		func(value int) bool { return value == 150 },
		-1,
	},
	{
		Ints{100},
		func(value int) bool { return value == 100 },
		0,
	},
	{
		Ints{1, 2},
		func(value int) bool { return value == 2 },
		1,
	},
	{
		Ints{1, 2, 3},
		func(value int) bool { return value == 3 },
		2,
	},
}

func TestInts_FindFirstUsing(t *testing.T) {
	for _, test := range intsFindFirstUsingTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, test.ss.FindFirstUsing(test.expression))
		})
	}
}

var intsEqualsTests = []struct {
	ss       Ints
	rhs      Ints
	expected bool
}{
	{nil, nil, true},
	{Ints{}, Ints{}, true},
	{nil, Ints{}, true},
	{Ints{1, 2}, Ints{1, 2}, true},
	{Ints{1, 2}, Ints{1, 5}, false},
	{Ints{1, 2}, Ints{1}, false},
	{Ints{1}, Ints{2}, false},
	{Ints{1}, nil, false},
}

func TestInts_Equals(t *testing.T) {
	for _, test := range intsEqualsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Equals(test.rhs))
		})
	}
}

var intsShiftAndUnshiftTests = []struct {
	ss      Ints
	shifted int
	shift   Ints
	params  Ints
	unshift Ints
}{
	{
		nil,
		0,
		nil,
		nil,
		Ints{},
	},
	{
		nil,
		0,
		nil,
		Ints{},
		Ints{},
	},
	{
		nil,
		0,
		nil,
		Ints{1, 2},
		Ints{1, 2},
	},
	{
		Ints{},
		0,
		nil,
		nil,
		Ints{},
	},
	{
		Ints{},
		0,
		nil,
		Ints{},
		Ints{},
	},
	{
		Ints{},
		0,
		nil,
		Ints{1, 2},
		Ints{1, 2},
	},
	{
		Ints{1},
		1,
		nil,
		Ints{2},
		Ints{2, 1},
	},
	{
		Ints{1, 2},
		1,
		Ints{2},
		Ints{3},
		Ints{3, 1, 2},
	},
	{
		Ints{1, 2},
		1,
		Ints{2},
		Ints{3, 4},
		Ints{3, 4, 1, 2},
	},
}

func TestInts_ShiftAndUnshift(t *testing.T) {
	for _, test := range intsShiftAndUnshiftTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableInts(t, &test.ss)()
			shifted, shift := test.ss.Shift()
			assert.Equal(t, test.shifted, shifted)
			assert.Equal(t, test.shift, shift)
			assert.Equal(t, test.unshift, test.ss.Unshift(test.params...))
		})
	}
}
func TestInts_Join(t *testing.T) {
	assert.Equal(t, "", Ints(nil).Join("a"))
	assert.Equal(t, "", Ints{}.Join("a"))
	assert.Equal(t, "1-2-3", Ints{1, 2, 3}.Join("-"))
	assert.Equal(t, "1--2-3", Ints{1, -2, 3}.Join("-"))
}

func TestInts_Pop(t *testing.T) {

	numbers := Ints{42, 999}

	assert.Equal(t, 42, *numbers.Pop())
	assert.Equal(t, Ints{999}, numbers)

	assert.Equal(t, 999, *numbers.Pop())
	assert.Equal(t, Ints{}, numbers)
}

func TestInts_Group(t *testing.T) {
	assert.Equal(t, map[int]int{}, Ints(nil).Group())

	assert.Equal(t, map[int]int{
		1: 1,
	}, Ints{1}.Group())

	assert.Equal(t, map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}, Ints{1, 2, 2, 3, 3, 3}.Group())
}

func TestInts_IntersectUsing(t *testing.T) {
	equals := func(a, b int) (bool, int) {
		return a == b, a
	}
	assert.Equal(t, Ints(nil), Ints{0, 1}.IntersectUsing(equals))
	assert.Equal(t, Ints{}, Ints{0, 1}.IntersectUsing(equals, Ints{2}))
	assert.Equal(t, Ints{0}, Ints{0, 1}.IntersectUsing(equals, Ints{0}))

	// We have to sort the slice because it is always returned in random order.
	assert.Equal(t, Ints{1}, Ints{0, 1}.IntersectUsing(equals, Ints{1, 2}, Ints{0, 1, 2}).Sort())
	assert.Equal(t, Ints{}, Ints{0, 1}.IntersectUsing(equals, Ints{1}, Ints{0}).Sort())
	assert.Equal(t, Ints{1}, Ints{0, 1}.IntersectUsing(equals, Ints{1}, Ints{1, 2, 3}).Sort())
	assert.Equal(t, Ints{}, Ints{0, 1}.IntersectUsing(equals, Ints{0}, Ints{1, 2, 3}).Sort())
	assert.Equal(t, Ints{}, Ints{0, 1}.IntersectUsing(equals, Ints{1, 2, 3}, Ints{0}).Sort())
}

func TestInts_Insert(t *testing.T) {

	assert.Equal(t, Ints{}, Ints(nil).Insert(0))
	assert.Equal(t, Ints{2, 1}, Ints{1}.Insert(0, 2))
	assert.Equal(t, Ints{1, 2}, Ints{1}.Insert(1, 2))
	assert.Equal(t, Ints{1, 2, 3}, Ints{1, 3}.Insert(1, 2))
}
