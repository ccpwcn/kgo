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

// MarshalJSON 实现 json.Marshaler。
// 之前用 MarshalText 承载 JSON 是错误的：encoding/json 会把 TextMarshaler 的输出
// 当作一个字符串再转义，导致 Set 嵌套在结构体中被序列化成 "[1,2,3]"（字符串）
// 而非真正的 JSON 数组 [1,2,3]。改用 json.Marshaler 后输出为原生数组。
func (s *Set[T]) MarshalJSON() (data []byte, err error) {
	return json.Marshal(s.ToArray())
}

// UnmarshalJSON 实现 json.Unmarshaler，从 JSON 数组反序列化回 Set。
func (s *Set[T]) UnmarshalJSON(data []byte) (err error) {
	var items []T
	if err = json.Unmarshal(data, &items); err != nil {
		return err
	}
	s.m = make(map[T]struct{})
	s.Add(items...)
	return nil
}
