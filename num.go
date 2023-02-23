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
