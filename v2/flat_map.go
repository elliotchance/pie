package pie

// FlatMap maps each element into a new slice, then flattens the result.
func FlatMap[T any, U any](ss []T, fn func(T) []U) []U {
	if ss == nil {
		return nil
	}

	return Flat(Map(ss, fn))
}
