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

func Test_SimpleUuid_BillionIdMustNotDuplicated(t *testing.T) {
	const (
		total     = 1_000_000_000
		pageSize  = 1_000_000
		pageCount = total/pageSize + 1
	)
	var (
		buffer sync.Map
		wg     sync.WaitGroup
		task   = func(page int) {
			defer wg.Done()
			for i := 0; i < pageSize; i++ {
				id := SimpleUuid()
				if _, ok := buffer.Load(id); ok {
					t.Errorf("duplicated UUID %s", id)
				}
				buffer.Store(id, true)
				if page%10_000_000 == 0 {
					t.Logf("page %d has been validated OK!", page)
				}
			}
		}
	)
	for page := 1; page <= pageCount; page++ {
		wg.Add(1)
		go task(page)
	}
	wg.Wait()
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
	for i := 1; i < 2_000_000; i++ {
		id := Uuid()
		if m[id] {
			t.Errorf("duplicated UUID %s", id)
		}
		m[id] = true
	}
}

func Test_SimpleUuid_Million(t *testing.T) {
	m := make(map[string]bool)
	for i := 1; i < 2_000_000; i++ {
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
		for i := 0; i < 5_000_000; i++ {
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
		for i := 0; i < 5_000_000; i++ {
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
