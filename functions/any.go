package functions

//
func (ss SliceType) Any(fn func(value ElementType) bool) bool {
	for _, value := range ss {
		if fn(value) {
			return true
		}
	}

	return false
}
