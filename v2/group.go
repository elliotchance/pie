package pie

// Group returns a map of the value with an individual count.
func Group[T comparable](ss []T) map[T]int {
	group := map[T]int{}
	for _, n := range ss {
		group[n]++
	}

	return group
}
