package kgo

import (
	"reflect"
	"testing"
	"time"
)

func Test_LocalCacheGetAndSet(t *testing.T) {
	const (
		key   = "local_cache_set_and_get_test"
		value = 1
	)
	lc := NewLocalCache(10)
	if err := lc.Set(key, value); err != nil {
		t.Errorf("LocalCache set error %+v", err)
	} else if v, ok := lc.Get(key); !ok {
		t.Errorf("LocalCache get failure")
	} else if vv, ok := v.(int); !ok {
		t.Errorf("LocalCache get value type failure")
	} else if vv != value {
		t.Errorf("LocalCache get value value failure")
	}
}

func Test_LocalCacheGetAndSet_WhenArray(t *testing.T) {
	const (
		key = "local_cache_set_and_get_test"
	)
	var (
		value = [3]int{1, 2, 3}
	)
	lc := NewLocalCache(10)
	if err := lc.Set(key, value); err != nil {
		t.Errorf("LocalCache set error %+v", err)
	} else if v, ok := lc.Get(key); !ok {
		t.Errorf("LocalCache get failure")
	} else if vv, ok := v.([3]int); !ok {
		t.Errorf("LocalCache get array value type failure")
	} else if !reflect.DeepEqual(vv, value) {
		t.Errorf("LocalCache get array value value failure, expected: %+v, actual: %+v", value, vv)
	}
}

func Test_LocalCacheGetAndSet_WhenSlice(t *testing.T) {
	const (
		key = "local_cache_set_and_get_test"
	)
	var (
		value = []int{1, 2, 3}
	)
	lc := NewLocalCache(10)
	if err := lc.Set(key, value); err != nil {
		t.Errorf("LocalCache set error %+v", err)
	} else if v, ok := lc.Get(key); !ok {
		t.Errorf("LocalCache get failure")
	} else if vv, ok := v.([]int); !ok {
		t.Errorf("LocalCache get array value type failure")
	} else if !reflect.DeepEqual(vv, value) {
		t.Errorf("LocalCache get array value value failure, expected: %+v, actual: %+v", value, vv)
	}
}

func Test_LocalCacheGetAndSet_WhenMap(t *testing.T) {
	const (
		key = "local_cache_set_and_get_test"
	)
	var (
		value = map[string]int{"1": 1, "2": 2, "3": 3}
	)
	lc := NewLocalCache(10)
	if err := lc.Set(key, value); err != nil {
		t.Errorf("LocalCache set error %+v", err)
	} else if v, ok := lc.Get(key); !ok {
		t.Errorf("LocalCache get failure")
	} else if vv, ok := v.(map[string]int); !ok {
		t.Errorf("LocalCache get map value type failure")
	} else if !reflect.DeepEqual(vv, value) {
		t.Errorf("LocalCache get map value value failure, expected: %+v, actual: %+v", value, vv)
	}
}

func Test_LocalCacheGetAndSet_WhenStruct(t *testing.T) {
	const (
		key = "local_cache_set_and_get_test"
	)
	type User struct {
		Name string
		Age  int
	}
	var (
		value = User{Name: "zhang", Age: 10}
	)
	lc := NewLocalCache(10)
	if err := lc.Set(key, value); err != nil {
		t.Errorf("LocalCache set error %+v", err)
	} else if v, ok := lc.Get(key); !ok {
		t.Errorf("LocalCache get failure")
	} else if vv, ok := v.(User); !ok {
		t.Errorf("LocalCache get struct value type failure")
	} else if !reflect.DeepEqual(vv, value) {
		t.Errorf("LocalCache get struct value value failure, expected: %+v, actual: %+v", value, vv)
	}
}

func Test_LocalCacheDelete(t *testing.T) {
	const (
		key   = "local_cache_set_and_get_test"
		value = 1
	)
	lc := NewLocalCache(10)
	if err := lc.Set(key, value); err != nil {
		t.Errorf("LocalCache set error %+v", err)
	} else {
		lc.Delete(key)
		if _, ok := lc.Get(key); ok {
			t.Errorf("LocalCache delete error %+v", err)
		}
	}
}

func Test_LocalCacheExists(t *testing.T) {
	const (
		key   = "local_cache_set_and_get_test"
		value = 1
	)
	lc := NewLocalCache(10)
	if err := lc.Set(key, value); err != nil {
		t.Errorf("LocalCache set error %+v", err)
	} else if !lc.Exists(key) {
		t.Errorf("LocalCache exists failure")
	}
}

func Test_LocalCacheSetWithTtl(t *testing.T) {
	const (
		key   = "local_cache_set_and_get_test"
		value = 1
	)
	lc := NewLocalCache(10)
	if err := lc.SetWithTtl(key, value, 100*time.Millisecond); err != nil {
		t.Errorf("LocalCache SetWithTtl error %+v", err)
	} else {
		time.Sleep(150 * time.Millisecond)
		if lc.Exists(key) {
			t.Errorf("LocalCache SetWithTtl failure")
		}
	}
}

func Test_LocalCacheSetWithTime(t *testing.T) {
	const (
		key   = "local_cache_set_and_get_test"
		value = 1
	)
	lc := NewLocalCache(10)
	if err := lc.SetWithTime(key, value, time.Now().Add(100*time.Millisecond)); err != nil {
		t.Errorf("LocalCache SetWithTime error %+v", err)
	} else {
		time.Sleep(150 * time.Millisecond)
		if lc.Exists(key) {
			t.Errorf("LocalCache SetWithTime failure")
		}
	}
}

func Test_LocalCacheSize(t *testing.T) {
	const (
		key   = "local_cache_set_and_get_test"
		value = 1
	)
	lc := NewLocalCache(10)
	if size := lc.Size(); size != 0 {
		t.Errorf("LocalCache origin size is not zero")
	} else if err := lc.SetWithTtl(key, value, 100*time.Millisecond); err != nil {
		t.Errorf("LocalCache set error %+v", err)
	} else if lc.Size() != 1 {
		t.Errorf("LocalCache size failure")
	} else {
		time.Sleep(150 * time.Millisecond)
		if lc.Size() != 0 {
			t.Errorf("LocalCache size failure")
		}
	}
}
