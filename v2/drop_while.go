package pie

// Drop items from the slice while f(item) is true.
// Afterwards, return every element until the slice is empty. It follows the
// same logic as the dropwhile() function from itertools in Python.
func DropWhile[T comparable](ss []T, f func(s T) bool) (ss2 []T) {
	ss2 = make([]T, len(ss))
	copy(ss2, ss)
	for i, value := range ss2 {
		if !f(value) {
			return ss2[i:]
		}
	}

	return []T{}
}
