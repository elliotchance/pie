package pie

// AreUnique will return true if the slice contains elements that are all
// different (unique) from each other.
func AreUnique[T comparable](ss []T) bool {
	return len(Unique(ss)) == len(ss)
}
