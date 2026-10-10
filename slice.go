package kgo

import (
	"reflect"
)

type ChsSort interface {
}

// SlicePagination 对一个切片进行内存分页
func SlicePagination[T any](data []T, pageSize int) (paged [][]T) {
	if data == nil {
		return nil
	}
	if pageSize <= 0 {
		return nil
	}
	var count = len(data)
	var pageCount = count / pageSize
	if count%pageSize != 0 {
		pageCount += 1
	}
	paged = make([][]T, 0, pageCount)
	for page := 1; page <= pageCount; page++ {
		segment := make([]T, 0, pageSize)
		start := (page - 1) * pageSize
		for offset := start; offset < start+pageSize && offset < count; offset++ {
			segment = append(segment, data[offset])
		}
		if len(segment) > 0 {
			paged = append(paged, segment)
		}
	}
	return paged
}

// Contains 一个切片是否包含指定元素
// 判断 data 是否包含 element
func Contains[T any](data []T, element T) bool {
	for _, datum := range data {
		if reflect.DeepEqual(datum, element) {
			return true
		}
	}
	return false
}

// ContainsAll 一个切片是否包含所有指定元素
// 判断 data 是否包含 elements 中的所有元素
func ContainsAll[T any](data []T, elements []T) bool {
	for _, element := range elements {
		if !Contains(data, element) {
			return false
		}
	}
	return true
}

// ContainsAny 一个切片是否包含任意指定元素
// 判断 data 是否包含 elements 中的任意一个元素
func ContainsAny[T any](data []T, elements []T) bool {
	for _, element := range elements {
		if Contains(data, element) {
			return true
		}
	}
	return false
}

// SameElements 两个切片是否拥有相同的元素，不考虑它们的顺序，只要元素相同即可（重复元素的出现次数也需一致）
// 判断 data1 和 data2 是否包含相同的元素
func SameElements[T any](data1, data2 []T) bool {
	if len(data1) != len(data2) {
		return false
	}
	// used 标记 data2 中已被匹配的元素，避免重复元素被重复匹配
	used := make([]bool, len(data2))
	for _, datum1 := range data1 {
		found := false
		for j, datum2 := range data2 {
			if !used[j] && reflect.DeepEqual(datum1, datum2) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Intersection 两个切片的交集
// 返回在 s1 和 s2 中都存在的元素集合
// 注意：T 约束为 comparable，因为需要用作 map 的键；不可比较类型（切片/map/函数等）会在编译期被拒绝，避免运行时 panic。
func Intersection[T comparable](s1, s2 []T) (results []T) {
	hash := make(map[T]bool, len(s1))
	for _, elem := range s1 {
		hash[elem] = true
	}
	for _, elem := range s2 {
		if hash[elem] {
			results = append(results, elem)
		}
	}
	return results
}

// Union 两个切片的合集，如果遇到重复元素，只保留1个
// 返回在 s1 和 s2 中存在的元素集合
// 注意：T 约束为 comparable（需用作 map 键），不可比较类型会在编译期被拒绝。
func Union[T comparable](s1, s2 []T) (results []T) {
	hash := make(map[T]bool, len(s1)+len(s2))
	for _, elem := range s1 {
		hash[elem] = true
	}
	for _, elem := range s2 {
		hash[elem] = true
	}
	for elem := range hash {
		results = append(results, elem)
	}
	return results
}

// ReplaceOrAppend 替换或者追加一个元素
// s: 切片
// e: 要替换或者追加的元素
// fieldName: 结构体字段名，用于判断是否已存在， <strong>注意：</strong> 指定的字段必须已导出，否则不会生效）
// 返回操作后的新切片
func ReplaceOrAppend[T any](s []T, e T, fieldName string) []T {
	// 获取 e 的 fieldName 字段值
	ev := reflect.ValueOf(e)
	if ev.Kind() == reflect.Ptr {
		ev = ev.Elem()
	}
	if ev.Kind() != reflect.Struct {
		return append(s, e) // 非结构体直接追加
	}

	field := ev.FieldByName(fieldName)
	if !field.IsValid() || !field.CanInterface() {
		// 字段不存在或未导出，直接返回
		return s
	}
	targetValue := field.Interface()

	// 遍历切片，检查是否已有相同 fieldName 的元素
	for i, elem := range s {
		v := reflect.ValueOf(elem)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct {
			continue
		}

		fv := v.FieldByName(fieldName)
		if !fv.IsValid() {
			continue
		}

		currentValue := fv.Interface()
		if reflect.DeepEqual(currentValue, targetValue) {
			// 复制切片，避免修改原数据
			newSlice := make([]T, len(s))
			copy(newSlice, s)
			newSlice[i] = e
			return newSlice
		}
	}

	// 没有找到匹配项，追加
	return append(s, e)
}

// ReplaceOrAppendFunc 替换或追加元素，使用自定义比较函数
func ReplaceOrAppendFunc[T any](s []T, e T, isEqual func(a, b T) bool) []T {
	for i, elem := range s {
		if isEqual(elem, e) {
			newSlice := make([]T, len(s))
			copy(newSlice, s)
			newSlice[i] = e
			return newSlice
		}
	}
	return append(s, e)
}

// Diff 两个切片的差集，以第一个参数s1为基准，即：返回在s1中存在 且 在s2中不存在的元素集合
// 注意：T 约束为 comparable（需用作 map 键），不可比较类型会在编译期被拒绝。
func Diff[T comparable](s1, s2 []T) (results []T) {
	hash := make(map[T]bool, len(s2))
	for _, elem := range s2 {
		hash[elem] = true
	}
	for _, elem := range s1 {
		if !hash[elem] {
			results = append(results, elem)
		}
	}
	return results
}

// SplitCounter 根据给定的页大小和总数，生成一个二维切片
func SplitCounter[T any](pageSize, count int) (items [][]T) {
	if pageSize <= 0 || count <= 0 {
		return
	}
	pageCount := count / pageSize
	remainder := count % pageSize
	if remainder != 0 {
		pageCount++
	}
	for page := 0; page < pageCount; page++ {
		size := pageSize
		// 最后一段且不能整除时，容量为余数
		if page == pageCount-1 && remainder != 0 {
			size = remainder
		}
		items = append(items, make([]T, 0, size))
	}
	return items
}

// MergeSlice 多个切片合并
func MergeSlice[T any](sliceList ...[]T) (results []T) {
	if len(sliceList) == 0 {
		return
	}
	var totalCount = 0
	for _, slice := range sliceList {
		totalCount += len(slice)
	}
	results = make([]T, 0, totalCount)
	for _, slice := range sliceList {
		results = append(results, slice...)
	}
	return results
}
