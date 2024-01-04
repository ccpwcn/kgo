package kgo

import "reflect"

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
func Contains[T any](data []T, element T) bool {
	for _, datum := range data {
		if reflect.DeepEqual(datum, element) {
			return true
		}
	}
	return false
}

// ContainsAll 一个切片是否包含所有指定元素
func ContainsAll[T any](data []T, elements []T) bool {
	for _, element := range elements {
		if !Contains(data, element) {
			return false
		}
	}
	return true
}

// ContainsAny 一个切片是否包含任意指定元素
func ContainsAny[T any](data []T, elements []T) bool {
	for _, element := range elements {
		if Contains(data, element) {
			return true
		}
	}
	return false
}

// SameElements 两个切片是否拥有相同的元素，不考虑它们的顺序，只要元素相同即可
func SameElements[T any](data1, data2 []T) bool {
	if len(data1) != len(data2) {
		return false
	}
	for _, datum1 := range data1 {
		found := false
		for _, datum2 := range data2 {
			if reflect.DeepEqual(datum1, datum2) {
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
func Intersection[T any](s1, s2 []T) (results []T) {
	hash := make(map[any]bool)
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
func Union[T any](s1, s2 []T) (results []T) {
	hash := make(map[any]bool)
	for _, elem := range s1 {
		hash[elem] = true
	}
	for _, elem := range s2 {
		hash[elem] = true
	}
	for elem := range hash {
		results = append(results, elem.(T))
	}
	return results
}

// Diff 两个切片的差集，以第一个参数s1为基准，即：返回在s1中存在 且 在s2中不存在的元素集合
func Diff[T any](s1, s2 []T) (results []T) {
	hash := make(map[any]bool)
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
