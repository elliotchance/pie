package functions

//
func (ss SliceType) First() ElementType {
	return ss.FirstOr(ElementZeroValue)
}
