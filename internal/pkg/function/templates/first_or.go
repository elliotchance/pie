package templates

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss SliceType) FirstOr(defaultValue ElementType) ElementType {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}
