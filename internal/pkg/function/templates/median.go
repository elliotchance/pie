package templates

// Median returns the value separating the higher half from the lower half of a
// data sample.
//
// Zero is returned if there are no elements in the slice.
func (ss SliceType) Median() ElementType {
	l := len(ss)

	switch {
	case l == 0:
		return ElementZeroValue

	case l == 1:
		return ss[0]
	}

	sorted := ss.Sort()

	if l%2 != 0 {
		return sorted[l/2]
	}

	return (sorted[l/2-1] + sorted[l/2]) / 2
}
