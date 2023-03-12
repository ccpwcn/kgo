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
  - GetExeDir 获得可执行程序所在目录
  - GetWorkDir 获得工作目录
  - IsExists 文件是否存在
  - MustRead 读取一个文件内容到[]byte中，确定文件很小时，直接使用这个方法，很省事
- [x] 数字相关操作
  - NumJoinStr 将一个数字（int、float）数组合并成一个字符串数组
  - Nums2Strings 将一个数字（int、float）数组合转换一个字符串数组
- [x] 字符串相关操作
  - Clean 清理字符串，将其中的特殊字符（括号、标点符号、转义字符等等）一律转换成下划线
  - JoinElements 将任意基本类型的数组使用英文逗号拼接成一个字符串
- [x] 本地缓存相关操作
  - Set 设置缓存项，支持仅设置缓存，也支持同时给缓存添加一个过期时间
  - Get 获得缓存内容
  - Delete 删除缓存
  - Exists 判断缓存项是否存在
  - Size 获得缓存大小
- [ ] Map相关操作
- [x] 结构体相关操作
  - JoinStructsField 将任意结构体数组中的指定字段的值使用英文逗号拼接成一个字符串，例如：用户列表中，所有用户ID拼成一个字符串
  - PickStructsField 将任意结构体数组中的指定字段的值提取出来形成一个保持原类型的数组，例如：用户列表中，所有用户ID提取成一个用户ID数组
- [x] 雪花算法
  - 通用实现方法，初始化一次，到处随时获得ID，多goroutine并发安全

🍕🍕🍕更多使用方法请参见测试用例或代码注释，都非常简单易用。

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
BenchmarkSnowflake_Concurrent_Id-8                    10         228476450 ns/op        24371260 B/op     893551 allocs/op
BenchmarkSnowflake_Concurrent_Id-8                    10         221438730 ns/op        24255018 B/op     893571 allocs/op
BenchmarkSnowflake_Concurrent_Id-8                    10         228759660 ns/op        24163129 B/op     893506 allocs/op
BenchmarkSnowflake_Concurrent_Id-8                    10         224021280 ns/op        24152733 B/op     893484 allocs/op
BenchmarkSnowflake_Concurrent_Id-8                    12         282877567 ns/op        33124897 B/op    1027776 allocs/op
BenchmarkSnowflake_Concurrent_Id-8                    12         282631367 ns/op        33600299 B/op    1027714 allocs/op
BenchmarkSnowflake_Concurrent_Id-8                    12         282288167 ns/op        33314116 B/op    1027732 allocs/op
BenchmarkSnowflake_Concurrent_Id-8                    10         221747070 ns/op        24494428 B/op     893501 allocs/op
BenchmarkSnowflake_Concurrent_Id-8                    12         284988817 ns/op        33560668 B/op    1027819 allocs/op
BenchmarkSnowflake_Concurrent_Id-8                    12         282686433 ns/op        33517448 B/op    1027905 allocs/op
BenchmarkSnowflake_Concurrent_Id_String                                                                                   
BenchmarkSnowflake_Concurrent_Id_String-8             12         352885200 ns/op        38611372 B/op    1161149 allocs/op
BenchmarkSnowflake_Concurrent_Id_String-8             10         273024470 ns/op        28924257 B/op    1004706 allocs/op
BenchmarkSnowflake_Concurrent_Id_String-8             10         290659680 ns/op        28560619 B/op    1004655 allocs/op
BenchmarkSnowflake_Concurrent_Id_String-8             10         276047140 ns/op        28892588 B/op    1004743 allocs/op
BenchmarkSnowflake_Concurrent_Id_String-8              8         249127200 ns/op        26594543 B/op     847427 allocs/op
BenchmarkSnowflake_Concurrent_Id_String-8             10         288484240 ns/op        29090812 B/op    1004658 allocs/op
BenchmarkSnowflake_Concurrent_Id_String-8             10         304520150 ns/op        28779200 B/op    1004660 allocs/op
BenchmarkSnowflake_Concurrent_Id_String-8             10         269651730 ns/op        28522777 B/op    1004678 allocs/op
BenchmarkSnowflake_Concurrent_Id_String-8             10         276155210 ns/op        28976923 B/op    1004669 allocs/op
BenchmarkSnowflake_Concurrent_Id_String-8             10         281306900 ns/op        28545998 B/op    1004693 allocs/op
PASS
ok      github.com/ccpwcn/kgo   59.227s
```
表现很是稳定。