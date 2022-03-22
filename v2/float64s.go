package pie

import (
	"golang.org/x/exp/constraints"
)

// Float64s transforms each element to a float64.
func Float64s[T constraints.Ordered](ss []T) []float64 {
	return Map(ss, Float64[T])
}
