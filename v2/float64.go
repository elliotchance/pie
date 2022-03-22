package pie

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

// Float64 transforms a value into a float64. This should only be used on slices
// that resolve to strings that represent numbers. An invalid value will use
// zero.
func Float64[T constraints.Ordered](x T) float64 {
	f, _ := strconv.ParseFloat(String(x), 64)

	return f
}
