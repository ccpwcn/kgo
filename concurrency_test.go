package kgo

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 本文件覆盖并发正确性与本地缓存的行为边界，配合 `go test -race` 使用。

// waitUntil 在超时前轮询条件，避免 -race / CI 调度抖动下固定 Sleep 误判。
func waitUntil(timeout time.Duration, cond func() bool) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if cond() {
			return true
		}
		time.Sleep(5 * time.Millisecond)
	}
	return cond()
}

// --- SimpleLocalCache 行为边界 ---

// maxCount<=0 表示不限容量
func TestCache_UnlimitedWhenMaxCountNonPositive(t *testing.T) {
	for _, max := range []int64{0, -1} {
		lc := NewLocalCache(max)
		for i := 0; i < 100; i++ {
			if err := lc.Set(fmt.Sprintf("k%d", i), i); err != nil {
				t.Fatalf("maxCount=%d 时应不限容量，却报错：%v", max, err)
			}
		}
		if lc.Size() != 100 {
			t.Errorf("maxCount=%d 期望 Size=100，实际 %d", max, lc.Size())
		}
	}
}

// 覆盖已存在的键不应重复计数
func TestCache_OverwriteDoesNotInflateSize(t *testing.T) {
	lc := NewLocalCache(10)
	_ = lc.Set("k", 1)
	_ = lc.Set("k", 2)
	_ = lc.Set("k", 3)
	if lc.Size() != 1 {
		t.Errorf("覆盖同一键后 Size 期望 1，实际 %d", lc.Size())
	}
	if v, ok := lc.Get("k"); !ok || v.(int) != 3 {
		t.Errorf("覆盖后应取到最新值 3，实际 %v(ok=%v)", v, ok)
	}
}

// 删除不存在的键应安全，且不使计数变负
func TestCache_DeleteNonExistentIsSafe(t *testing.T) {
	lc := NewLocalCache(10)
	lc.Delete("not-exist")
	if lc.Size() != 0 {
		t.Errorf("删除不存在的键后 Size 期望 0，实际 %d", lc.Size())
	}
	_ = lc.Set("k", 1)
	lc.Delete("not-exist")
	if lc.Size() != 1 {
		t.Errorf("误删导致 Size 错误，期望 1，实际 %d", lc.Size())
	}
}

// 容量上限：新键达到上限后应被拒绝，覆盖旧键仍可成功
func TestCache_CapacityLimit(t *testing.T) {
	lc := NewLocalCache(2)
	if err := lc.Set("a", 1); err != nil {
		t.Fatalf("写入 a 失败：%v", err)
	}
	if err := lc.Set("b", 2); err != nil {
		t.Fatalf("写入 b 失败：%v", err)
	}
	if err := lc.Set("c", 3); err == nil {
		t.Errorf("已达上限，写入新键 c 应失败")
	}
	// 覆盖已有键不受上限影响
	if err := lc.Set("a", 100); err != nil {
		t.Errorf("覆盖已有键 a 不应受上限限制：%v", err)
	}
	if lc.Size() != 2 {
		t.Errorf("Size 期望 2，实际 %d", lc.Size())
	}
}

// TTL 覆盖：旧定时器不应误删覆盖后的新值
func TestCache_TtlOverwriteNotDeletedByStaleTimer(t *testing.T) {
	lc := NewLocalCache(10)
	_ = lc.SetWithTtl("k", "v1", 50*time.Millisecond)
	_ = lc.SetWithTtl("k", "v2", 200*time.Millisecond) // 覆盖，取消旧定时器
	// 超过旧 TTL 后键仍应存在
	time.Sleep(80 * time.Millisecond)
	if !lc.Exists("k") {
		t.Fatalf("旧定时器误删了覆盖后的值")
	}
	if v, _ := lc.Get("k"); v.(string) != "v2" {
		t.Errorf("期望取到 v2，实际 %v", v)
	}
	if !waitUntil(500*time.Millisecond, func() bool { return !lc.Exists("k") }) {
		t.Errorf("新 TTL 到期后应被删除")
	}
}

// 用普通 Set 覆盖带 TTL 的键后，应取消过期，变为永久有效
func TestCache_PlainSetCancelsTtl(t *testing.T) {
	lc := NewLocalCache(10)
	_ = lc.SetWithTtl("k", "v1", 60*time.Millisecond)
	_ = lc.Set("k", "v2") // 取消 TTL
	time.Sleep(120 * time.Millisecond)
	if !lc.Exists("k") {
		t.Errorf("普通 Set 覆盖后应变为永久有效，但键被删除了")
	}
}

// SetWithTime 传入过去时间应报错
func TestCache_SetWithTimePast(t *testing.T) {
	lc := NewLocalCache(10)
	if err := lc.SetWithTime("k", 1, time.Now().Add(-time.Second)); err == nil {
		t.Errorf("过去时间应返回 error")
	}
}

// --- 并发安全（-race） ---

func TestCache_ConcurrentMixedOps(t *testing.T) {
	lc := NewLocalCache(1000)
	var wg sync.WaitGroup
	goroutines, ops := 20, 300
	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for i := 0; i < ops; i++ {
				key := fmt.Sprintf("k%d", i%50)
				switch i % 5 {
				case 0:
					_ = lc.Set(key, i)
				case 1:
					_, _ = lc.Get(key)
				case 2:
					_ = lc.SetWithTtl(key, i, 20*time.Millisecond)
				case 3:
					lc.Delete(key)
				default:
					_ = lc.Exists(key)
					_ = lc.Size()
				}
			}
		}(g)
	}
	wg.Wait()
}

// 并发写入不同键时，成功数量应恰好等于容量上限（check-and-set 原子性）
func TestCache_ConcurrentCapacityNotExceeded(t *testing.T) {
	const maxCount = 10
	const goroutines = 50
	lc := NewLocalCache(maxCount)
	var wg sync.WaitGroup
	var success int64
	var mu sync.Mutex
	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			if err := lc.Set(fmt.Sprintf("key-%d", g), g); err == nil {
				mu.Lock()
				success++
				mu.Unlock()
			}
		}(g)
	}
	wg.Wait()
	if success != maxCount {
		t.Errorf("成功写入数期望 %d，实际 %d", maxCount, success)
	}
	if lc.Size() != maxCount {
		t.Errorf("Size 期望 %d，实际 %d", maxCount, lc.Size())
	}
}

func TestConcurrent_Snowflake_Unique(t *testing.T) {
	if err := InitSnowflake(1, 1); err != nil {
		t.Fatalf("初始化失败：%v", err)
	}
	var wg sync.WaitGroup
	var seen sync.Map
	var dup int64
	goroutines, per := 8, 2000
	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < per; i++ {
				id := SnowflakeId()
				if _, loaded := seen.LoadOrStore(id, struct{}{}); loaded {
					atomic.AddInt64(&dup, 1)
				}
			}
		}()
	}
	wg.Wait()
	if dup != 0 {
		t.Errorf("并发下出现 %d 个重复 Snowflake ID", dup)
	}
}

func TestConcurrent_Uuid_Unique(t *testing.T) {
	var wg sync.WaitGroup
	var seen sync.Map
	var dup int64
	goroutines, per := 8, 2000
	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < per; i++ {
				id := Uuid()
				if _, loaded := seen.LoadOrStore(id, struct{}{}); loaded {
					atomic.AddInt64(&dup, 1)
				}
			}
		}()
	}
	wg.Wait()
	if dup != 0 {
		t.Errorf("并发下出现 %d 个重复 UUID", dup)
	}
}

func TestConcurrent_UuidV7_Unique(t *testing.T) {
	var wg sync.WaitGroup
	var seen sync.Map
	var dup int64
	goroutines, per := 4, 500
	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < per; i++ {
				id := UuidV7()
				if _, loaded := seen.LoadOrStore(id, struct{}{}); loaded {
					atomic.AddInt64(&dup, 1)
				}
			}
		}()
	}
	wg.Wait()
	if dup != 0 {
		t.Errorf("并发下出现 %d 个重复 UUIDv7", dup)
	}
}
