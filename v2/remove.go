package pie

// Removes element at index idx from input slice.
// Returns true if succeeded, false if the index out of bounds.
func Remove[T any](ss *[]T, idx int) bool {
	if idx < 0 || idx >= len(*ss) {
		return false
	}
	*ss = append((*ss)[:idx], (*ss)[idx+1:]...)
	return true
}
