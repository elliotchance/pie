package templates

// Last returns the last element, or zero. Also see LastOr().
func (ss SliceType) Last() ElementType {
	return ss.LastOr(ElementZeroValue)
}
