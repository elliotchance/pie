package pie

// Insert a value at an index.
func Insert[T any](ss []T, index int, values ...T) []T {
	if index >= len(ss) {
		return append(ss, values...)
	}

	return append(ss[:index], append(values, ss[index:]...)...)
}
