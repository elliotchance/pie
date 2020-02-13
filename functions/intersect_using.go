package functions

// IntersectUsing returns items that exist in all lists, using equals function
//
// It returns slice without any duplicates.
// If zero slice arguments are provided, then nil is returned.
func (ss SliceType) IntersectUsing(equals func(ElementType, ElementType) (bool, ElementType), slices ...SliceType) (ss2 SliceType) {
	if slices == nil {
		return nil
	}

	found := map[ElementType]int{}

	for _, e1 := range ss {
		for _, s2 := range slices {
			foundInSlice := false
			for _, e2 := range s2 {
				chekFound, checkValue := equals(e1, e2)
				if chekFound {
					found[checkValue]++
					foundInSlice = true
					break // if found the element don't check other elements in this slice
				}
			}
			if !foundInSlice {
				break // if not found in this slice don't check other slices
			}
		}
	}
	ss2 = SliceType{}

	for value, count := range found {
		if count == len(slices) {
			ss2 = ss2.Append(value)
		}
	}
	return ss2
}
