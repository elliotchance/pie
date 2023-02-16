package pie

// Flat flattens the two-dimensional slice into one-dimensional slice.
func Flat[T any](ss [][]T) (ss2 []T) {
	for _, s := range ss {
		ss2 = append(ss2, s...)
	}

	return ss2
}
