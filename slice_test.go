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

func TestSlice_Intersection_Union_Diff_Integer(t *testing.T) {
	type args[T any] struct {
		data1 []T
		data2 []T
	}
	type testCase[T any] struct {
		name             string
		args             args[T]
		intersectionWant []T
		unionWant        []T
		diffWant         []T
	}
	tests := []testCase[int]{
		{
			name:             "SliceIntersection_Union_Diff_Integer",
			args:             args[int]{data1: []int{1, 2, 3, 4}, data2: []int{1, 2, 8, 10}},
			intersectionWant: []int{1, 2},
			unionWant:        []int{1, 2, 3, 4, 8, 10},
			diffWant:         []int{3, 4},
		},
	}
	for _, tt := range tests {
		if got := Intersection[int](tt.args.data1, tt.args.data2); !SameElements(got, tt.intersectionWant) {
			t.Errorf("Intersection() = %v, want %v", got, tt.intersectionWant)
		}
		if got := Union[int](tt.args.data1, tt.args.data2); !SameElements(got, tt.unionWant) {
			t.Errorf("Union() = %v, want %v", got, tt.unionWant)
		}
		if got := Diff[int](tt.args.data1, tt.args.data2); !SameElements(got, tt.diffWant) {
			t.Errorf("Diff() = %v, want %v", got, tt.diffWant)
		}
	}
}

func TestSlice_Intersection_Union_Diff_String(t *testing.T) {
	type args[T any] struct {
		data1 []T
		data2 []T
	}
	type testCase[T any] struct {
		name             string
		args             args[T]
		intersectionWant []T
		unionWant        []T
		diffWant         []T
	}
	tests := []testCase[string]{
		{
			name:             "SliceIntersection_Union_Diff_String",
			args:             args[string]{data1: []string{"1", "2", "3", "4"}, data2: []string{"1", "2", "8", "10"}},
			intersectionWant: []string{"1", "2"},
			unionWant:        []string{"1", "2", "3", "4", "8", "10"},
			diffWant:         []string{"3", "4"},
		},
	}
	for _, tt := range tests {
		if got := Intersection[string](tt.args.data1, tt.args.data2); !SameElements(got, tt.intersectionWant) {
			t.Errorf("Intersection() = %v, want %v", got, tt.intersectionWant)
		}
		if got := Union[string](tt.args.data1, tt.args.data2); !SameElements(got, tt.unionWant) {
			t.Errorf("Union() = %v, want %v", got, tt.unionWant)
		}
		if got := Diff[string](tt.args.data1, tt.args.data2); !SameElements(got, tt.diffWant) {
			t.Errorf("Diff() = %v, want %v", got, tt.diffWant)
		}
	}
}
