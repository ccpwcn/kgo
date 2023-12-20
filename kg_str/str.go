package kg_str

import (
	"fmt"
	"unicode/utf8"
)

const DefaultCharCount = 1

type MaskOption func(masker *Masker)

// Masker 字符串脱敏工具
type Masker struct {
	maskCount int
	left      int
	right     int
	maskChar  rune
}

// WithMaskCount 脱敏字符数量，默认是输入字符串总字符数量-左侧保留字符数-右侧保留字符数，即：默认脱敏之后的字符串字符数量和原字符串相同
func WithMaskCount(maskCount int) MaskOption {
	return func(masker *Masker) {
		masker.maskCount = maskCount
	}
}

// WithLeft 脱敏字符左侧保留字符数量，未指定时使用默认值 DefaultCharCount
func WithLeft(left int) MaskOption {
	return func(masker *Masker) {
		masker.left = left
	}
}

// WithRight 脱敏字符右侧保留字符数量，未指定时使用默认值 DefaultCharCount
func WithRight(right int) MaskOption {
	return func(masker *Masker) {
		masker.right = right
	}
}

// WithMaskChar 脱敏字符，如果未指定，默认使用星号*
func WithMaskChar(maskChar rune) MaskOption {
	return func(masker *Masker) {
		masker.maskChar = maskChar
	}
}

// NewMasker 创建脱敏工具实例
//
//	如果您既没有指定左侧保留字符数量，也没有指定右侧保留字符数量，那么默认左右各保留一个字符
//	如果您没有指定脱敏字符，那么默认使用星号*
func NewMasker(options ...MaskOption) *Masker {
	masker := &Masker{}
	for _, option := range options {
		option(masker)
	}
	// 判断并设置默认值
	if masker.maskChar == 0 {
		masker.maskChar = '*'
	}
	if masker.left == 0 && masker.right == 0 {
		masker.left = DefaultCharCount
		masker.right = DefaultCharCount
	}
	return masker
}

// Mask 脱敏任意字符串
// 入参： s 要脱敏的字符串
//
//	注意1：如果入参字符串太短，可能默认的左侧保留 DefaultCharCount、右侧保留 DefaultCharCount 无法满足您的需求，此时您在实例化时要指定相应选项
//	注意2：如果入参字符串为空，也返回空，不会报错
//	注意3：如果指定了脱敏字符串的数量，那么返回的字符串长度有可能和您提供的入参字符串的长度并不相同
func (mk *Masker) Mask(s string) string {
	var (
		left, middle, right string
	)

	size := utf8.RuneCountInString(s)
	if size == 0 {
		return ""
	}
	// 如果初始化的时候没有设置starCount，则默认用字符串长度减去左右保留的长度
	maskCount := mk.maskCount
	if maskCount == 0 {
		maskCount = size - mk.left - mk.right
	}

	i := 0
	for _, n := range s {
		if i < mk.left {
			left += fmt.Sprintf("%c", n)
		} else if i >= size-mk.right {
			right += fmt.Sprintf("%c", n)
		} else if utf8.RuneCountInString(middle) < maskCount {
			middle += fmt.Sprintf("%c", mk.maskChar)
		}
		i++
	}
	// 如果字符串太短，以至于无法脱敏，进行扩张
	middleSize := len(middle)
	if maskCount > middleSize {
		for i := 0; i < maskCount-middleSize; i++ {
			middle += fmt.Sprintf("%c", mk.maskChar)
		}
	}
	if middle == "" {
		middle = fmt.Sprintf("%c", mk.maskChar)
	}
	return left + middle + right
}
