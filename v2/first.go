package pie

// First returns the first element or a zero value if there are no elements.
func First[T any](ss []T) T {
	if len(ss) == 0 {
		var zeroValue T

		return zeroValue
	}

	return ss[0]
}
