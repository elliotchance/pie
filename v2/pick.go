package pie

// Pick returns a new map with the given keys.
func Pick[K comparable, V any](keys []K, m map[K]V) map[K]V {
	keysLength := len(keys)
	mapLength := len(m)

	if keysLength == 0 || mapLength == 0 {
		return make(map[K]V, 0)
	}

	result := make(map[K]V, len(keys))

	for _, key := range keys {
		if value, ok := m[key]; ok {
			result[key] = value
		}
	}

	return result
}
