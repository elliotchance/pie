package pie

import (
	"math/rand"
	"testing"
)

func BenchmarkFloatMedianSmall(b *testing.B) { benchmarkFloatMedian(b, 20) }

func BenchmarkFloatMedianMedium(b *testing.B) { benchmarkFloatMedian(b, 800) }
func BenchmarkFloatMedianLarge(b *testing.B)  { benchmarkFloatMedian(b, 1000000) }

func benchmarkFloatMedian(b *testing.B, size int) {
	// Make the random numbers below deterministic
	rand.Seed(123)

	a := make(Float64s, size)
	for i := range a {
		// As many possible values as slots in the slice.
		// Positives and negatives.
		// Variety, with some duplicates.
		a[i] = -0.5 + rand.Float64()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// m := a.MedianOld()
		// m := a.medianCheck()
		m := a.Median()
		sinkFloats += m
	}
}

// Prevent compiler from agressively optimizing away the result
var sinkFloats float64
