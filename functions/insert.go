package functions

// Insert a value at an index
func (ss SliceType) Insert(index int, values ...ElementType) SliceType {
	if index >= ss.Len() {
		return SliceType.Extend(ss, SliceType(values))
	}

	return SliceType.Extend(ss[:index], SliceType(values), ss[index:])
}
