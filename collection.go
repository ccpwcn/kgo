package kgo

type Collection[T comparable] interface {
	Add(item ...T)
	Remove(item T) bool
	Clear()
	Contains(item T) bool
	Len() int
	Empty() bool
}
