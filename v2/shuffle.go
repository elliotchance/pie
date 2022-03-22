package pie

import (
	"math/rand"
)

// Shuffle returns a new shuffled slice by your rand.Source. The original slice
// is not modified.
func Shuffle[T any](ss []T, source rand.Source) []T {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 2 {
		return ss
	}

	shuffled := make([]T, n)
	copy(shuffled, ss)

	rnd := rand.New(source)
	rnd.Shuffle(n, func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}
