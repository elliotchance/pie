package functions

// DropTop will return the rest slice after dropping the top n elements
// if the slice has less elements then n that'll return empty slice
// if n < 0 it'll return empty slice.
func (ss SliceType) DropTop(n int) (drop SliceType) {
	if n < 0 {
		return
	}

	for i := n; i < len(ss); i++ {
		drop = append(drop, ss[i])
	}

	return
}
