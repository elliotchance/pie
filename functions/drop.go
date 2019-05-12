package functions

// Drop will return the rest slice after dropping the first n elements
// if the slice has less elements then n that'll return empty slice
// if n <= 0 it'll return all copied elements.
func (ss SliceType) Drop(n int) (drop SliceType) {
	if n <= 0 {
		drop = make(SliceType, len(ss))
		copy(drop, ss)
		return
	}

	for i := n; i < len(ss) && n > 0; i++ {
		drop = append(drop, ss[i])
		n--
	}

	return
}
