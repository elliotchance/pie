package functions

import "math/rand"

func (ss SliceType) Shuffle() SliceType {
	for i := len(ss) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}
