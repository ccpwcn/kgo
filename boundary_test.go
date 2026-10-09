package kgo

import (
	"strconv"
	"testing"
)

// 本文件集中覆盖边界与异常路径，防止此前修复的缺陷回归。

// --- Stack：空栈必须返回 (零值, error) 而不是 panic ---

func TestBoundary_Stack_EmptyPopPeek(t *testing.T) {
	s := NewStack[int]()
	v, err := s.Pop()
	if err == nil {
		t.Errorf("空栈 Pop 期望返回 error")
	}
	if v != 0 {
		t.Errorf("空栈 Pop 期望返回零值，实际 %v", v)
	}
	v2, err2 := s.Peek()
	if err2 == nil {
		t.Errorf("空栈 Peek 期望返回 error")
	}
	if v2 != 0 {
		t.Errorf("空栈 Peek 期望返回零值，实际 %v", v2)
	}
}

func TestBoundary_Stack_EmptyPopPeek_StringType(t *testing.T) {
	s := NewStack[string]()
	v, err := s.Pop()
	if err == nil || v != "" {
		t.Errorf("空栈(string) Pop 期望 (\"\", error)，实际 (%q, %v)", v, err)
	}
}

// --- Nums2Strings：必须返回等长、无逗号的字符串数组 ---

func TestBoundary_Nums2Strings(t *testing.T) {
	got := Nums2Strings([]int{1, 2, 3})
	if len(got) != 3 {
		t.Fatalf("期望长度 3，实际 %d (%#v)", len(got), got)
	}
	for i, want := range []string{"1", "2", "3"} {
		if got[i] != want {
			t.Errorf("索引 %d 期望 %q，实际 %q", i, want, got[i])
		}
	}
}

func TestBoundary_Nums2Strings_Pointer(t *testing.T) {
	a, b := 10, 20
	got := Nums2Strings([]*int{&a, &b})
	if len(got) != 2 || got[0] != "10" || got[1] != "20" {
		t.Errorf("指针数组转换错误：%#v", got)
	}
}

func TestBoundary_Nums2Strings_Empty(t *testing.T) {
	got := Nums2Strings([]int{})
	if len(got) != 0 {
		t.Errorf("空输入期望空输出，实际 %#v", got)
	}
}

func TestBoundary_Nums2Strings_Float(t *testing.T) {
	got := Nums2Strings([]float64{1.5})
	if len(got) != 1 || got[0] != strconv.FormatFloat(1.5, 'g', -1, 64) {
		// %v 对 1.5 输出 "1.5"
		if got[0] != "1.5" {
			t.Errorf("float 转换错误：%#v", got)
		}
	}
}

// --- SameElements：需按多重集比较，重复元素次数不同应判为不同 ---

func TestBoundary_SameElements_Duplicates(t *testing.T) {
	if SameElements([]int{1, 1, 2}, []int{1, 2, 2}) {
		t.Errorf("[1,1,2] 与 [1,2,2] 应判为不同")
	}
	if !SameElements([]int{1, 2, 2}, []int{2, 1, 2}) {
		t.Errorf("[1,2,2] 与 [2,1,2] 应判为相同")
	}
	if !SameElements([]int{1, 1, 1}, []int{1, 1, 1}) {
		t.Errorf("[1,1,1] 与 [1,1,1] 应判为相同")
	}
	if SameElements([]int{1, 1, 1}, []int{1, 1, 2}) {
		t.Errorf("[1,1,1] 与 [1,1,2] 应判为不同")
	}
}

func TestBoundary_SameElements_LengthMismatch(t *testing.T) {
	if SameElements([]int{1, 2}, []int{1, 2, 3}) {
		t.Errorf("长度不同应判为不同")
	}
	if SameElements([]int{}, nil) == false {
		// 两者长度都为 0，应判为相同
		t.Errorf("两个空切片应判为相同")
	}
}

// --- SplitCounter：整除时不应产生多余空段 ---

func TestBoundary_SplitCounter_ExactMultiple(t *testing.T) {
	items := SplitCounter[int](5, 10)
	if len(items) != 2 {
		t.Fatalf("count=10,pageSize=5 期望 2 段，实际 %d", len(items))
	}
	for i, it := range items {
		if cap(it) != 5 {
			t.Errorf("第 %d 段容量期望 5，实际 %d", i, cap(it))
		}
	}
}

func TestBoundary_SplitCounter_WithRemainder(t *testing.T) {
	items := SplitCounter[int](5, 13)
	if len(items) != 3 {
		t.Fatalf("count=13,pageSize=5 期望 3 段，实际 %d", len(items))
	}
	wantCaps := []int{5, 5, 3}
	for i, it := range items {
		if cap(it) != wantCaps[i] {
			t.Errorf("第 %d 段容量期望 %d，实际 %d", i, wantCaps[i], cap(it))
		}
	}
}

func TestBoundary_SplitCounter_Invalid(t *testing.T) {
	if got := SplitCounter[int](0, 10); got != nil {
		t.Errorf("pageSize=0 期望 nil，实际 %#v", got)
	}
	if got := SplitCounter[int](5, 0); got != nil {
		t.Errorf("count=0 期望 nil，实际 %#v", got)
	}
	if got := SplitCounter[int](-1, -1); got != nil {
		t.Errorf("负参数期望 nil，实际 %#v", got)
	}
}

// --- MaskChineseNameEx：left/right 越界不得 panic ---

func TestBoundary_MaskChineseNameEx_OutOfRange(t *testing.T) {
	cases := []struct {
		name        string
		left, right int
	}{
		{"张", 1, 1},   // left+right > size
		{"张", 5, 5},   // 远超长度
		{"张三", -1, 5}, // 负数 + 超长
		{"张一二", 1, 1}, // 正常
		{"", 1, 1},    // 空串
	}
	for _, c := range cases {
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("MaskChineseNameEx(%q,%d,%d) panic: %v", c.name, c.left, c.right, r)
				}
			}()
			_ = MaskChineseNameEx(c.name, c.left, c.right)
		}()
	}
}

func TestBoundary_MaskChineseNameEx_Normal(t *testing.T) {
	if got := MaskChineseNameEx("张一二", 1, 1); got != "张*二" {
		t.Errorf("期望 张*二，实际 %q", got)
	}
	if got := MaskChineseNameEx("张", 1, 1); got != "张" {
		t.Errorf("单字符越界钳制后期望 张，实际 %q", got)
	}
}

// --- Snowflake：参数越界返回 error 而非 panic ---

// 未初始化时应 panic（必须排在任何成功的 InitSnowflake 之前运行；
// 若已被其它测试初始化则跳过，避免依赖测试执行顺序）
func TestBoundary_Snowflake_UninitializedPanics(t *testing.T) {
	if snowflakeOf() != nil {
		t.Skip("snowflake 已被初始化，跳过未初始化路径")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("未初始化调用 SnowflakeId 期望 panic，但没有")
		}
	}()
	_ = SnowflakeId()
}

func TestBoundary_InitSnowflake_InvalidParams(t *testing.T) {
	if err := InitSnowflake(-1, 1); err == nil {
		t.Errorf("workerId=-1 期望返回 error")
	}
	if err := InitSnowflake(workerIdMax+1, 1); err == nil {
		t.Errorf("workerId 超上限期望返回 error")
	}
	if err := InitSnowflake(1, -1); err == nil {
		t.Errorf("dataCenterId=-1 期望返回 error")
	}
	if err := InitSnowflake(1, datacenterIdMax+1); err == nil {
		t.Errorf("dataCenterId 超上限期望返回 error")
	}
	// 合法边界值不应报错（此调用会成功初始化全局单例，故排在未初始化用例之后）
	if err := InitSnowflake(workerIdMax, datacenterIdMax); err != nil {
		t.Errorf("合法上限值不应报错：%v", err)
	}
}
