package pie

// Omit returns a new map without the given keys.
func Omit[K comparable, V any](keys []K, m map[K]V) map[K]V {
	keysLength := len(keys)
	mapLength := len(m)

	if keysLength == 0 || mapLength == 0 {
		return m
	}

	result := make(map[K]V, len(keys))

	for key, value := range m {
		if !containsKey(keys, key) {
			result[key] = value
		}
	}

	return result
}

func containsKey[K comparable](keys []K, key K) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}
