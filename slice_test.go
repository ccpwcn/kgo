package kgo

import (
	"reflect"
	"testing"
)

func TestSlicePagination(t *testing.T) {
	type args[T any] struct {
		data     []T
		pageSize int
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		wantPaged [][]T
	}
	tests := []testCase[int]{
		{
			name: "int pagination 5",
			args: args[int]{
				data:     []int{1, 2, 3, 4, 5},
				pageSize: 2,
			},
			wantPaged: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name: "int pagination 10",
			args: args[int]{
				data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				pageSize: 5,
			},
			wantPaged: [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}},
		},
		{
			name: "int pagination 29",
			args: args[int]{
				data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
				pageSize: 10,
			},
			wantPaged: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, {21, 22, 23, 24, 25, 26, 27, 28, 29}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPaged := SlicePagination(tt.args.data, tt.args.pageSize); !reflect.DeepEqual(gotPaged, tt.wantPaged) {
				t.Errorf("SlicePagination() = %v, want %v", gotPaged, tt.wantPaged)
			}
		})
	}
}

func TestContainsString(t *testing.T) {
	type args[T any] struct {
		data    []T
		element T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[string]{
		{name: "stringsContains", args: args[string]{data: []string{"abc", "def", "xyz"}, element: "abc"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.data, tt.args.element); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInteger(t *testing.T) {
	type args[T any] struct {
		data    []T
		element T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{name: "integerContains", args: args[int]{data: []int{1, 2, 3}, element: 1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.data, tt.args.element); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsFloat(t *testing.T) {
	type args[T any] struct {
		data    []T
		element T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[float64]{
		{name: "floatContains", args: args[float64]{data: []float64{1.1, 2.2, 3.3}, element: 1.1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.data, tt.args.element); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAllStrings(t *testing.T) {
	type args[T any] struct {
		data     []T
		elements []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[string]{
		{name: "stringsContainsAll", args: args[string]{data: []string{"abc", "def", "xyz"}, elements: []string{"abc", "def"}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAll(tt.args.data, tt.args.elements); got != tt.want {
				t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAllIntegers(t *testing.T) {
	type args[T any] struct {
		data     []T
		elements []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{name: "integerContainsAll", args: args[int]{data: []int{1, 2, 3}, elements: []int{1, 2}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAll(tt.args.data, tt.args.elements); got != tt.want {
				t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAllFloats(t *testing.T) {
	type args[T any] struct {
		data     []T
		elements []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[float64]{
		{name: "floatContainsAll", args: args[float64]{data: []float64{1.1, 2.2, 3.3}, elements: []float64{1.1, 2.2}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAll(tt.args.data, tt.args.elements); got != tt.want {
				t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAnyString(t *testing.T) {
	type args[T any] struct {
		data     []T
		elements []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[string]{
		{name: "stringsContainsAny", args: args[string]{data: []string{"abc", "def", "xyz"}, elements: []string{"abc", "opq"}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAny(tt.args.data, tt.args.elements); got != tt.want {
				t.Errorf("ContainsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAnyInteger(t *testing.T) {
	type args[T any] struct {
		data     []T
		elements []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{name: "integerContainsAny", args: args[int]{data: []int{1, 2, 3, 4}, elements: []int{10, 1, 2, 300}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAny(tt.args.data, tt.args.elements); got != tt.want {
				t.Errorf("ContainsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
