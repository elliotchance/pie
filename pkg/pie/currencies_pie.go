package pie

// Keys returns the keys in the map. All of the items will be unique.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func (m currencies) Keys() []string {
	// Avoid allocation
	l := len(m)
	if l == 0 {
		return nil
	}

	i := 0
	keys := make([]string, len(m))
	for key := range m {
		keys[i] = key
		i++
	}

	return keys
}

// Values returns the values in the map.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func (m currencies) Values() []currency {
	// Avoid allocation
	l := len(m)
	if l == 0 {
		return nil
	}

	i := 0
	keys := make([]currency, len(m))
	for _, value := range m {
		keys[i] = value
		i++
	}

	return keys
}
