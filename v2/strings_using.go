package pie

// StringsUsing transforms each element to a string.
func StringsUsing[T any](ss []T, transform func(T) string) []string {
	return Map(ss, transform)
}
