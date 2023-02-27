package pie

// GroupBy groups slice elements by key returned by getKey function for each
// slice element.
//
// It returns a map in which slices of elements of original slice are matched
// to keys defined by getKey function. It returns non-nil map, if empty or nil
// slice is passed.
//
// For example, if you want to group integers by their remainder from division
// by 5, you can use this function as follows:
//
//	_ = pie.GroupBy(
//	    []int{23, 76, 37, 11, 23, 47},
//	    func(num int) int {
//	        return num % 5
//	    },
//	)
//
// In above case map {1:[76, 11], 2:[37, 47], 3:[23, 23]} is returned.
func GroupBy[T comparable, U any](values []U, getKey func(U) T) map[T][]U {
	groups := make(map[T][]U)

	for _, val := range values {
		key := getKey(val)
		groups[key] = append(groups[key], val)
	}

	return groups
}
