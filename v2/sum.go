package pie

import "golang.org/x/exp/constraints"

// Sum is the sum of all of the elements.
func Sum[T constraints.Integer | constraints.Float](ss []T) (sum T) {
	for _, s := range ss {
		sum += s
	}

	return
}
