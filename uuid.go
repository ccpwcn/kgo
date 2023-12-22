package kgo

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"sync"
)

// A UUID is a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC
// 4122.
type UUID [16]byte

const randPoolSize = 16 * 16

var (
	rander  = rand.Reader // random function
	poolMu  sync.Mutex
	poolPos = randPoolSize     // protected with poolMu
	pool    [randPoolSize]byte // protected with poolMu
)

func init() {
	newRandomFromPool()
}

func Uuid() string {
	uuid := newRandomFromPool()
	var buf [36]byte
	encodeHex(buf[:], uuid)
	return B2S(buf[:])
}

func SimpleUuid() string {
	uuid := newRandomFromPool()
	var buf [32]byte
	encodeHexShort(buf[:], uuid)
	return B2S(buf[:])
}

func newRandomFromPool() UUID {
	var uuid UUID
	poolMu.Lock()
	if poolPos == randPoolSize {
		if _, err := io.ReadFull(rander, pool[:]); err != nil {
			poolMu.Unlock()
			panic(err)
		}
		poolPos = 0
	}
	copy(uuid[:], pool[poolPos:(poolPos+16)])
	poolPos += 16
	poolMu.Unlock()

	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid
}

func encodeHex(dst []byte, uuid UUID) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}

func encodeHexShort(dst []byte, uuid UUID) {
	hex.Encode(dst, uuid[:4])
	hex.Encode(dst[8:12], uuid[4:6])
	hex.Encode(dst[12:16], uuid[6:8])
	hex.Encode(dst[16:20], uuid[5:7])
	hex.Encode(dst[12:], uuid[6:])
}
