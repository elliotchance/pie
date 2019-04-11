package functions

//
func (ss SliceType) FirstOr(defaultValue ElementType) ElementType {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}
