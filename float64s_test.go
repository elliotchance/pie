package pie_test

import (
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"testing"
)

var float64sContainsTests = []struct {
	ss       pie.Float64s
	contains float64
	expected bool
}{
	{nil, 1, false},
	{pie.Float64s{1, 2, 3}, 1, true},
	{pie.Float64s{1, 2, 3}, 2, true},
	{pie.Float64s{1, 2, 3}, 3, true},
	{pie.Float64s{1, 2, 3}, 4, false},
	{pie.Float64s{1, 2, 3}, 5, false},
	{pie.Float64s{1, 2, 3}, 6, false},
	{pie.Float64s{1, 5, 3}, 5, true},
}

func TestFloat64s_Contains(t *testing.T) {
	for _, test := range float64sContainsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.expected, test.ss.Contains(test.contains))
		})
	}
}

var float64sOnlyAndWithoutTests = []struct {
	ss                pie.Float64s
	condition         func(float64) bool
	expectedOnly      pie.Float64s
	expectedWithout   pie.Float64s
	expectedTransform pie.Float64s
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
		pie.Float64s{1, 2, 3},
		func(s float64) bool {
			return s != 2
		},
		pie.Float64s{1, 3},
		pie.Float64s{2},
		pie.Float64s{6.2, 7.2, 8.2},
	},
}

func TestFloat64s_Only(t *testing.T) {
	for _, test := range float64sOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.expectedOnly, test.ss.Only(test.condition))
		})
	}
}

func TestFloat64s_Without(t *testing.T) {
	for _, test := range float64sOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.expectedWithout, test.ss.Without(test.condition))
		})
	}
}

func TestFloat64s_Transform(t *testing.T) {
	for _, test := range float64sOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.expectedTransform, test.ss.Transform(pie.AddFloat64(5.2)))
		})
	}
}

var float64sFirstAndLastTests = []struct {
	ss             pie.Float64s
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
		pie.Float64s{100},
		100,
		100,
		100,
		100,
	},
	{
		pie.Float64s{1, 2},
		1,
		1,
		2,
		2,
	},
	{
		pie.Float64s{1, 2, 3},
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
	ss            pie.Float64s
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
			assert.Equal(t, test.min, pie.Float64s(test.ss).Min())
		})
	}
}

func TestFloat64s_Max(t *testing.T) {
	for _, test := range float64sStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.max, pie.Float64s(test.ss).Max())
		})
	}
}

func TestFloat64s_Sum(t *testing.T) {
	for _, test := range float64sStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.sum, pie.Float64s(test.ss).Sum())
		})
	}
}

func TestFloat64s_Len(t *testing.T) {
	for _, test := range float64sStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.len, pie.Float64s(test.ss).Len())
		})
	}
}

func TestFloat64s_Average(t *testing.T) {
	for _, test := range float64sStatsTests {
		t.Run("", func(t *testing.T) {
			defer assertImmutableFloat64s(t, &test.ss)()
			assert.Equal(t, test.average, pie.Float64s(test.ss).Average())
		})
	}
}

var float64sJSONTests = []struct {
	ss         pie.Float64s
	jsonString string
}{
	{
		nil,
		`[]`, // Instead of null.
	},
	{
		pie.Float64s{},
		`[]`,
	},
	{
		pie.Float64s{12.3},
		`[12.3]`,
	},
	{
		pie.Float64s{23, -2.5, 3424, 12.3},
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
	ss       pie.Float64s
	sorted   pie.Float64s
	reversed pie.Float64s
	areSorted bool
}{
	{
		nil,
		nil,
		nil,
		true,
	},
	{
		pie.Float64s{},
		pie.Float64s{},
		pie.Float64s{},
		true,
	},
	{
		pie.Float64s{789},
		pie.Float64s{789},
		pie.Float64s{789},
		true,
	},
	{
		pie.Float64s{12.789, -13.2, 789},
		pie.Float64s{-13.2, 12.789, 789},
		pie.Float64s{789, -13.2, 12.789},
		false,
	},
	{
		pie.Float64s{12.789, -13.2, 1.234e6, 789},
		pie.Float64s{-13.2, 12.789, 789, 1.234e6},
		pie.Float64s{789, 1.234e6, -13.2, 12.789},
		false,
	},
	{
		pie.Float64s{-13.2, 12.789},
		pie.Float64s{-13.2, 12.789},
		pie.Float64s{12.789, -13.2},
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
