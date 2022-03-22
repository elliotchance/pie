package pie

import (
	"golang.org/x/exp/constraints"
)

// Ints transforms each element to an integer.
func Ints[T constraints.Ordered](ss []T) []int {
	return Map(ss, Int[T])
}
