package pie

// Flat flattens the two-dimensional slice into one-dimensional slice.
func Flat[T any](ss [][]T) (ss2 []T) {
	if ss == nil {
		return nil
	}

	ss2 = make([]T, 0)
	for _, s := range ss {
		ss2 = append(ss2, s...)
	}

	return
}
