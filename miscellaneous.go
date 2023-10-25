package kgo

import "runtime"

// RunFuncName 获取正在运行的函数名，调用也也可以
func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
