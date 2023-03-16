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
- [x] Map相关操作
  - HasKey 是否拥有某个Key 
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

## 非泛型接口
此接口只会返回int64类型的ID。具体使用方法参见`[snowflake_test.go](snowflake_test.go)`中的代码示例。

测试方法：在PowerShell中执行的测试命令：
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
BenchmarkSnowflake_Concurrent_Id-8             9         203077967 ns/op        21983367 B/op     625820 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            12         277161517 ns/op        31229477 B/op     761023 allocs/op
BenchmarkSnowflake_Concurrent_Id-8             9         204514756 ns/op        21902075 B/op     625815 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            12         268109500 ns/op        30850743 B/op     761026 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            10         220143410 ns/op        22856056 B/op     671238 allocs/op
BenchmarkSnowflake_Concurrent_Id-8             9         227427922 ns/op        21570741 B/op     625823 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            10         245171650 ns/op        22504872 B/op     671258 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            10         232327030 ns/op        22703912 B/op     671208 allocs/op
BenchmarkSnowflake_Concurrent_Id-8             9         223790600 ns/op        21771869 B/op     625841 allocs/op
BenchmarkSnowflake_Concurrent_Id-8            10         239567270 ns/op        23161156 B/op     671290 allocs/op
PASS                                                                                                              
ok      github.com/ccpwcn/kgo   24.803s
```

## 泛型接口
泛型接口支持返回int64和string两种类型的ID，你可以在调用函数的时候指定返回类型，具体调用代码参见`[snowflake_generic_test.go](snowflake_generic_test.go)`中的代码示例。

测试方法：在PowerShell中执行的测试命令：
```shell
go test -v -bench="Benchmark_GenericSnowflake_Concurrent_Id" -run=none -count=10 -benchmem
```
命令输出：
```text
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
Benchmark_GenericSnowflake_Concurrent_Id
Benchmark_GenericSnowflake_Concurrent_Id-8                    10         224993810 ns/op        24819956 B/op     893709 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id-8                    12         282375483 ns/op        32965658 B/op    1027759 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id-8                    12         283087400 ns/op        33369010 B/op    1027736 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id-8                    12         287084958 ns/op        33401450 B/op    1027784 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id-8                    12         284687817 ns/op        33059180 B/op    1027716 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id-8                    12         287791083 ns/op        32944784 B/op    1027756 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id-8                    10         234368130 ns/op        24079393 B/op     893518 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id-8                    10         229153660 ns/op        24258326 B/op     893513 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id-8                    12         292083508 ns/op        33392757 B/op    1027721 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id-8                    10         230679260 ns/op        24353888 B/op     893491 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id_String
Benchmark_GenericSnowflake_Concurrent_Id_String-8             10         278951190 ns/op        28652905 B/op    1004702 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id_String-8              8         248918450 ns/op        26022041 B/op     847370 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id_String-8             10         275623760 ns/op        28527036 B/op    1004687 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id_String-8             10         280000980 ns/op        28835009 B/op    1004686 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id_String-8             10         279923380 ns/op        29070018 B/op    1004786 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id_String-8             10         294371500 ns/op        28708047 B/op    1004654 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id_String-8             10         276980950 ns/op        28606474 B/op    1004658 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id_String-8             12         354999700 ns/op        39017543 B/op    1161183 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id_String-8             12         356444367 ns/op        38855564 B/op    1161128 allocs/op
Benchmark_GenericSnowflake_Concurrent_Id_String-8             10         281144360 ns/op        28855772 B/op    1004697 allocs/op
PASS
ok      github.com/ccpwcn/kgo   62.156s
```

## 综合性能
经过多次冷启动测试和热启动测试，有如下结论：
1. 总体上各接口的综合性能表现很是稳定，不会有明显的波动。
2. 非泛型返回int64类型的ID是最快的。
3. 泛型int64类型ID接口比非泛型接口慢不到5%，可以忽略不计。这也从侧面反应出Golang的泛型设计还是非常优秀的，性能损失非常小。
4. 泛型string类型ID接口比泛型int64接口慢约15%到18%，比非泛型ID接口慢20%。这也可以理解，int64类型的大小是确定的，它在栈上，string类型对于程序来说大小是不确定的，它在堆上，自然就慢了。

综合一张图如下所示：
![雪花算法性能分析.png](assets/snowflake_algorithm_performance_analysis.png)

> 之所以提供string类型ID的方法，还是为了便于一些场合性能要求不那么高，却要以string类型保存ID的地方，就不必非常麻烦的每次都要转换一下或者自己再封装一下了。