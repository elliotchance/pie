package pie

// Any will return true if any callbacks return true. It follows the same logic
// as the any() function in Python.
//
// If the list is empty then false is always returned.
func Any[T any](ss []T, fn func(value T) bool) bool {
	for _, value := range ss {
		if fn(value) {
			return true
		}
	}

	return false
}
