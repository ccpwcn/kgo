package kgo

import "testing"

func Test_SetInstance(t *testing.T) {
	s := NewSet[string]()
	s.Add("1")
	s.Add("2")
	s.Add("3")
	if s.Len() != 3 {
		t.Error("Expected 3, got", s.Len())
	}
}

func Test_SetRemoveItem(t *testing.T) {
	s := NewSet[string]()
	s.Add("1")
	s.Remove("1")
	if s.Len() != 0 {
		t.Error("Expected 3, got", s.Len())
	}
}

func Test_SetClear(t *testing.T) {
	s := NewSet[string]()
	s.Add("1")
	s.Clear()
	if s.Len() != 0 {
		t.Error("Expected 3, got", s.Len())
	}
}
func Test_SetContains(t *testing.T) {
	s := NewSet[string]()
	s.Add("123")
	if !s.Contains("123") {
		t.Error("Expected has 123, but not")
	}
}

func Test_SetLenAndEmpty(t *testing.T) {
	s := NewSet[string]()
	s.Add("123")
	s.Add("456")
	s.Add("789")
	if s.Len() != 3 {
		t.Error("Expected has 123, but got ", s.Len())
	}
	s.Clear()
	if !s.Empty() {
		t.Error("Expected empty, but not, it length ", s.Len())
	}
}
