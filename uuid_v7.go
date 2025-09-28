package kgo

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var lastTimestamp int64
var seq uint16

// UuidV7 并发安全的 UUID v7 生成器
func UuidV7() string {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now().UnixMilli()

	// 处理同一毫秒内的序列号
	if now == lastTimestamp {
		seq++
		if seq > 0x0FFF { // 12位序列号最大值
			// 等待下一毫秒
			for now <= lastTimestamp {
				time.Sleep(100 * time.Microsecond)
				now = time.Now().UnixMilli()
			}
			seq = 0
		}
	} else {
		seq = 0
	}
	lastTimestamp = now

	var uuid [2]uint64

	// 第一部分：时间戳（48位） + 版本（4位） + 序列号（12位）
	timestampHigh := uint64(now) << 16
	version := uint64(0x7000)

	uuid[0] = timestampHigh | version | uint64(seq)

	// 第二部分：变体（2位） + 随机数（62位）
	variant := uint64(0x8000) << 48
	randBytes := make([]byte, 8)
	_, _ = rand.Read(randBytes)
	randPart := uint64(binary.BigEndian.Uint64(randBytes)) & 0x3FFFFFFFFFFFFFFF

	uuid[1] = variant | randPart

	return formatStandardUUID(uuid)
}

// SimpleUuidV7 并发安全的简化 UUID
func SimpleUuidV7() string {
	uuidStr := UuidV7()
	// 移除连字符
	return fmt.Sprintf("%s%s%s%s%s",
		uuidStr[0:8],
		uuidStr[9:13],
		uuidStr[14:18],
		uuidStr[19:23],
		uuidStr[24:36])
}

// formatStandardUUID 将 128 位数据格式化为标准 UUID 字符串
func formatStandardUUID(uuid [2]uint64) string {
	// UUID 格式：xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uint32(uuid[0]>>32),        // 时间戳高32位
		uint16(uuid[0]>>16)&0xFFFF, // 时间戳低16位 + 版本4位
		uint16(uuid[0])&0xFFFF,     // 随机数12位 + 填充
		uint16(uuid[1]>>48)&0xFFFF, // 变体 + 随机数高2位
		uuid[1]&0xFFFFFFFFFFFF)     // 随机数低48位
}
