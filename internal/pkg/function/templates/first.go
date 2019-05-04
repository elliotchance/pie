package templates

// First returns the first element, or zero. Also see FirstOr().
func (ss SliceType) First() ElementType {
	return ss.FirstOr(ElementZeroValue)
}
