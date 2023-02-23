package kgo

import (
	"fmt"
	"reflect"
	"strings"
)

var specificSymbolEn = []rune{
	'`', '~', '@', '!', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+', '=',
	'|', '\\', '[', ']', '{', '}', ':', ';', '"', '\'', '<', '>', ',', '?', '/', ' ',
}

var specificSymbolCn = []rune{
	'“', '”', '（', '）', '、', '《', '》', '，', '；', '：', '？', '！', '…', '―',
	'—', '·', '¨', '｜', '〃', '‘', '’', '～', '‖', '∶', '＂', '＇', '｀', '｜',
	'〔', '〕', '〈', '〉', '「', '」', '『', '』', '〖', '〗', '【', '】', '（', '）', '［', '］', '｛', '｝', ' ', ' ',
}

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
