package functions

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func (ss SliceType) All(fn func(value ElementType) bool) bool {
	for _, value := range ss {
		if !fn(value) {
			return false
		}
	}

	return true
}
