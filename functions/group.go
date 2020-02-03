package functions

// Group returns a map of the value with an individual count.
//
func (ss SliceType) Group() map[ElementType]int {
	group := map[ElementType]int{}
	for _, n := range ss {
		group[n]++
	}
	return group
}
