package functions

// LastOr returns the last element or a default value if there are no elements.
func (ss SliceType) LastOr(defaultValue ElementType) ElementType {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}
