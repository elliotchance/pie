package pie

import "golang.org/x/exp/constraints"

// Shift will return two values: the shifted value and the rest slice.
func Shift[T constraints.Integer | constraints.Float](ss []T) (T, []T) {
	return FirstOr(ss, 0), DropTop(ss, 1)
}
