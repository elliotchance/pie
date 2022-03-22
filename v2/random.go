package pie

import (
	"golang.org/x/exp/constraints"
	"math/rand"
)

// Random returns a random element by your rand.Source, or zero.
func Random[T constraints.Integer | constraints.Float](ss []T, source rand.Source) T {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 1 {
		return 0
	}

	if n < 2 {
		return ss[0]
	}

	rnd := rand.New(source)
	i := rnd.Intn(n)

	return ss[i]
}
