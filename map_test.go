package kgo

import "testing"

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
