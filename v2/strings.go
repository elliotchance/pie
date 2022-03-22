package pie

import (
	"golang.org/x/exp/constraints"
)

// Strings transforms each element to a string.
//
// If the element type implements fmt.Stringer it will be used. Otherwise it
// will fallback to the result of:
//
//   fmt.Sprintf("%v")
//
func Strings[T constraints.Ordered](ss []T) []string {
	return Map(ss, String[T])
}
