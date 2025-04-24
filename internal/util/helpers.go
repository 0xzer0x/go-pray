package util

func MapKeys[K comparable, V any](mp map[K]V) []K {
	keys := make([]K, len(mp))

	i := 0
	for k := range mp {
		keys[i] = k
		i++
	}
	return keys
}

func FindInMap[K comparable, V comparable](mp map[K]V, value V) K {
	var key K
	var val V
	for key, val = range mp {
		if val == value {
			break
		}
	}
	return key
}
