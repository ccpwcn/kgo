package kgo

import (
	"crypto/rand"
	"fmt"
	"math"
)

const (
	defaultNanoIdSize = 21
	defaultAlphabet   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
)

// NanoId 生成一个默认大小的NanoID
func NanoId() (string, error) {
	return newNanoIdWithSize(defaultNanoIdSize)
}

// MustNanoId 生成一个默认大小的NanoID，如果中途出现错误，抛出 panic
func MustNanoId() string {
	id, err := newNanoIdWithSize(defaultNanoIdSize)
	if err != nil {
		panic(err)
	}
	return id
}

func newNanoIdWithSize(size int) (string, error) {
	return newNanoIdWithAlphabet(size, defaultAlphabet)
}

func newNanoIdWithAlphabet(size int, alphabet string) (string, error) {
	if alphabet == "" || len(alphabet) >= 256 {
		return "", fmt.Errorf("alphabet must contain between 1 and 255 symbols")
	}

	if size <= 0 {
		return "", fmt.Errorf("size must be greater than zero")
	}

	mask := (2 << (int)(math.Log2(float64(len(alphabet)-1)))) - 1
	step := int(math.Ceil(1.6 * float64(mask) * float64(size) / float64(len(alphabet))))

	idBuilder := make([]byte, 0, size)
	bytes := make([]byte, step)

	for {
		_, err := rand.Read(bytes)
		if err != nil {
			return "", err
		}

		for _, b := range bytes {
			alphabetIndex := int(b) & mask

			if alphabetIndex < len(alphabet) {
				idBuilder = append(idBuilder, alphabet[alphabetIndex])
				if len(idBuilder) == size {
					return B2S(idBuilder), nil
				}
			}
		}
	}
}
