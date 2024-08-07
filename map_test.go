package kgo

import (
	"testing"
)

func TestHasKey(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
		k K
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want bool
	}
	tests := []testCase[string, int]{
		{
			name: "exists", args: args[string, int]{m: map[string]int{"k": 1}, k: "k"}, want: true,
		},
		{
			name: "notExists", args: args[string, int]{m: map[string]int{"k": 1}, k: "kk"}, want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasKey(tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("HasKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeys(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want []K
	}
	tests := []testCase[string, int]{
		{name: "string keys", args: args[string, int]{map[string]int{"a": 1, "b": 2, "c": 3}}, want: []string{"b", "a", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapKeys[string, int](tt.args.m); !ContainsAll(got, tt.want) {
				t.Errorf("MapKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysEmptyStruct(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want []K
	}
	tests := []testCase[int64, struct{}]{
		{
			name: "string keys",
			args: args[int64, struct{}]{
				map[int64]struct{}{
					1: {},
					2: {},
					3: {},
				},
			}, want: []int64{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapKeys[int64, struct{}](tt.args.m); !ContainsAll(got, tt.want) {
				t.Errorf("MapKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
