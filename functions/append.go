package functions

// Append will return a new slice with the elements appended to the end.
//
// It is acceptable to provide zero arguments.
func (ss SliceType) Append(elements ...ElementType) (result SliceType) {
	return append(append(result, ss...), elements...)
}
