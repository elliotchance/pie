package pie

import "sort"

// Removes elements at indices in idx from input slice, returns resulting slice.
// If an index is out of bounds, skip it.
func Delete[T any](ss []T, idx ...int) []T {
	// short path O(n) inplace
	if len(idx) == 1 {
		i := idx[0]

		if i < 0 || i >= len(ss) {
			return ss
		}
		return append(ss[:i], ss[i+1:]...)
	}

	// long path O(mLog(m) + n)
	sort.Ints(idx)

	ss2 := make([]T, 0, len(ss))

	prev := 0
	for _, i := range idx {
		if i < 0 || i >= len(ss) {
			continue
		}
		// Copy by consecutive chunks instead of one by one
		ss2 = append(ss2, ss[prev:i]...)
		prev = i + 1
	}
	ss2 = append(ss2, ss[prev:]...)

	return ss2
}
