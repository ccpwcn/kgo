package kgo

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable](items ...T) *Set[T] {
	s := &Set[T]{
		m: make(map[T]struct{}),
	}
	s.Add(items...)
	return s
}

func (s *Set[T]) Add(item ...T) {
	for _, item := range item {
		s.m[item] = struct{}{}
	}
}
func (s *Set[T]) Remove(item T) bool {
	if s.Contains(item) {
		delete(s.m, item)
		return true
	} else {
		return false
	}
}

func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Empty() bool {
	return len(s.m) == 0
}
