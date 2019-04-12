package functions

// AreUnique will return true if the slice contains elements that are all
// different (unique) from each other.
func (ss SliceType) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
}
