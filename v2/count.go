package pie

// Count returns the number of elements in the slice that are equal to
// lookingFor. It is a more concise alternative to filtering the slice and
// taking its length.
//
// When using slices of pointers it will only compare by address, not value.
func Count[T comparable](ss []T, lookingFor T) int {
	count := 0
	for _, s := range ss {
		if s == lookingFor {
			count++
		}
	}

	return count
}
