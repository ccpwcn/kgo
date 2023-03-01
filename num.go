package kgo

import (
	"fmt"
	"reflect"
)

func NumJoinStr[T IntUintFloat | IntUintFloatPtr](nums []T) string {
	var s string
	for i, n := range nums {
		pv := reflect.ValueOf(n)
		if pv.Kind() == reflect.Ptr {
			if i < len(nums)-1 {
				s += fmt.Sprintf("%v,", pv.Elem())
			} else {
				s += fmt.Sprintf("%v", pv.Elem())
			}
		} else {
			if i < len(nums)-1 {
				s += fmt.Sprintf("%v,", n)
			} else {
				s += fmt.Sprintf("%v", n)
			}
		}
	}
	return `[` + s + `]`
}

// Nums2Strings 将数字数组转换成字符串数组，在某些情况下，比如数字太大了，就可以自动转换，避免JS中精度丢失
func Nums2Strings[T IntUintFloat | IntUintFloatPtr](nums []T) []string {
	var ss = make([]string, len(nums), len(nums))
	for i, n := range nums {
		pv := reflect.ValueOf(n)
		if pv.Kind() == reflect.Ptr {
			if i < len(nums)-1 {
				ss[i] = fmt.Sprintf("%v", pv.Elem())
			} else {
				ss[i] += fmt.Sprintf("%v", pv.Elem())
			}
		} else {
			if i < len(nums)-1 {
				ss[i] = fmt.Sprintf("%v,", n)
			} else {
				ss[i] = fmt.Sprintf("%v", n)
			}
		}
	}
	return []string{}
}
