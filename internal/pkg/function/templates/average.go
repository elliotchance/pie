package templates

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss SliceType) Average() float64 {
	if l := ElementType(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}
