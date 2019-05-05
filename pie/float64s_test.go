package pie

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/elliotchance/testify-stats/assert"
)

var float64sContainsTests = []struct {
	ss       Float64s
	contains float64
	expected bool
}{
	{nil, 1, false},
	{Float64s{1, 2, 3}, 1, true},
	{Float64s{1, 2, 3}, 2, true},
	{Float64s{1, 2, 3}, 3, true},
	{Float64s{1, 2, 3}, 4, false},
	{Float64s{1, 2, 3}, 5, false},
	{Float64s{1, 2, 3}, 6, false},
	{Float64s{1, 5, 3}, 5, true},
}

func TestFloat64s_Contains(t *testing.T) {
	for _, test := range float64sContainsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Contains(test.contains))
		})
	}
}

var float64sSelectTests = []struct {
	ss                Float64s
	condition         func(float64) bool
	expectedFilter    Float64s
	expectedFilterNot Float64s
	expectedMap       Float64s
}{
	{
		nil,
		func(s float64) bool {
			return s == 5
		},
		nil,
		nil,
		nil,
	},
	{
		Float64s{1, 2, 3},
		func(s float64) bool {
			return s != 2
		},
		Float64s{1, 3},
		Float64s{2},
		Float64s{6.2, 7.2, 8.2},
	},
}

func TestFloat64s_Filter(t *testing.T) {
	for _, test := range float64sSelectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.expectedFilter, test.ss.Filter(test.condition))
		})
	}
}

func TestFloat64s_FilterNot(t *testing.T) {
	for _, test := range float64sSelectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.expectedFilterNot, test.ss.FilterNot(test.condition))
		})
	}
}

func TestFloat64s_Map(t *testing.T) {
	for _, test := range float64sSelectTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.expectedMap, test.ss.Map(func(a float64) float64 {
				return a + 5.2
			}))
		})
	}
}

var float64sFirstAndLastTests = []struct {
	ss             Float64s
	first, firstOr float64
	last, lastOr   float64
}{
	{
		nil,
		0,
		102,
		0,
		202,
	},
	{
		Float64s{100},
		100,
		100,
		100,
		100,
	},
	{
		Float64s{1, 2},
		1,
		1,
		2,
		2,
	},
	{
		Float64s{1, 2, 3},
		1,
		1,
		3,
		3,
	},
}

func TestFloat64s_FirstOr(t *testing.T) {
	for _, test := range float64sFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.firstOr, test.ss.FirstOr(102))
		})
	}
}

func TestFloat64s_LastOr(t *testing.T) {
	for _, test := range float64sFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.lastOr, test.ss.LastOr(202))
		})
	}
}

func TestFloat64s_First(t *testing.T) {
	for _, test := range float64sFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.first, test.ss.First())
		})
	}
}

func TestFloat64s_Last(t *testing.T) {
	for _, test := range float64sFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.last, test.ss.Last())
		})
	}
}

var float64sStatsTests = []struct {
	ss            Float64s
	min, max, sum float64
	len           int
	average       float64
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
		[]float64{},
		0,
		0,
		0,
		0,
		0,
	},
	{
		[]float64{1.5},
		1.5,
		1.5,
		1.5,
		1,
		1.5,
	},
	{
		[]float64{2.2, 3.1, 5.1, 1.9},
		1.9,
		5.1,
		12.3,
		4,
		3.075,
	},
}

func TestFloat64s_Min(t *testing.T) {
	for _, test := range float64sStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.min, Float64s(test.ss).Min())
		})
	}
}

func TestFloat64s_Max(t *testing.T) {
	for _, test := range float64sStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.max, Float64s(test.ss).Max())
		})
	}
}

func TestFloat64s_Sum(t *testing.T) {
	for _, test := range float64sStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.sum, Float64s(test.ss).Sum())
		})
	}
}

func TestFloat64s_Len(t *testing.T) {
	for _, test := range float64sStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.len, Float64s(test.ss).Len())
		})
	}
}

func TestFloat64s_Average(t *testing.T) {
	for _, test := range float64sStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.average, Float64s(test.ss).Average())
		})
	}
}

var float64sJSONTests = []struct {
	ss         Float64s
	jsonString string
}{
	{
		nil,
		`[]`, // Instead of null.
	},
	{
		Float64s{},
		`[]`,
	},
	{
		Float64s{12.3},
		`[12.3]`,
	},
	{
		Float64s{23, -2.5, 3424, 12.3},
		`[23,-2.5,3424,12.3]`,
	},
}

func TestFloat64s_JSONString(t *testing.T) {
	for _, test := range float64sJSONTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.jsonString, test.ss.JSONString())
		})
	}
}

var float64sSortTests = []struct {
	ss        Float64s
	sorted    Float64s
	reversed  Float64s
	areSorted bool
}{
	{
		nil,
		nil,
		nil,
		true,
	},
	{
		Float64s{},
		Float64s{},
		Float64s{},
		true,
	},
	{
		Float64s{789},
		Float64s{789},
		Float64s{789},
		true,
	},
	{
		Float64s{12.789, -13.2, 789},
		Float64s{-13.2, 12.789, 789},
		Float64s{789, -13.2, 12.789},
		false,
	},
	{
		Float64s{12.789, -13.2, 1.234e6, 789},
		Float64s{-13.2, 12.789, 789, 1.234e6},
		Float64s{789, 1.234e6, -13.2, 12.789},
		false,
	},
	{
		Float64s{-13.2, 12.789},
		Float64s{-13.2, 12.789},
		Float64s{12.789, -13.2},
		true,
	},
}

func TestFloat64s_Sort(t *testing.T) {
	for _, test := range float64sSortTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.sorted, test.ss.Sort())
		})
	}
}

func TestFloat64s_Reverse(t *testing.T) {
	for _, test := range float64sSortTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.reversed, test.ss.Reverse())
		})
	}
}

func TestFloat64s_AreSorted(t *testing.T) {
	for _, test := range float64sSortTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.areSorted, test.ss.AreSorted())
		})
	}
}

var float64sUniqueTests = []struct {
	ss        Float64s
	unique    Float64s
	areUnique bool
}{
	{
		nil,
		nil,
		true,
	},
	{
		Float64s{},
		Float64s{},
		true,
	},
	{
		Float64s{789},
		Float64s{789},
		true,
	},
	{
		Float64s{12.789, -13.2, 12.789},
		Float64s{-13.2, 12.789},
		false,
	},
	{
		Float64s{12.789, -13.2, 1.234e6, 789},
		Float64s{-13.2, 12.789, 789, 1.234e6},
		true,
	},
}

func TestFloat64s_Unique(t *testing.T) {
	for _, test := range float64sUniqueTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()

			// We have to sort the unique slice because it is always returned in
			// random order.
			assert.Equal(t, test.unique, test.ss.Unique().Sort())
		})
	}
}

func TestFloat64s_AreUnique(t *testing.T) {
	for _, test := range float64sUniqueTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.areUnique, test.ss.AreUnique())
		})
	}
}

var float64sToStringsTests = []struct {
	ss        Float64s
	transform func(float64) string
	expected  Strings
}{
	{
		nil,
		func(s float64) string {
			return "foo"
		},
		nil,
	},
	{
		Float64s{},
		func(s float64) string {
			return fmt.Sprintf("%f!", s)
		},
		nil,
	},
	{
		Float64s{6.2, 7.2, 8.2},
		func(s float64) string {
			return fmt.Sprintf("%.2f!", s)
		},
		Strings{"6.20!", "7.20!", "8.20!"},
	},
}

func TestFloat64s_ToStrings(t *testing.T) {
	for _, test := range float64sToStringsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.ToStrings(test.transform))
		})
	}
}

func TestFloat64s_Append(t *testing.T) {
	assert.Equal(t,
		Float64s{}.Append(),
		Float64s{},
	)

	assert.Equal(t,
		Float64s{}.Append(1.23),
		Float64s{1.23},
	)

	assert.Equal(t,
		Float64s{}.Append(1.23, 2.34),
		Float64s{1.23, 2.34},
	)

	assert.Equal(t,
		Float64s{1.23}.Append(2.34),
		Float64s{1.23, 2.34},
	)

	assert.Equal(t,
		Float64s{1.23}.Append(2.34, 5.67),
		Float64s{1.23, 2.34, 5.67},
	)
}

func TestFloat64s_Extend(t *testing.T) {
	assert.Equal(t,
		Float64s{}.Extend(),
		Float64s{},
	)

	assert.Equal(t,
		Float64s{}.Extend([]float64{1.23}),
		Float64s{1.23},
	)

	assert.Equal(t,
		Float64s{}.Extend([]float64{1.23, 2.34}),
		Float64s{1.23, 2.34},
	)

	assert.Equal(t,
		Float64s{1.23}.Extend([]float64{2.34}),
		Float64s{1.23, 2.34},
	)

	assert.Equal(t,
		Float64s{1.23}.Extend([]float64{2.34, 5.67}),
		Float64s{1.23, 2.34, 5.67},
	)
}

func TestFloat64s_All(t *testing.T) {
	assert.True(t,
		Float64s{}.All(func(value float64) bool {
			return false
		}),
	)

	assert.True(t,
		Float64s{}.All(func(value float64) bool {
			return false
		}),
	)

	// None
	assert.False(t,
		Float64s{12.3, 4.56}.All(func(value float64) bool {
			return value == 1
		}),
	)

	// Some
	assert.False(t,
		Float64s{12.3, 4.56}.All(func(value float64) bool {
			return value == 12.3
		}),
	)

	// All
	assert.True(t,
		Float64s{12.3, 4.56}.All(func(value float64) bool {
			return value > 0
		}),
	)
}

func TestFloat64s_Any(t *testing.T) {
	assert.False(t,
		Float64s{}.Any(func(value float64) bool {
			return true
		}),
	)

	assert.False(t,
		Float64s{}.Any(func(value float64) bool {
			return true
		}),
	)

	// None
	assert.False(t,
		Float64s{12.3, 4.56}.Any(func(value float64) bool {
			return value == 1
		}),
	)

	// Some
	assert.True(t,
		Float64s{12.3, 4.56}.Any(func(value float64) bool {
			return value == 12.3
		}),
	)

	// All
	assert.True(t,
		Float64s{12.3, 4.56}.Any(func(value float64) bool {
			return value > 0
		}),
	)
}

var float64sShuffleTests = []struct {
	ss       Float64s
	expected Float64s
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
		Float64s{},
		Float64s{},
		rand.NewSource(0),
	},
	{
		Float64s{12.3, 2.34, 4.56},
		Float64s{2.34, 12.3, 4.56},
		rand.NewSource(0),
	},
	{
		Float64s{12.3},
		Float64s{12.3},
		rand.NewSource(0),
	},
}

func TestFloat64s_Shuffle(t *testing.T) {
	for _, test := range float64sShuffleTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Shuffle(test.source))
		})
	}
}

var float64sTopAndBottomTests = []struct {
	ss     Float64s
	n      int
	top    Float64s
	bottom Float64s
}{
	{
		nil,
		1,
		nil,
		nil,
	},
	{
		Float64s{},
		1,
		nil,
		nil,
	},
	{
		Float64s{1.23, 2.34},
		1,
		Float64s{1.23},
		Float64s{2.34},
	},
	{
		Float64s{1.23, 2.34},
		3,
		Float64s{1.23, 2.34},
		Float64s{2.34, 1.23},
	},
	{
		Float64s{1.23, 2.34},
		0,
		nil,
		nil,
	},
	{
		Float64s{1.23, 2.34},
		-1,
		nil,
		nil,
	},
}

func TestFloat64s_Top(t *testing.T) {
	for _, test := range float64sTopAndBottomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.top, test.ss.Top(test.n))
		})
	}
}

func TestFloat64s_Bottom(t *testing.T) {
	for _, test := range float64sTopAndBottomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.bottom, test.ss.Bottom(test.n))
		})
	}
}

func TestFloat64s_Median(t *testing.T) {
	assert.Equal(t, 0.0, Float64s{}.Median())
	assert.Equal(t, 12.3, Float64s{12.3}.Median())
	assert.Equal(t, 8.4, Float64s{12.3, 4.5}.Median())
	assert.Equal(t, 4.5, Float64s{2.1, 12.3, 4.5}.Median())
}

func TestFloat64s_Each(t *testing.T) {
	var values []float64

	values = []float64{}
	Float64s{}.Each(func(value float64) {
		values = append(values, value)
	})
	assert.Equal(t, []float64{}, values)

	values = []float64{}
	Float64s{435.34, 8923.1}.Each(func(value float64) {
		values = append(values, value)
	})
	assert.Equal(t, []float64{435.34, 8923.1}, values)
}

var float64sRandomTests = []struct {
	ss       Float64s
	expected float64
	source   rand.Source
}{
	{
		nil,
		0.0,
		nil,
	},
	{
		nil,
		0.0,
		rand.NewSource(0),
	},
	{
		Float64s{},
		0.0,
		rand.NewSource(0),
	},
	{
		Float64s{12.3, 2.34, 4.56},
		12.3,
		rand.NewSource(0),
	},
	{
		Float64s{12.3, 2.34, 4.56},
		4.56,
		rand.NewSource(1),
	},
	{
		Float64s{12.3},
		12.3,
		rand.NewSource(0),
	},
}

func TestFloat64s_Random(t *testing.T) {
	for _, test := range float64sRandomTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Random(test.source))
		})
	}
}

var float64sAbsTests = []struct {
	ss  Float64s
	abs Float64s
}{
	{
		Float64s{1, 2, 3, 4, 5},
		Float64s{1, 2, 3, 4, 5},
	},
	{
		Float64s{636, -5828, 444, -29281, 0},
		Float64s{636, 5828, 444, 29281, 0},
	},
	{
		Float64s{-584.2727, -47474.2112, 96843, -0.000004, 13},
		Float64s{584.2727, 47474.2112, 96843, 0.000004, 13},
	},
}

func TestFloat64s_Abs(t *testing.T) {
	for _, test := range float64sAbsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.abs, test.ss.Abs())
		})
	}
}

var float64sSendTests = []struct {
	ss            Float64s
	recieveDelay  time.Duration
	canceledDelay time.Duration
	expected      Float64s
}{
	{
		nil,
		0,
		0,
		nil,
	},
	{
		Float64s{1.2, 3.2},
		0,
		0,
		Float64s{1.2, 3.2},
	},
	{
		Float64s{1.2, 3.2},
		time.Millisecond * 30,
		time.Millisecond * 10,
		Float64s{1.2},
	},
	{
		Float64s{1.2, 3.2},
		time.Millisecond * 3,
		time.Millisecond * 10,
		Float64s{1.2, 3.2},
	},
}

func TestFloat64s_Send(t *testing.T) {
	for _, test := range float64sSendTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			ch := make(chan float64)
			actual := getFloat64sFromChan(ch, test.recieveDelay)
			ctx := createContextByDelay(test.canceledDelay)

			actualSended := test.ss.Send(ctx, ch)
			close(ch)

			assert.Equal(t, test.expected, actualSended)
			assert.Equal(t, test.expected, actual())
		})
	}
}
