package kgo

import (
	"encoding/json"
)

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

func (s *Set[T]) Add(items ...T) {
	for _, item := range items {
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
func (s *Set[T]) ToArray() []T {
	var items = make([]T, 0, s.Len())
	for k := range s.m {
		items = append(items, k)
	}
	return items
}

func (s *Set[T]) MarshalText() (data []byte, err error) {
	return json.Marshal(s.ToArray())
}

func (s *Set[T]) UnmarshalText(data []byte) (err error) {
	var items []T
	err = json.Unmarshal(data, &items)
	s.m = make(map[T]struct{})
	s.Add(items...)
	return err
}
