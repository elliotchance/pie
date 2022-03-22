package pie

import (
	"sort"
)

// SortStableUsing works similar to sort.SliceStable. However, unlike sort.SliceStable the
// slice returned will be reallocated as to not modify the input slice.
func SortStableUsing[T comparable](ss []T, less func(a, b T) bool) []T {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]T, len(ss))
	copy(sorted, ss)
	sort.SliceStable(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}
