package pie

// FirstOr returns the first element or a default value if there are no
// elements.
func FirstOr[T any](ss []T, defaultValue T) T {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}
