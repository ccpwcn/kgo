package kgo

import (
	"sync"
	"testing"
)

func TestMustNanoID(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(MustNanoId())
	}
}

func TestNormalNanoID(t *testing.T) {
	for i := 0; i < 20; i++ {
		if id, err := NanoId(); err != nil {
			t.Error(err)
		} else {
			t.Log(id)
		}
	}
}

func TestMustNanoID_CountLetter(t *testing.T) {
	var m = make(map[byte]int)
	for i := 0; i < 1000_0000; i++ {
		id := MustNanoId()
		for j := 0; j < len(id); j++ {
			m[id[j]]++
		}
	}
	for k, v := range m {
		t.Logf("%q: %d\n", k, v)
	}
}

func TestNormalNanoID_CheckDuplicate(t *testing.T) {
	const count = 100_0000
	ids := make(map[string]bool, count)
	for i := 0; i < count; i++ {
		if id, err := NanoId(); err != nil {
			t.Error(err)
		} else if ids[id] {
			t.Error("duplicate Id", id)
		}
	}
}

func TestNormalNanoID_CheckDuplicate5(t *testing.T) {
	const count = 500_0000
	ids := make(map[string]bool, count)
	for i := 0; i < count; i++ {
		if id, err := NanoId(); err != nil {
			t.Error(err)
		} else if ids[id] {
			t.Error("duplicate Id", id)
		}
	}
}

func TestNormalNanoIDCheck_Duplicate10(t *testing.T) {
	const count = 1000_0000
	ids := make(map[string]bool, count)
	for i := 0; i < count; i++ {
		if id, err := NanoId(); err != nil {
			t.Error(err)
		} else if ids[id] {
			t.Error("duplicate Id", id)
		}
	}
}

func TestNormalNanoIDCheck_Duplicate100(t *testing.T) {
	const count = 1_0000_0000
	ids := make(map[string]bool, count)
	for i := 0; i < count; i++ {
		if id, err := NanoId(); err != nil {
			t.Error(err)
		} else if ids[id] {
			t.Error("duplicate Id", id)
		}
	}
}

func Benchmark_NanoId(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buffer sync.Map
		for j := 0; j < 500_0000; j++ {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					id := MustNanoId()
					if _, ok := buffer.Load(id); ok {
						b.Errorf("duplicated Id %s", id)
					}
					buffer.Store(id, true)
				}
			})
		}
	}
}
