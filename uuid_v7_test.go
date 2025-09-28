package kgo

import (
	"regexp"
	"testing"
)

func TestUuidFormat(t *testing.T) {
	// 测试标准格式
	uuid := UuidV7()
	matched, _ := regexp.MatchString(`^[0-9a-f]{8}-[0-9a-f]{4}-7[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`, uuid)
	if !matched {
		t.Errorf("Invalid UUID v7 format: %s", uuid)
	}
	t.Logf("UUID: %s", uuid)

	// 测试简化格式
	simple := SimpleUuidV7()
	if len(simple) != 32 {
		t.Errorf("Simple UUID should be 32 chars, got %d: %s", len(simple), simple)
	}
	t.Logf("Simple UUID: %s", simple)
}

func TestUuidUniqueness(t *testing.T) {
	// 测试生成大量 UUID 确保唯一性
	uuids := make(map[string]bool)
	for i := 0; i < 10000; i++ {
		uuid := UuidV7()
		if uuids[uuid] {
			t.Errorf("Duplicate UUID generated: %s", uuid)
		}
		uuids[uuid] = true
	}
}

func TestSimpleUuidFormat(t *testing.T) {
	// 测试简化格式
	uuid := SimpleUuidV7()
	if len(uuid) != 32 {
		t.Errorf("Simple UUID should be 32 chars, got %d: %s", len(uuid), uuid)
	}
	t.Logf("Simple UUID: %s", uuid)
}

func TestSimpleUuidUniqueness(t *testing.T) {
	// 测试生成大量 无连接符的 UUID 确保唯一性
	uuids := make(map[string]bool)
	for i := 0; i < 10000; i++ {
		uuid := SimpleUuidV7()
		if uuids[uuid] {
			t.Errorf("Duplicate UUID generated: %s", uuid)
		}
		uuids[uuid] = true
	}
}
