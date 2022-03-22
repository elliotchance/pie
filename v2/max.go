package pie

import (
	"golang.org/x/exp/constraints"
)

// Max is the maximum value, or zero.
func Max[T constraints.Ordered](ss []T) (min T) {
	if len(ss) == 0 {
		return
	}

	min = ss[0]
	for _, s := range ss {
		if s > min {
			min = s
		}
	}

	return
}
