package kgo

func HasKey[K comparable, V any](m map[K]V, k K) bool {
	if _, ok := m[k]; ok {
		return true
	} else {
		return false
	}
}
