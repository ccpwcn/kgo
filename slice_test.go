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

func TestSplitCounter(t *testing.T) {
	type args struct {
		count    int
		pageSize int
	}
	type testCase[T any] struct {
		name string
		args args
		want [][]T
	}
	tests := []testCase[string]{
		{
			name: "SplitCounter1",
			args: args{count: 13, pageSize: 5},
			want: [][]string{{"", "", "", "", ""}, {"", "", "", "", ""}, {"", ""}},
		},
		{
			name: "SplitCounter2",
			args: args{count: 3, pageSize: 5},
			want: [][]string{{"", "", ""}},
		},
		{
			name: "SplitCounter3",
			args: args{count: 10, pageSize: 5},
			want: [][]string{{"", "", "", "", ""}, {"", "", "", "", ""}},
		},
	}
	for _, test := range tests {
		pageCount := test.args.count/test.args.pageSize + 1
		t.Run(test.name, func(t *testing.T) {
			items := SplitCounter[string](test.args.pageSize, test.args.count)
			if len(items) != pageCount {
				t.Errorf("len(items) = %v, want 3", len(items))
			}
			for i, item := range items {
				if i+1 < pageCount && cap(item) != test.args.pageSize {
					t.Errorf("cap(item) = %v, want %v", cap(item), test.args.pageSize)
				}
				if i+1 == pageCount && cap(item) != test.args.count%test.args.pageSize {
					t.Errorf("cap(item) = %v, want %v", cap(item), test.args.count%test.args.pageSize)
				}
			}
		})
	}
}

func TestMergeSlice_String(t *testing.T) {
	type args[T any] struct {
		data1 []T
		data2 []T
	}
	type testCase[T any] struct {
		name      string
		args      args[string]
		mergeWant []string
	}
	tests := []testCase[string]{
		{
			name:      "TestMergeSlice_String",
			args:      args[string]{data1: []string{"1", "2", "3"}, data2: []string{"8", "10"}},
			mergeWant: []string{"1", "2", "3", "8", "10"},
		},
	}
	for _, tt := range tests {
		if got := MergeSlice[string](tt.args.data1, tt.args.data2); !SameElements(got, tt.mergeWant) {
			t.Errorf("MergeSlice() = %v, want %v", got, tt.mergeWant)
		}
	}
}

func TestMergeSlice_String_Multi(t *testing.T) {
	type args[T any] struct {
		data1 []T
		data2 []T
		data3 []T
	}
	type testCase[T any] struct {
		name      string
		args      args[string]
		mergeWant []string
	}
	tests := []testCase[string]{
		{
			name:      "TestMergeSlice_String",
			args:      args[string]{data1: []string{"1", "2", "3"}, data2: []string{"8", "10"}, data3: []string{"10"}},
			mergeWant: []string{"1", "2", "3", "8", "10", "10"},
		},
	}
	for _, tt := range tests {
		if got := MergeSlice[string](tt.args.data1, tt.args.data2, tt.args.data3); !SameElements(got, tt.mergeWant) {
			t.Errorf("MergeSlice() = %v, want %v", got, tt.mergeWant)
		}
	}
}

func TestMergeSlice_Int(t *testing.T) {
	type args[T any] struct {
		data1 []T
		data2 []T
	}
	type testCase[T any] struct {
		name      string
		args      args[int]
		mergeWant []int
	}
	tests := []testCase[string]{
		{
			name:      "TestMergeSlice_String",
			args:      args[int]{data1: []int{1, 2, 3, 4}, data2: []int{5, 6, 7}},
			mergeWant: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		if got := MergeSlice[int](tt.args.data1, tt.args.data2); !SameElements(got, tt.mergeWant) {
			t.Errorf("MergeSlice() = %v, want %v", got, tt.mergeWant)
		}
	}
}

func TestMergeSlice_Int_Multi(t *testing.T) {
	type args[T any] struct {
		data1 []T
		data2 []T
		data3 []T
	}
	type testCase[T any] struct {
		name      string
		args      args[int]
		mergeWant []int
	}
	tests := []testCase[string]{
		{
			name:      "TestMergeSlice_String",
			args:      args[int]{data1: []int{1, 2, 3, 4}, data2: []int{5, 6, 7}, data3: []int{17, 18}},
			mergeWant: []int{1, 2, 3, 4, 5, 6, 7, 17, 18},
		},
	}
	for _, tt := range tests {
		if got := MergeSlice[int](tt.args.data1, tt.args.data2, tt.args.data3); !SameElements(got, tt.mergeWant) {
			t.Errorf("MergeSlice() = %v, want %v", got, tt.mergeWant)
		}
	}
}

func TestMergeSlice_Float(t *testing.T) {
	type args[T any] struct {
		data1 []T
		data2 []T
	}
	type testCase[T any] struct {
		name      string
		args      args[float64]
		mergeWant []float64
	}
	tests := []testCase[float64]{
		{
			name:      "TestMergeSlice_String",
			args:      args[float64]{data1: []float64{1.1, 2.2, 3.3}, data2: []float64{8.8}},
			mergeWant: []float64{1.1, 2.2, 3.3, 8.8},
		},
	}
	for _, tt := range tests {
		if got := MergeSlice[float64](tt.args.data1, tt.args.data2); !SameElements(got, tt.mergeWant) {
			t.Errorf("MergeSlice() = %v, want %v", got, tt.mergeWant)
		}
	}
}

func TestMergeSlice_Float_Multi(t *testing.T) {
	type args[T any] struct {
		data1 []T
		data2 []T
		data3 []T
	}
	type testCase[T any] struct {
		name      string
		args      args[float64]
		mergeWant []float64
	}
	tests := []testCase[float64]{
		{
			name:      "TestMergeSlice_String",
			args:      args[float64]{data1: []float64{1.1, 2.2, 3.3}, data2: []float64{8.8}, data3: []float64{18.8, 19.9, 20.1}},
			mergeWant: []float64{1.1, 2.2, 3.3, 8.8, 18.8, 19.9, 20.1},
		},
	}
	for _, tt := range tests {
		if got := MergeSlice[float64](tt.args.data1, tt.args.data2, tt.args.data3); !SameElements(got, tt.mergeWant) {
			t.Errorf("MergeSlice() = %v, want %v", got, tt.mergeWant)
		}
	}
}

func TestMergeSlice_Struct(t *testing.T) {
	type student struct {
		name string
		age  int
	}
	type args[T any] struct {
		data1 []T
		data2 []T
	}
	type testCase[T any] struct {
		name      string
		args      args[student]
		mergeWant []student
	}
	tests := []testCase[string]{
		{
			name:      "TestMergeSlice_String",
			args:      args[student]{data1: []student{{name: "s1", age: 1}, {name: "s2", age: 2}}, data2: []student{{name: "s10", age: 10}}},
			mergeWant: []student{{name: "s1", age: 1}, {name: "s2", age: 2}, {name: "s10", age: 10}},
		},
	}
	for _, tt := range tests {
		if got := MergeSlice[student](tt.args.data1, tt.args.data2); !SameElements(got, tt.mergeWant) {
			t.Errorf("MergeSlice() = %v, want %v", got, tt.mergeWant)
		}
	}
}

func TestMergeSlice_Struct_Multi(t *testing.T) {
	type student struct {
		name string
		age  int
	}
	type args[T any] struct {
		data1 []T
		data2 []T
		data3 []T
	}
	type testCase[T any] struct {
		name      string
		args      args[student]
		mergeWant []student
	}
	tests := []testCase[string]{
		{
			name: "TestMergeSlice_String",
			args: args[student]{
				data1: []student{
					{name: "s1", age: 1},
					{name: "s2", age: 2},
				},
				data2: []student{
					{name: "s10", age: 10},
				},
				data3: []student{
					{name: "s100", age: 100},
					{name: "s101", age: 101},
				},
			},
			mergeWant: []student{
				{name: "s1", age: 1},
				{name: "s2", age: 2},
				{name: "s10", age: 10},
				{name: "s100", age: 100},
				{name: "s101", age: 101},
			},
		},
	}
	for _, tt := range tests {
		if got := MergeSlice[student](tt.args.data1, tt.args.data2, tt.args.data3); !SameElements(got, tt.mergeWant) {
			t.Errorf("MergeSlice() = %v, want %v", got, tt.mergeWant)
		}
	}
}

func TestReplaceOrAppend_IntField(t *testing.T) {
	type student struct {
		Id   int64
		Name string
	}
	type args[T any] struct {
		data []T
		obj  T
	}
	type testCase struct {
		name string
		args args[student]
		want []student
	}
	tests := []testCase{
		{
			name: "TestReplaceOrAppend_IntField",
			args: args[student]{
				data: []student{
					{Id: 1, Name: "s1"},
					{Id: 2, Name: "s2"},
				},
				obj: student{Id: 1, Name: "s1_1"},
			},
			want: []student{
				{Id: 1, Name: "s1_1"},
				{Id: 2, Name: "s2"},
			},
		},
		{
			name: "TestReplaceOrAppend_IntField2",
			args: args[student]{
				data: []student{
					{Id: 1, Name: "s1"},
					{Id: 2, Name: "s2"},
					{Id: 3, Name: "s3"},
				},
				obj: student{Id: 3, Name: "s3_3"},
			},
			want: []student{
				{Id: 1, Name: "s1"},
				{Id: 2, Name: "s2"},
				{Id: 3, Name: "s3_3"},
			},
		},
	}
	for _, tt := range tests {
		if got := ReplaceOrAppend[student](tt.args.data, tt.args.obj, "Id"); !SameElements(got, tt.want) {
			t.Errorf("ReplaceOrAppend() = %v, want %v", got, tt.want)
		}
	}
}

func TestReplaceOrAppend_StringField(t *testing.T) {
	type student struct {
		Id   int64
		Name string
	}
	type args[T any] struct {
		data []T
		obj  T
	}
	type testCase struct {
		name string
		args args[student]
		want []student
	}
	tests := []testCase{
		{
			name: "TestReplaceOrAppend_StringField",
			args: args[student]{
				data: []student{
					{Id: 1, Name: "s1"},
					{Id: 2, Name: "s2"},
				},
				obj: student{Id: 11, Name: "s1"},
			},
			want: []student{
				{Id: 11, Name: "s1"},
				{Id: 2, Name: "s2"},
			},
		},
		{
			name: "TestReplaceOrAppend_StringField2",
			args: args[student]{
				data: []student{
					{Id: 1, Name: "s1"},
					{Id: 2, Name: "s2"},
					{Id: 3, Name: "s3"},
				},
				obj: student{Id: 33, Name: "s3"},
			},
			want: []student{
				{Id: 1, Name: "s1"},
				{Id: 2, Name: "s2"},
				{Id: 33, Name: "s3"},
			},
		},
	}
	for _, tt := range tests {
		if got := ReplaceOrAppend[student](tt.args.data, tt.args.obj, "Name"); !SameElements(got, tt.want) {
			t.Errorf("ReplaceOrAppend() = %v, want %v", got, tt.want)
		}
	}
}

func TestReplaceOrAppendFunc_IntField(t *testing.T) {
	type student struct {
		Id   int64
		Name string
	}
	type args[T any] struct {
		data []T
		obj  T
	}
	type testCase struct {
		name string
		args args[student]
		want []student
	}
	tests := []testCase{
		{
			name: "TestReplaceOrAppend_IntField",
			args: args[student]{
				data: []student{
					{Id: 1, Name: "s1"},
					{Id: 2, Name: "s2"},
				},
				obj: student{Id: 1, Name: "s1_1"},
			},
			want: []student{
				{Id: 1, Name: "s1_1"},
				{Id: 2, Name: "s2"},
			},
		},
		{
			name: "TestReplaceOrAppend_IntField2",
			args: args[student]{
				data: []student{
					{Id: 1, Name: "s1"},
					{Id: 2, Name: "s2"},
					{Id: 3, Name: "s3"},
				},
				obj: student{Id: 3, Name: "s3_3"},
			},
			want: []student{
				{Id: 1, Name: "s1"},
				{Id: 2, Name: "s2"},
				{Id: 3, Name: "s3_3"},
			},
		},
	}
	for _, tt := range tests {
		if got := ReplaceOrAppendFunc[student](tt.args.data, tt.args.obj, func(a, b student) bool {
			return a.Id == b.Id
		}); !SameElements(got, tt.want) {
			t.Errorf("ReplaceOrAppend() = %v, want %v", got, tt.want)
		}
	}
}

func TestReplaceOrAppendFunc_StringField(t *testing.T) {
	type student struct {
		Id   int64
		Name string
	}
	type args[T any] struct {
		data []T
		obj  T
	}
	type testCase struct {
		name string
		args args[student]
		want []student
	}
	tests := []testCase{
		{
			name: "TestReplaceOrAppend_StringField",
			args: args[student]{
				data: []student{
					{Id: 1, Name: "s1"},
					{Id: 2, Name: "s2"},
				},
				obj: student{Id: 11, Name: "s1"},
			},
			want: []student{
				{Id: 11, Name: "s1"},
				{Id: 2, Name: "s2"},
			},
		},
		{
			name: "TestReplaceOrAppend_StringField2",
			args: args[student]{
				data: []student{
					{Id: 1, Name: "s1"},
					{Id: 2, Name: "s2"},
					{Id: 3, Name: "s3"},
				},
				obj: student{Id: 33, Name: "s3"},
			},
			want: []student{
				{Id: 1, Name: "s1"},
				{Id: 2, Name: "s2"},
				{Id: 33, Name: "s3"},
			},
		},
	}
	for _, tt := range tests {
		if got := ReplaceOrAppendFunc[student](tt.args.data, tt.args.obj, func(a, b student) bool {
			return a.Name == b.Name
		}); !SameElements(got, tt.want) {
			t.Errorf("ReplaceOrAppend() = %v, want %v", got, tt.want)
		}
	}
}
