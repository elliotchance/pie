package functions

//
func (ss SliceType) Append(elements ...ElementType) SliceType {
	return append(ss, elements...)
}
