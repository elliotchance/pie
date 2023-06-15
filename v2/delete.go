package pie

// Removes element at index idx from input slice, returns resulting slice.
// If the index out of bounds, returns unchanged input slice.
func Delete[T any](ss []T, idx int) []T {
	if idx < 0 || idx >= len(ss) {
		return ss
	}
	return append(ss[:idx], ss[idx+1:]...)
}
