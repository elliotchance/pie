package pie

import (
	"math/rand"
	"testing"
)

func BenchmarkIntMedianSmall(b *testing.B) { benchmarkIntMedian(b, 20) }

func BenchmarkIntMedianMedium(b *testing.B) { benchmarkIntMedian(b, 800) }
func BenchmarkIntMedianLarge(b *testing.B)  { benchmarkIntMedian(b, 1000000) }

func benchmarkIntMedian(b *testing.B, size int) {
	// Make the random numbers below deterministic
	rand.Seed(123)

	a := make(Ints, size)
	for i := range a {
		// As many possible values as slots in the slice.
		// Negatives and positives.
		// Variety, with some duplicates.
		a[i] = -size/2 + rand.Intn(size)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := a.Median()
		sinkInts += m
	}
}

// Prevent compiler from agressively optimizing away the result
var sinkInts int
