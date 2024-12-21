package kgo

import (
	"errors"
	"reflect"
)

type Stack[T any] struct {
	elements []any
}

// NewStack 创建一个新的栈
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push 向栈中添加元素
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop 从栈中移除并返回栈顶元素
func (s *Stack[T]) Pop() (T, error) {
	if len(s.elements) == 0 {
		return reflect.Zero(reflect.TypeOf(new(T))).Interface().(T), errors.New("stack is empty")
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element.(T), nil
}

// Peek 返回栈顶元素但不移除它
func (s *Stack[T]) Peek() (T, error) {
	if len(s.elements) == 0 {
		return reflect.Zero(reflect.TypeOf(new(T))).Interface().(T), errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1].(T), nil
}

// Size 返回栈中元素的数量
func (s *Stack[T]) Size() int {
	return len(s.elements)
}
