package functions

import (
	"math/rand"
	"time"
)

// Max is the maximum value, or zero.
func (ss SliceType) Shuffle() SliceType {
	if len(ss) < 2 {
		return ss
	}

	shuffled := make([]ElementType, len(ss))
	copy(shuffled, ss)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}
