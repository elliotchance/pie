package pie

// FindFirstUsing will return the index of the first element when the callback
// returns true or -1 if no element is found.
// It follows the same logic as the findIndex() function in Javascript.
//
// If the list is empty then -1 is always returned.
func FindFirstUsing[T any](ss []T, fn func(value T) bool) int {
	for idx, value := range ss {
		if fn(value) {
			return idx
		}
	}

	return -1
}
