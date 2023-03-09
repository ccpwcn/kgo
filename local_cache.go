package kgo

import (
	"errors"
	"sync"
	"sync/atomic"
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

type SimpleLocalCache struct {
	container *sync.Map
	maxCount  int64
	usedCount int64
}

// NewLocalCache 初始化一个缓存实例，一般将其初始化为全局单例对象
func NewLocalCache(maxCount int64) *SimpleLocalCache {
	var m sync.Map
	return &SimpleLocalCache{container: &m, maxCount: maxCount}
}

// Get 获得一个缓存
func (receiver *SimpleLocalCache) Get(key string) (value interface{}, ok bool) {
	return receiver.container.Load(key)
}

// Set 设置一个缓存
func (receiver *SimpleLocalCache) Set(key string, value interface{}) (err error) {
	if receiver.maxCount > 0 && receiver.usedCount < receiver.maxCount {
		receiver.container.Store(key, value)
		atomic.AddInt64(&receiver.usedCount, 1)
		return nil
	} else {
		return errors.New("缓存已经满了")
	}
}

// SetWithTtl 设置一个缓存并指定有效时间
func (receiver *SimpleLocalCache) SetWithTtl(key string, value interface{}, duration time.Duration) (err error) {
	if receiver.maxCount > 0 && receiver.usedCount < receiver.maxCount {
		receiver.container.Store(key, value)
		tm := time.After(duration)
		go func() {
			<-tm // tm会在时间到达之后自动释放，不必手动close
			receiver.container.Delete(key)
			atomic.AddInt64(&receiver.usedCount, -1)
		}()
		atomic.AddInt64(&receiver.usedCount, 1)
		return nil
	} else {
		return errors.New("缓存已经满了")
	}
}

// SetWithTime 设置一个缓存并指定在未来一个时间之前有效
func (receiver *SimpleLocalCache) SetWithTime(key string, value interface{}, future time.Time) (err error) {
	now := time.Now()
	if future.After(now) {
		return receiver.SetWithTtl(key, value, future.Sub(now))
	} else {
		return errors.New("指定的时间不能是一个过去的时间")
	}
}

func (receiver *SimpleLocalCache) Delete(key string) {
	receiver.container.Delete(key)
}

// Exists 判断缓存是否存在
func (receiver *SimpleLocalCache) Exists(key string) (exists bool) {
	if _, ok := receiver.container.Load(key); ok {
		return true
	} else {
		return false
	}
}
func (receiver *SimpleLocalCache) Size() int64 {
	return receiver.usedCount
}
