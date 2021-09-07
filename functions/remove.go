package functions

// Remove returns a new slice that does not include any of the items.
func (ss SliceType) Remove(items ...ElementType) (result SliceType) {
	if len(items) == 0 {
		return items
	}

	ss2 := make(map[ElementType]bool, len(items))
	for _, item := range items {
		ss2[item] = true
	}

	result = SliceType{}
	for _, v := range ss {
		if !ss2[v] {
			result = append(result, v)
		}
	}
	return
}
