package kgo

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
	"unsafe"
)

// 英文特殊符号
var specificSymbolEn = []rune{
	'`', '~', '@', '!', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+', '=',
	'|', '\\', '[', ']', '{', '}', ':', ';', '"', '\'', '<', '>', ',', '?', '/', ' ',
}

// 中文特殊符号
var specificSymbolCn = []rune{
	'“', '”', '（', '）', '、', '《', '》', '，', '；', '：', '？', '！', '…', '―',
	'—', '·', '¨', '｜', '〃', '‘', '’', '～', '‖', '∶', '＂', '＇', '｀', '｜',
	'〔', '〕', '〈', '〉', '「', '」', '『', '』', '〖', '〗', '【', '】', '（', '）', '［', '］', '｛', '｝', ' ', ' ',
}

// Clean 清理字符串，去除其实的各种特殊符号，一律替换为下划线
func Clean(str string) string {
	var onlyCharacters string
	onlyCharacters = strings.ReplaceAll(str, ".", "__")
	onlyCharacters = strings.ReplaceAll(onlyCharacters, "。", "__")
	for _, ch := range specificSymbolEn {
		onlyCharacters = strings.ReplaceAll(onlyCharacters, string(ch), "_")
	}
	for _, ch := range specificSymbolCn {
		onlyCharacters = strings.ReplaceAll(onlyCharacters, string(ch), "_")
	}
	onlyCharacters = strings.ReplaceAll(onlyCharacters, "-", "_")
	return onlyCharacters
}

// JoinElements 将多个元素拼接成一个字符串，入参请使用基本类型，不处理结构体
func JoinElements[T IntUintFloat | IntUintFloatPtr | string | *string | bool | *bool](elements []T) (s string) {
	if elements == nil || len(elements) == 0 {
		return ""
	} else {
		for _, element := range elements {
			pv := reflect.ValueOf(element)
			if pv.Kind() == reflect.Ptr {
				s += fmt.Sprintf("%v", pv.Elem()) + ","
			} else {
				s += fmt.Sprintf("%v", element) + ","
			}
		}
		return strings.TrimRight(s, ",")
	}
}

// B2S byte切片转为string
func B2S(b []byte) (s string) {
	return *(*string)(unsafe.Pointer(&b))
}

// S2B string转为byte切片
func S2B(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

// MaskChineseName 中文姓名脱敏
func MaskChineseName(name string) (masked string) {
	size := utf8.RuneCountInString(name)
	if size == 0 {
		return ""
	}
	for i, n := range name {
		if i == 0 {
			masked = fmt.Sprintf("%c", n)
		} else {
			masked += "*"
		}
	}
	return masked
}

// MaskChineseNameEx 中文姓名脱敏，可以指定左右保留字符数量
// 示例：MaskChineseNameEx("张一二", 1, 1) => "*一*" 左边1位替代、右边1位替代
func MaskChineseNameEx(name string, left, right int) (masked string) {
	size := utf8.RuneCountInString(name)
	if size == 0 {
		return ""
	}
	i := 0
	for _, n := range name {
		if i < left || i >= size-right {
			masked += fmt.Sprintf("%c", n)
		} else {
			masked += "*"
		}
		i++
	}
	return masked
}

// MaskChineseMobile 中国手机号脱敏
//
// 输入 13012345678 得到 130****5678 这是最常见的方式，为了简单省事，就不加参数了
func MaskChineseMobile(tel string) (masked string) {
	size := len(tel)
	if size == 0 {
		return ""
	}
	for i, n := range tel {
		if i < 3 || i >= 7 {
			masked += fmt.Sprintf("%c", n)
		} else {
			masked += "*"
		}
	}
	return masked
}

// MaskChineseIdCard34 中国身份证号脱敏，仅支持18位身份证号，但是这足够了
//
// left 左侧保留几位
// right 右边保留几位
//
// 示例 MaskChineseIdCard34("101101199010101234") 得到 101***********1234
//
// 请对输入的身份证号自行校验，否则可能触发意料之外的结果
func MaskChineseIdCard34(idCard string) (masked string) {
	return MaskChineseIdCard(idCard, 3, 4)
}

// MaskChineseIdCard64 中国身份证号脱敏，仅支持18位身份证号，但是这足够了
//
// left 左侧保留几位
// right 右边保留几位
//
// 示例 MaskChineseIdCard64("101101199010101234") 得到 101101********1234
//
// 请对输入的身份证号自行校验，否则可能触发意料之外的结果
func MaskChineseIdCard64(idCard string) (masked string) {
	return MaskChineseIdCard(idCard, 6, 4)
}

// MaskChineseIdCard11 中国身份证号脱敏，仅支持18位身份证号，但是这足够了
//
// left 左侧保留几位
// right 右边保留几位
//
// 示例 MaskChineseIdCard11("101101199010101234") 得到 1****************4
//
// 请对输入的身份证号自行校验，否则可能触发意料之外的结果
func MaskChineseIdCard11(idCard string) (masked string) {
	return MaskChineseIdCard(idCard, 1, 1)
}

// MaskChineseIdCard 中国身份证号脱敏，仅支持18位身份证号，但是这足够了
//
// left 左侧保留几位
// right 右边保留几位
//
// 示例 MaskChineseIdCard("101101199010101234", 3, 4) 得到 101***********1234
//
// 请对输入的身份证号自行校验，否则可能触发意料之外的结果
func MaskChineseIdCard(idCard string, left, right int) (masked string) {
	size := len(idCard)
	if size == 0 {
		return ""
	}
	for i, n := range idCard {
		if i < left || i >= 18-right {
			masked += fmt.Sprintf("%c", n)
		} else {
			masked += "*"
		}
	}
	return masked
}

// MaskAnyString 任意字符串脱敏
//
// left 左侧保留几位
// right 右边保留几位
//
// 示例 MaskAnyString("一二三四五", 1, 1) 得到 一***五
func MaskAnyString(s string, left, right int) (masked string) {
	size := utf8.RuneCountInString(s)
	if size == 0 {
		return ""
	}
	i := 0
	for _, n := range s {
		if i < left || i >= size-right {
			masked += fmt.Sprintf("%c", n)
		} else {
			masked += "*"
		}
		i++
	}
	return masked
}
