package pie

// Diff returns the elements that needs to be added or removed from the first
// slice to have the same elements in the second slice.
//
// The order of elements is not taken into consideration, so the slices are
// treated sets that allow duplicate items.
//
// The added and removed returned may be blank respectively, or contain upto as
// many elements that exists in the largest slice.
func Diff[T comparable](ss []T, against []T) (added, removed []T) {
	diffOneWay := func(ss1, ss2raw []T) (result []T) {
		set := make(map[T]struct{}, len(ss1))

		for _, s := range ss1 {
			set[s] = struct{}{}
		}

		for _, s := range ss2raw {
			if _, ok := set[s]; ok {
				delete(set, s) // remove duplicates
			} else {
				result = append(result, s)
			}
		}
		return
	}

	added = diffOneWay(ss, against)
	removed = diffOneWay(against, ss)

	return
}
