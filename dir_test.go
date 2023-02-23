package kgo

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestGetWorkDir(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Error("get current dir failed")
	} else {
		tests := []struct {
			name string
			want string
		}{
			{"work dir", filepath.Dir(filename)},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := GetWorkDir(); got != tt.want {
					t.Errorf("GetWorkDir() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestIsExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "exists", args: args{filename: "dir.go"}, want: true},
		{name: "exists", args: args{filename: "dir_no.go"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExists(tt.args.filename); got != tt.want {
				t.Errorf("IsExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustRead(t *testing.T) {
	const (
		content = "Must Read function testing"
	)
	if f, err := os.CreateTemp(os.TempDir(), "test_must_read"); err != nil {
		t.Errorf("create testing temp file failed, %+v", err)
	} else if _, err := f.WriteString(content); err != nil {
		t.Errorf("write string content into testing temp file failed, %+v", err)
	} else {
		type args struct {
			filename string
		}
		tests := []struct {
			name string
			args args
			want []byte
		}{
			{name: "testing read file", args: args{filename: f.Name()}, want: []byte(content)},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := MustRead(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("MustRead() = %v, want %v", got, tt.want)
				}
			})
		}
		if err := f.Close(); err != nil {
			t.Errorf("write string content into testing file %s done, but delete temp file failed, %+v", f.Name(), err)
		}
	}
}
