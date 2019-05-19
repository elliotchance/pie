package functions

// NotEquals  compare elements of slice
// and return true if they are not equal
//
// if element realizes Equals interface it uses that method, in other way uses default compare
func (ss SliceType) NotEquals(rhs SliceType) bool {
	// It's been done for generator, see issue #143
	var eq = ss.Equals

	return !eq(rhs)
}
