package kgo

import "testing"

func TestNewStack(t *testing.T) {
	stack := NewStack[int]()
	if stack != nil {
		t.Log("stack is not nil")
	} else {
		t.Errorf("stack is nil")
	}
}

func TestStack_Push(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("stack size is not 3")
	} else {
		t.Log("stack size is 3")
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	element, err := stack.Pop()
	if err != nil {
		t.Errorf("pop error: %s", err)
	} else if element != 3 {
		t.Errorf("pop element is not 3")
	}
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	element, err := stack.Peek()
	if err != nil {
		t.Errorf("peek error: %s", err)
	} else if element != 3 {
		t.Errorf("peek element is not 3")
	}
}

func TestStack_Size(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("stack size is not 3")
	} else {
		t.Log("stack size is 3")
	}
}
