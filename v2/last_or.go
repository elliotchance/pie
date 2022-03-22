package pie

// LastOr returns the last element or a default value if there are no elements.
func LastOr[T any](ss []T, defaultValue T) T {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}
