package pie

// Equals compare elements from the start to the end,
//
// if they are the same is considered the slices are equal if all elements are
// the same is considered the slices are equal
// if each slice == nil is considered that they're equal
//
// if element realizes Equals interface it uses that method, in other way uses
// default compare
func Equals[T comparable](ss []T, rhs []T) bool {
	if len(ss) != len(rhs) {
		return false
	}

	for i := range ss {
		if ss[i] != rhs[i] {
			return false
		}
	}

	return true
}
