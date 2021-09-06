package functions

// Remove items from slice when item existed
func (ss SliceType) Remove(items ...ElementType) (result SliceType, removedCnt int) {
	result = SliceType{}
	for _, v := range ss {
		found := false
		for _, i := range items {
			if i == v {
				found = true
				break
			}
		}
		if !found {
			result = append(result, v)
		} else {
			removedCnt++
		}
	}
	return
}
