package functions

// Bottom will return n elements from bottom
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss SliceType) Bottom(n int) (top SliceType) {
	var lastIndex = len(ss) - 1
	for i := lastIndex; i > -1 && n > 0; i-- {
		top = append(top, ss[i])
		n--
	}

	return
}
