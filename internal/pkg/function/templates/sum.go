package templates

// Sum is the sum of all of the elements.
func (ss SliceType) Sum() (sum ElementType) {
	for _, s := range ss {
		sum += s
	}

	return
}
