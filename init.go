package kgo

type IntUintBig interface {
	~int | ~int64 | ~uint | ~uint64
}

type IntUintFloat interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type IntUintFloatPtr interface {
	*int | *int8 | *int16 | *int32 | *int64 | *uint | *uint8 | *uint16 | *uint32 | *uint64 | *float32 | *float64
}
