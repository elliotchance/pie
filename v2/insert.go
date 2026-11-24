package pie

// Insert returns a new slice with values inserted at an index.
// An index past the end appends the values.
func Insert[T any](ss []T, index int, values ...T) []T {
	if index >= len(ss) {
		index = len(ss)
	}

	// Copy into separate storage, since values may share ss's backing array.
	result := make([]T, 0, len(ss)+len(values))
	if ss == nil && len(values) == 0 {
		result = nil
	}
	result = append(result, ss[:index]...)
	result = append(result, values...)
	return append(result, ss[index:]...)
}
