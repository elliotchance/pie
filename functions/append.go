package functions

// Append will return a new slice with the elements appended to the end. It is a
// wrapper for the internal append(). It is offered as a function so that it can
// more easily chained.
//
// It is acceptable to provide zero arguments.
func (ss SliceType) Append(elements ...ElementType) SliceType {
	return append(ss, elements...)
}
