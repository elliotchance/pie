package pie_test

import (
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"testing"
)

var intsContainsTests = []struct {
	ss       pie.Ints
	contains int
	expected bool
}{
	{nil, 1, false},
	{pie.Ints{1, 2, 3}, 1, true},
	{pie.Ints{1, 2, 3}, 2, true},
	{pie.Ints{1, 2, 3}, 3, true},
	{pie.Ints{1, 2, 3}, 4, false},
	{pie.Ints{1, 2, 3}, 5, false},
	{pie.Ints{1, 2, 3}, 6, false},
	{pie.Ints{1, 5, 3}, 5, true},
}

func TestInts_Contains(t *testing.T) {
	for _, test := range intsContainsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, test.ss.Contains(test.contains))
		})
	}
}

func TestIntsContains(t *testing.T) {
	for _, test := range intsContainsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.IntsContains(test.ss, test.contains))
		})
	}
}

var intsOnlyAndWithoutTests = []struct {
	ss                pie.Ints
	condition         func(int) bool
	expectedOnly      pie.Ints
	expectedWithout   pie.Ints
	expectedTransform pie.Ints
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
		pie.Ints{1, 2, 3},
		func(s int) bool {
			return s != 2
		},
		pie.Ints{1, 3},
		pie.Ints{2},
		pie.Ints{6, 7, 8},
	},
}

func TestInts_Only(t *testing.T) {
	for _, test := range intsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedOnly, test.ss.Only(test.condition))
		})
	}
}

func TestIntsOnly(t *testing.T) {
	for _, test := range intsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, []int(test.expectedOnly), pie.IntsOnly(test.ss, test.condition))
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

func TestIntsWithout(t *testing.T) {
	for _, test := range intsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, []int(test.expectedWithout), pie.IntsWithout(test.ss, test.condition))
		})
	}
}

func TestInts_Transform(t *testing.T) {
	for _, test := range intsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedTransform, test.ss.Transform(pie.AddInt(5)))
		})
	}
}

func TestIntsTransform(t *testing.T) {
	for _, test := range intsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, []int(test.expectedTransform), pie.IntsTransform(test.ss, pie.AddInt(5)))
		})
	}
}

var intsFirstAndLastTests = []struct {
	ss             pie.Ints
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
		pie.Ints{100},
		100,
		100,
		100,
		100,
	},
	{
		pie.Ints{1, 2},
		1,
		1,
		2,
		2,
	},
	{
		pie.Ints{1, 2, 3},
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

func TestIntsFirstOr(t *testing.T) {
	for _, test := range intsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.firstOr, pie.IntsFirstOr(test.ss, 102))
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

func TestIntsLastOr(t *testing.T) {
	for _, test := range intsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.lastOr, pie.IntsLastOr(test.ss, 202))
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

func TestIntsFirst(t *testing.T) {
	for _, test := range intsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.first, pie.IntsFirst(test.ss))
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

func TestIntsLast(t *testing.T) {
	for _, test := range intsFirstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.last, pie.IntsLast(test.ss))
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

func TestIntsMin(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.min, pie.IntsMin(test.ss))
		})
	}
}

func TestInts_Min(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.min, pie.Ints(test.ss).Min())
		})
	}
}

func TestIntsMax(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.max, pie.IntsMax(test.ss))
		})
	}
}

func TestInts_Max(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.max, pie.Ints(test.ss).Max())
		})
	}
}

func TestIntsSum(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.sum, pie.IntsSum(test.ss))
		})
	}
}

func TestInts_Sum(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.sum, pie.Ints(test.ss).Sum())
		})
	}
}

func TestInts_Len(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.len, pie.Ints(test.ss).Len())
		})
	}
}

func TestIntsAverage(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.average, pie.IntsAverage(test.ss))
		})
	}
}

func TestInts_Average(t *testing.T) {
	for _, test := range intsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.average, pie.Ints(test.ss).Average())
		})
	}
}
