package pie

// Last returns the last element or a zero value if there are no elements.
func Last[T any](ss []T) T {
	if len(ss) == 0 {
		var zeroValue T

		return zeroValue
	}

	return ss[len(ss)-1]
}
