package pie

import (
	"math/rand"
	"testing"
)

func BenchmarkIntMedianSmall(b *testing.B) { benchmarkIntMedian(b, 20) }

// func BenchmarkIntMedianMedium(b *testing.B) { benchmarkIntMedian(b, 60) }

func BenchmarkIntMedianMedium(b *testing.B) { benchmarkIntMedian(b, 800) }
func BenchmarkIntMedianLarge(b *testing.B)  { benchmarkIntMedian(b, 1000000) }

func benchmarkIntMedian(b *testing.B, size int) {
	a := make(Ints, size)
	for i := range a {
		// As many possible values as slots in the slice.
		// Variety, with some duplicates.
		a[i] = rand.Intn(size)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// m := a.MedianOld()
		// m := a.Median()
		m := a.median()
		sink += m
	}
}

// Prevent compiler from agressively optimizing away the result
var sink = 0
