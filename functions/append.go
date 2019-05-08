package functions

// Append will return a new slice with the elements appended to the end.
//
// It is acceptable to provide zero arguments.
func (ss SliceType) Append(elements ...ElementType) SliceType {
	// Copy ss, to make sure no memory is overlapping between input and
	// output. See issue #97.
	result := append(SliceType{}, ss...)

	result = append(result, elements...)
	return result
}
