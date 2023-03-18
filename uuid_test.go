package kgo

import (
	"sync"
	"testing"
)

func Test_Uuid(t *testing.T) {
	m := make(map[string]bool)
	for i := 1; i < 5; i++ {
		id := Uuid()
		if m[id] {
			t.Errorf("duplicated UUID %s", id)
		} else {
			t.Log(id)
		}
		m[id] = true
	}
}

func Test_SimpleUuid(t *testing.T) {
	m := make(map[string]bool)
	for i := 1; i < 5; i++ {
		id := SimpleUuid()
		if m[id] {
			t.Errorf("duplicated UUID %s", id)
		} else {
			t.Log(id)
		}
		m[id] = true
	}
}

func Test_Uuid_Million(t *testing.T) {
	m := make(map[string]bool)
	for i := 1; i < 100_0000; i++ {
		id := Uuid()
		if m[id] {
			t.Errorf("duplicated UUID %s", id)
		}
		m[id] = true
	}
}

func Test_SimpleUuid_Million(t *testing.T) {
	m := make(map[string]bool)
	for i := 1; i < 100_0000; i++ {
		id := SimpleUuid()
		if m[id] {
			t.Errorf("duplicated UUID %s", id)
		}
		m[id] = true
	}
}

func Test_Uuid_Billion(t *testing.T) {
	m := make(map[string]bool)
	for i := 1; i < 1_0000_0000; i++ {
		id := Uuid()
		if m[id] {
			t.Errorf("duplicated UUID %s", id)
		}
		m[id] = true
	}
}

func Test_SimpleUuid_Billion(t *testing.T) {
	m := make(map[string]bool)
	for i := 1; i < 1_0000_0000; i++ {
		id := SimpleUuid()
		if m[id] {
			t.Errorf("duplicated UUID %s", id)
		}
		m[id] = true
	}
}

func Benchmark_Uuid(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buffer sync.Map
		for i := 0; i < 500_0000; i++ {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					id := Uuid()
					if _, ok := buffer.Load(id); ok {
						b.Errorf("duplicated UUID %s", id)
					}
					buffer.Store(id, true)
				}
			})
		}
	}
}

func Benchmark_SimpleUuid(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buffer sync.Map
		for i := 0; i < 500_0000; i++ {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					id := SimpleUuid()
					if _, ok := buffer.Load(id); ok {
						b.Errorf("duplicated UUID %s", id)
					}
					buffer.Store(id, true)
				}
			})
		}
	}
}
