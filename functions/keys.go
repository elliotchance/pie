package functions

//
func (m MapType) Keys() KeySliceType {
	// Avoid allocation
	l := len(m)
	if l == 0 {
		return nil
	}

	i := 0
	keys := make(KeySliceType, len(m))
	for key := range m {
		keys[i] = key
		i++
	}

	return keys
}
