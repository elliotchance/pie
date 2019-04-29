package functions

import (
	"math/rand"
)

// Shuffle returns shuffled slice by your rand.Source
func (ss SliceType) Shuffle(source rand.Source) SliceType {
	if len(ss) < 2 {
		return ss
	}

	shuffled := make([]ElementType, len(ss))
	copy(shuffled, ss)

	rnd := rand.New(source)
	rnd.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}
