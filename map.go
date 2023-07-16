package kgo

func HasKey[K comparable, V any](m map[K]V, k K) bool {
	if _, ok := m[k]; ok {
		return true
	} else {
		return false
	}
}

// MapKeys 取得Map的Key的集合
func MapKeys[K comparable, V any](m map[K]V) []K {
	i := 0
	keys := make([]K, len(m))
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
