package kgo

import (
	"sync"
	"testing"
)

func TestSnowflake_Id(t *testing.T) {
	if err := InitSnowflake(1, 1); err != nil {
		t.Errorf("雪花算法初始化失败了 %+v", err)
	} else {
		const count = 10
		var buffer = make(map[int64]int8, count*3)
		for i := 0; i < count; i++ {
			id := SnowflakeId()
			if _, ok := buffer[id]; ok {
				t.Errorf("ID %+v 重复了", id)
			} else {
				buffer[id] = 1
				t.Logf("%+v\n", id)
			}
		}
	}
}

func TestSnowflake_Concurrent_Id(t *testing.T) {
	if err := InitSnowflake(1, 1); err != nil {
		t.Errorf("雪花算法初始化失败了 %+v", err)
	} else {
		const count = 5
		const machine = 3
		var buffer sync.Map
		var wg sync.WaitGroup
		for i := 1; i <= machine; i++ {
			wg.Add(1)
			go func(i int64) {
				for j := 0; j < count; j++ {
					id := SnowflakeId()
					if _, ok := buffer.Load(id); ok {
						t.Errorf("ID %+v 重复了", id)
					} else {
						buffer.Store(id, 1)
						t.Logf("ID %+v\n", id)
					}
				}
				wg.Done()
			}(int64(i))
		}
		wg.Wait()
	}
}

func BenchmarkSnowflake_Concurrent_Id(b *testing.B) {
	if err := InitSnowflake(1, 1); err != nil {
		b.Errorf("雪花算法初始化失败了 %+v", err)
		return
	}

	b.ResetTimer()
	tests := []struct {
		name    string
		count   int
		wantIds []uint64
	}{
		{"3 ids", 3, []uint64{}},
		{"10 ids", 10, []uint64{}},
		{"100 ids", 100, []uint64{}},
		{"1000 ids", 1000, []uint64{}},
		{"10000 ids", 10000, []uint64{}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			var buffer sync.Map
			for i := 0; i < tt.count; i++ {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						id := SnowflakeId()
						if _, ok := buffer.Load(id); ok {
							b.Errorf("ID %+v 重复了", id)
						} else {
							buffer.Store(id, 1)
							// b.Logf("ID %+v\n", Id)
						}
					}
				})
			}
		}
	}
}
