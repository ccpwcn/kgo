package kgo

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetExeDir 获得可执行程序所在目录的绝对完整路径
func GetExeDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

// GetWorkDir 获得工作目录
func GetWorkDir() string {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return d
}

// IsExists 文件或目录是否存在
func IsExists(filename string) bool {
	_, err := os.Lstat(filename)
	return !os.IsNotExist(err)
}

// MustRead 读取一个存在的文件的完整内容，注意：使用这个函数之前，请你自行确定文件不会很大，否则可能会对你的程序的性能和资源开销造成影响
func MustRead(filename string) []byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("读取文件 %s 出现错误：%+v", filename, err))
	}
	return b
}
