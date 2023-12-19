package kgo

import (
	"fmt"
	"reflect"
	"strings"
)

// JoinStructsField 将一个结构体切片中的指定字段的值拼接成字符串，请注意：第二个参数指定的字段，要存在且已导出，否则结果可能不是预期的样子
func JoinStructsField[T interface{}](elements []T, fieldName string) (s string) {
	// 只处理导出了的字段
	if fieldName[0] < 'A' || fieldName[0] > 'Z' {
		return ""
	}
	if elements == nil || len(elements) == 0 {
		return ""
	} else {
		for _, element := range elements {
			v := reflect.ValueOf(element)
			rv := v.FieldByName(fieldName)
			if rv.Kind() != reflect.Invalid { // 判断字段是否存在
				fv := rv.Interface()
				switch fv.(type) {
				case uint, uint64, uint32, uint16, uint8, int, int64, int32, int16, int8, *uint, *uint64, *uint32, *uint16, *uint8, *int, *int64, *int32, *int16, *int8:
					s += fmt.Sprintf("%v", fv) + ","
				case float32, float64, *float32, *float64:
					s += fmt.Sprintf("%f", fv) + ","
				case string, *string:
					s += fmt.Sprintf("%s", fv) + ","
				case bool, *bool:
					s += fmt.Sprintf("%v", fv) + ","
				default:
					panic("嵌套类型暂不支持")
				}
			}
		}
		return strings.TrimRight(s, ",")
	}
}

// PickStructsField 将指定结构体切片中的指定字段的值提取出来，形成一个数组，请注意：第二个参数指定的字段，要存在且已导出，否则结果可能不是预期的样子
func PickStructsField[T interface{}, FT comparable](elements []T, fieldName string) (s []FT) {
	// 只处理导出了的字段
	if fieldName[0] < 'A' || fieldName[0] > 'Z' {
		return []FT{}
	}
	if elements == nil || len(elements) == 0 {
		return []FT{}
	} else {
		s = make([]FT, len(elements))
		for i, element := range elements {
			v := reflect.ValueOf(element)
			rv := v.FieldByName(fieldName)
			if rv.Kind() != reflect.Invalid { // 判断字段是否存在
				if fv, ok := rv.Interface().(FT); ok {
					s[i] = fv
				}
			}
		}
		return s
	}
}

// SliceGroupBy 指定一个结构体切片，按指定的字段对其进行分组
// 例如：
//
//	[{UserName: "zhang", Type: 1}, {UserName: "wang", Type: 1}, {UserName: "li", Type: 2}]
//
// 按Type这个字段对切片进行分组的结果是一个Map：
//
//	{
//	   1: [{UserName: "zhang", Type: 1}, {UserName: "wang", Type: 1}],
//	   2: [{UserName: "li", Type: 2}],
//	}
//
// T是结构体类型，KT是字段类型，字段类型必须是可比较的，即：comparable（因为它是Map的Key）
func SliceGroupBy[T any, KT comparable](data []T, fieldName string) map[KT][]T {
	// 只处理导出了的字段，如果给定的字段没有导出，那么什么也不做
	if fieldName[0] < 'A' || fieldName[0] > 'Z' {
		return map[KT][]T{}
	}
	if data == nil || len(data) == 0 {
		return map[KT][]T{}
	}
	// 遍历处理分组
	m := make(map[KT][]T)
	for _, datum := range data {
		v := reflect.ValueOf(datum)
		rv := v.FieldByName(fieldName)
		if rv.Kind() != reflect.Invalid { // 判断字段是否存在
			if fv, ok := rv.Interface().(KT); ok {
				m[fv] = append(m[fv], datum)
			}
		}
	}
	return m
}

// CopyFields 将一个结构体中的所有字段值全部复制到目标结构体中，可以指定忽略某些字段
// 注意：如果你的字段是非指针类型，那么是值拷贝，如果你的字段是指针类型，那么是引用拷贝
// 即：在字段是指针类型的情况下，即使复制值到目标结构体实例中了，修改源结构体实例的字段值，会影响目标结构体实例的值，因为是引用拷贝
func CopyFields[SRC any, DST any](src SRC, ignore ...string) (dst DST) {
	ignoreCount := len(ignore)
	srcValue := reflect.ValueOf(src)
	// 反射取值
	dstValue := reflect.ValueOf(&dst)
	// 去包装取实际类型
	dstValue = dstValue.Elem()
	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Type().Field(i)
		// 源中字段未导出，忽略
		if !srcField.IsExported() {
			continue
		}
		fieldName := srcValue.Type().Field(i).Name
		if ignoreCount > 0 && Contains(ignore, fieldName) {
			continue
		}
		srcFieldValue := srcValue.FieldByName(fieldName)
		dstFieldValue := dstValue.FieldByName(fieldName)
		df, ok := dstValue.Type().FieldByName(fieldName)
		// 目标中字段不存在，忽略
		if !ok {
			continue
		}
		// 目标中字段未导出，忽略
		if !df.IsExported() {
			continue
		}
		// 类型不匹配，不进行复制
		if srcFieldValue.Kind() != dstFieldValue.Kind() {
			continue
		}
		// 复制值
		if dstFieldValue.CanSet() {
			dstFieldValue.Set(srcFieldValue)
		}
	}
	return
}

// SetField 给任意结构体的任意字段设置一个新值
// newValue不可以传入反身过的值类型（即：不要传reflect.ValueOf(...)的结果），这会导致panic
func SetField(entity any, field string, newValue any) {
	val := reflect.ValueOf(entity)
	for val.Type().Kind() == reflect.Pointer {
		val = val.Elem()
	}
	fv := val.FieldByName(field)
	if !fv.CanSet() {
		return
	}
	fv.Set(reflect.ValueOf(newValue))
}
