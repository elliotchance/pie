package templates

// Values returns the values in the map.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func (m MapType) Values() []ElementType {
	// Avoid allocation
	l := len(m)
	if l == 0 {
		return nil
	}

	i := 0
	keys := make([]ElementType, len(m))
	for _, value := range m {
		keys[i] = value
		i++
	}

	return keys
}
