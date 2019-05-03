package functions

import (
	"math/rand"
)

// Random returns a random element by your rand.Source, or zero
func (ss SliceType) Random(source rand.Source) ElementType {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 1 {
		return ElementZeroValue
	}
	if n < 2 {
		return ss[0]
	}
	rnd := rand.New(source)
	i := rnd.Intn(n)
	return ss[i]
}
