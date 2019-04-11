package functions

//
func (ss SliceType) Last() ElementType {
	return ss.LastOr(ElementZeroValue)
}
