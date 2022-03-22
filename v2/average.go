package pie

import "golang.org/x/exp/constraints"

// Average is the average of all of the elements, or zero if there are no
// elements.
func Average[T constraints.Integer | constraints.Float](ss []T) float64 {
	if l := len(ss); l > 0 {
		return float64(Sum(ss)) / float64(l)
	}

	return 0
}
