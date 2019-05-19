package functions

// Equals compare elements of slice
//
// if all elements the same is considered that slices are equal
// if len(first_slice) != len(second_slice) they are not equal
// if slices == nil is considered that they're equal
//
// if element realizes Equals interface it uses that method, in other way uses default compare
func (ss SliceType) Equals(rhs SliceType) bool {
	if len(ss) != len(rhs) {
		return false
	}

	for i := range ss {
		if !ss[i].Equals(rhs[i]) {
			return false
		}
	}

	return true
}
