package pie

import "golang.org/x/exp/constraints"

// Abs returns the absolute value.
func Abs[T constraints.Integer | constraints.Float](val T) T {
	if val < 0 {
		return -val
	}

	return val
}
