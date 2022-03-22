package pie

// Contains returns true if the element exists in the slice.
//
// When using slices of pointers it will only compare by address, not value.
func Contains[T comparable](ss []T, lookingFor T) bool {
	for _, s := range ss {
		if lookingFor == s {
			return true
		}
	}

	return false
}
