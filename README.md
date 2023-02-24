# kgo 快速的Go开发工具简易工具集
![GitHub](https://img.shields.io/github/license/ccpwcn/kgo)
![GitHub tests (latest SemVer)](https://img.shields.io/badge/tests-passed-brightgreen)
![GitHub Workflow Status (with branch)](https://img.shields.io/github/actions/workflow/status/ccpwcn/kgo/go.yml)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/ccpwcn/kgo)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ccpwcn/kgo)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/ccpwcn/kgo)


想来想去，取了这么一个简单易用的名字，把一些常见的功能放上去，免得在各个项目中分别写一遍，同时开源出来，希望给同是从事Golang开发的朋友们有一丢丢小小的帮助。

> 本工具集仅依赖Go SDK，暂时不考虑依赖其他第三方。

# 前置条件
因为其中使用到了泛型，所以需要在Go 1.18以上版本中运行

# 功能目录
- [x] 目录文件相关操作
- [x] 数字相关操作
- [x] 字符串相关操作
- [ ] Map相关操作
- [x] 结构体相关操作
- [x] 雪花算法

# 引入和安装
非常简单，一个指令搞定：
```shell
go get -u github.com/ccpwcn/kgo
```

# 使用方法
请查看单元测试，那里就是测试代码，或者直接查看源码，都是非常简单的引用类，后面东西多了，复杂了，我再加上专门的使用说明文档吧。

# 雪花算法性能表现
为了确保在生产环境使用没有问题，我特意写了一个性能测试，好好对雪花算法进行了压力测试。
在PowerShell中执行的测试命令：
```shell
go test -v -bench="BenchmarkSnowflake_Concurrent_Id" -run=none -count=10 -benchmem
```
命令输出：
```text
goos: windows
goarch: amd64
pkg: github.com/ccpwcn/kgo
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkSnowflake_Concurrent_Id
BenchmarkSnowflake_Concurrent_Id-8             9         217772189 ns/op        21733468 B/op     625873 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            10         231429490 ns/op        22388100 B/op     671228 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            10         231598560 ns/op        22293514 B/op     671233 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            10         230639500 ns/op        22323937 B/op     671239 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            12         292305067 ns/op        31733692 B/op     761012 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            10         236379090 ns/op        22452506 B/op     671243 allocs/op
BenchmarkSnowflake_Concurrent_Id-8             9         212681622 ns/op        21538135 B/op     625780 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            10         232043890 ns/op        22591533 B/op     671354 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            10         233019430 ns/op        22484507 B/op     671258 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            10         234118470 ns/op        22525298 B/op     671292 allocs/op
PASS
ok      github.com/ccpwcn/kgo   24.862s
```
表现很是稳定。