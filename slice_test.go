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
