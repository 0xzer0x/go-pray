package util

func MapKeys[K comparable, V any](mp map[K]V) []K {
	var keys []K = make([]K, len(mp))

	i := 0
	for k := range mp {
		keys[i] = k
		i++
	}
	return keys
}
