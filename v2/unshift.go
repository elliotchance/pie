package pie

// Unshift adds one or more elements to the beginning of the slice
// and returns the new slice.
func Unshift[T any](ss []T, elements ...T) (unshift []T) {
	unshift = append([]T{}, elements...)
	unshift = append(unshift, ss...)

	return
}
