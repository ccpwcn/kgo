package kgo

import (
	"errors"
	"sync"
	"time"
)

type LocalCacheable interface {
	// Get 获得一个缓存
	Get(key string) (value interface{}, ok bool)
	// Set 设置一个缓存
	Set(key string, value interface{}) (err error)
	// SetWithTtl 设置一个缓存并指定有效时间
	SetWithTtl(key string, value interface{}, duration time.Duration) (err error)
	// SetWithTime 设置一个缓存并指定在未来一个时间之前有效
	SetWithTime(key string, value interface{}, future time.Time) (err error)
	// Delete 删除一个缓存
	Delete(key string)
	// Exists 判断缓存是否存在
	Exists(key string) (exists bool)
	// Size 缓存大小
	Size() int64
}

// ErrCacheFull 缓存条目数量达到上限时返回
var ErrCacheFull = errors.New("缓存已经满了")

// SimpleLocalCache 一个简单的本地内存缓存。
//
// 并发安全：所有操作通过读写锁保护，容量以实际存储的键数量为准，不会出现计数漂移。
//
// maxCount 语义：maxCount > 0 时表示最多可存放的键数量，达到上限后对新键的写入会返回 ErrCacheFull；
// maxCount <= 0 时表示不限制容量。对已存在的键进行覆盖写入不受容量限制。
type SimpleLocalCache struct {
	mu          sync.RWMutex
	container   map[string]interface{}
	timers      map[string]*time.Timer // 每个键对应的过期定时器
	generations map[string]uint64      // 每个键的写入代号，用于让被覆盖的旧定时器失效
	maxCount    int64
	genSeq      uint64 // 单调递增的代号发生器，仅在持有写锁时访问
}

// NewLocalCache 初始化一个缓存实例，一般将其初始化为全局单例对象。
// maxCount > 0 表示容量上限，maxCount <= 0 表示不限容量。
func NewLocalCache(maxCount int64) *SimpleLocalCache {
	return &SimpleLocalCache{
		container:   make(map[string]interface{}),
		timers:      make(map[string]*time.Timer),
		generations: make(map[string]uint64),
		maxCount:    maxCount,
	}
}

// Get 获得一个缓存
func (receiver *SimpleLocalCache) Get(key string) (value interface{}, ok bool) {
	receiver.mu.RLock()
	defer receiver.mu.RUnlock()
	value, ok = receiver.container[key]
	return value, ok
}

// Set 设置一个缓存（不带过期时间）。若该键此前带 TTL，覆盖后其过期定时器会被取消，变为永久有效。
func (receiver *SimpleLocalCache) Set(key string, value interface{}) (err error) {
	receiver.mu.Lock()
	defer receiver.mu.Unlock()
	return receiver.set(key, value, false, 0)
}

// SetWithTtl 设置一个缓存并指定有效时间
func (receiver *SimpleLocalCache) SetWithTtl(key string, value interface{}, duration time.Duration) (err error) {
	receiver.mu.Lock()
	defer receiver.mu.Unlock()
	return receiver.set(key, value, true, duration)
}

// set 实际的写入逻辑，调用前必须已持有写锁。
func (receiver *SimpleLocalCache) set(key string, value interface{}, withTtl bool, ttl time.Duration) (err error) {
	_, exists := receiver.container[key]
	// 仅当写入一个新键、且设置了容量上限、且已达上限时，才拒绝
	if !exists && receiver.maxCount > 0 && int64(len(receiver.container)) >= receiver.maxCount {
		return ErrCacheFull
	}

	// 取消该键已有的过期定时器与代号，避免旧定时器误删覆盖后的新值
	receiver.cancelTimerLocked(key)

	receiver.container[key] = value

	if withTtl {
		receiver.genSeq++
		gen := receiver.genSeq
		receiver.generations[key] = gen
		receiver.timers[key] = time.AfterFunc(ttl, func() {
			receiver.expire(key, gen)
		})
	}
	return nil
}

// expire 定时器回调：仅当代号匹配（说明期间没有被覆盖或删除）时才真正移除该键。
func (receiver *SimpleLocalCache) expire(key string, gen uint64) {
	receiver.mu.Lock()
	defer receiver.mu.Unlock()
	if cur, ok := receiver.generations[key]; !ok || cur != gen {
		return
	}
	delete(receiver.container, key)
	delete(receiver.timers, key)
	delete(receiver.generations, key)
}

// cancelTimerLocked 取消并清理某个键的定时器与代号，调用前必须已持有写锁。
func (receiver *SimpleLocalCache) cancelTimerLocked(key string) {
	if timer, ok := receiver.timers[key]; ok {
		timer.Stop()
		delete(receiver.timers, key)
	}
	delete(receiver.generations, key)
}

// SetWithTime 设置一个缓存并指定在未来一个时间之前有效
func (receiver *SimpleLocalCache) SetWithTime(key string, value interface{}, future time.Time) (err error) {
	now := time.Now()
	if future.After(now) {
		return receiver.SetWithTtl(key, value, future.Sub(now))
	}
	return errors.New("指定的时间不能是一个过去的时间")
}

// Delete 删除一个缓存。删除不存在的键是安全的，不会影响容量计数。
func (receiver *SimpleLocalCache) Delete(key string) {
	receiver.mu.Lock()
	defer receiver.mu.Unlock()
	receiver.cancelTimerLocked(key)
	delete(receiver.container, key)
}

// Exists 判断缓存是否存在
func (receiver *SimpleLocalCache) Exists(key string) (exists bool) {
	receiver.mu.RLock()
	defer receiver.mu.RUnlock()
	_, exists = receiver.container[key]
	return exists
}

// Size 缓存中当前实际存在的键数量
func (receiver *SimpleLocalCache) Size() int64 {
	receiver.mu.RLock()
	defer receiver.mu.RUnlock()
	return int64(len(receiver.container))
}
