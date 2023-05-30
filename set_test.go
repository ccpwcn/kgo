package kgo

import (
	"encoding/json"
	"testing"
)

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

func Test_SetToArray(t *testing.T) {
	s := NewSet[string]()
	s.Add("1")
	s.Add("2")
	s.Add("3")
	items := s.ToArray()
	contain := 0
	for _, item := range items {
		if item == "1" || item == "2" || item == "3" {
			contain++
		}
	}
	if contain != 3 {
		t.Error("toArray failed")
	}
}

func Test_MarshalAndUnMarshal(t *testing.T) {
	newSet := NewSet[string]()
	oldSet := NewSet[string]()
	oldSet.Add("1")
	oldSet.Add("2")
	oldSet.Add("3")
	if b, err := json.Marshal(oldSet); err != nil {
		t.Error("marshal set error", err)
	} else if err = json.Unmarshal(b, &newSet); err != nil {
		t.Error("unmarshal set error", err)
	} else if newSet.Len() != oldSet.Len() {
		t.Errorf("len not equal, old len %v, new len %v", oldSet.Len(), newSet.Len())
	} else if !newSet.Contains("1") || !newSet.Contains("2") || !newSet.Contains("3") {
		t.Error("element lost")
	}
}
