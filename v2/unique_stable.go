package pie

// UniqueStable works similar to Unique. However, unlike Unique
// the slice returned will be in previous relative order
func UniqueStable[T comparable](ss []T) []T {
	// Avoid the allocation. If there is one element or less it is already
	// unique.
	if len(ss) < 2 {
		return ss
	}

	seen := map[T]struct{}{}
	ret := make([]T, 0)

	for _, value := range ss {
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		ret = append(ret, value)
	}

	return ret
}
