package pie

import "golang.org/x/exp/constraints"

// Product is the product of all of the elements.
func Product[T constraints.Integer | constraints.Float](ss []T) (product T) {
	if len(ss) == 0 {
		return
	}

	product = ss[0]
	for _, s := range ss[1:] {
		product *= s
	}

	return
}
