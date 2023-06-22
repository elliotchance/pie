package pie

import (
	"golang.org/x/exp/constraints"
)

// Max is the maximum value, or zero.
func Max[T constraints.Ordered](ss []T) (max T) {
	if len(ss) == 0 {
		return
	}

	max = ss[0]
	for _, s := range ss {
		if s > max {
			max = s
		}
	}

	return
}
