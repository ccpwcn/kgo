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
