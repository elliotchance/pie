package functions

//
func (ss SliceType) Each(fn func(ElementType)) SliceType {
	for _, s := range ss {
		fn(s)
	}

	return ss
}
