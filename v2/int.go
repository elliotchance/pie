package pie

import (
	"golang.org/x/exp/constraints"
)

// Int transforms a value into an int. This should only be used on slices
// that resolve to strings that represent numbers. An invalid value will use
// zero and fractional values will be truncated.
func Int[T constraints.Ordered](x T) int {
	return int(Float64(x))
}
