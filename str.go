package kgo

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// 英文特殊符号
var (
	specificSymbolEn = []rune{
		'`', '~', '@', '!', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=',
		'|', '\\', '[', ']', '{', '}', ':', ';', '"', '\'', '<', '>', ',', '?', '/', ' ',
	}

	// 中文特殊符号
	specificSymbolCn = []rune{
		'“', '”', '（', '）', '、', '《', '》', '，', '；', '：', '？', '！', '…', '―',
		'—', '·', '¨', '｜', '〃', '‘', '’', '～', '‖', '∶', '＂', '＇', '｀', '｜',
		'〔', '〕', '〈', '〉', '「', '」', '『', '』', '〖', '〗', '【', '】', '（', '）', '［', '］', '｛', '｝', ' ', ' ',
	}

	// 英文单词判断正则表达式
	englishWordsPattern = regexp.MustCompile("[a-zA-Z]+")
)

// Clean 清理字符串，去除其实的各种特殊符号，一律替换为下划线，这样变会变成安全的字符串，在绝大多数场景下不会有问题
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

// R2S rune切片转为string
func R2S(b []rune) (s string) {
	return (string)(b)
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
//
// 示例1：
//
//	MaskChineseNameEx("张一二", 1, 1) // "张*二" 左边保留1个字符、右边保留1个字符
//
// 示例2：
//
//	name := "张一二"
//	MaskChineseNameEx(name, 0, utf8.RuneCountInString(name)-1) // "*一二" 左边保留0个字符、右边保留所有字符少1个，达到对姓氏脱敏的目的
func MaskChineseNameEx(name string, left, right int) (masked string) {
	// 将中文字符串转换为rune数组，方便进行字符级别的操作
	runes := []rune(name)
	size := len(runes)
	leftRunes := runes[:left]
	rightRunes := runes[size-right:]
	middleRunes := make([]rune, size-len(leftRunes)-len(rightRunes))
	for i := 0; i < len(middleRunes); i++ {
		middleRunes[i] = '*'
	}

	// 返回脱敏后的名字
	return string(leftRunes) + string(middleRunes) + string(rightRunes)
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
//
// 示例 MaskAnyString("一二三四五六七八九十", 3, 1) 得到 一二三******十
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

// ReverseString 字符串反转
func ReverseString(s string) string {
	runeSlice := []rune(s)
	for i, j := 0, len(runeSlice)-1; i < j; i, j = i+1, j-1 {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}
	return string(runeSlice)
}

// EnglishWordsCount 英文单词数量，以非英文字符分割之后统计
//
// 注意：这里假设了英文字符是 ASCII 字符集
func EnglishWordsCount(s string) int {
	// 使用正则表达式匹配英文单词
	words := englishWordsPattern.FindAllString(s, -1)
	return len(words)
}

// IsBlank 判断字符串是否为空
func IsBlank(s string) bool {
	if len(strings.TrimSpace(s)) == 0 {
		return true
	}
	for _, c := range s {
		if !isBlankChar(c) {
			return false
		}
	}
	return true
}

// IsNotBlank 判断字符串是否不为空
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// isBlankChar 判断字符是否为空白字符
func isBlankChar(c rune) bool {
	return unicode.IsSpace(c) ||
		c == '\ufeff' || // 字节次序标记字符
		c == '\u202a' || // 右到左标记
		c == '\u0000' || // 空字符
		c == '\u3164' || // 零宽度非连接符
		c == '\u2800' || // 零宽度空间
		c == '\u180e' // 蒙古文格式化字符
}

// MaxLen 获取字符串列表中最大长度
func MaxLen(strList ...string) int {
	maxLen := 0
	for _, str := range strList {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}
	return maxLen
}
