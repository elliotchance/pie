package functions

// Unshift adds one or more elements to the beginning of the slice
// and returns the new slice.
func (ss SliceType) Unshift(elements ...ElementType) (unshift SliceType) {
	unshift = append(SliceType{}, elements...)
	unshift = append(unshift, ss...)

	return
}
