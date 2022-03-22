package pie

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func All[T any](ss []T, fn func(value T) bool) bool {
	for _, value := range ss {
		if !fn(value) {
			return false
		}
	}

	return true
}
