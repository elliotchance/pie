package pie

// Keys returns the keys in the map. All of the items will be unique.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func Keys[K comparable, V any](m map[K]V) []K {
	// Avoid allocation
	l := len(m)
	if l == 0 {
		return nil
	}

	i := 0
	keys := make([]K, len(m))
	for key := range m {
		keys[i] = key
		i++
	}

	return keys
}
