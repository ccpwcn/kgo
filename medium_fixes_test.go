package kgo

import (
	"encoding/json"
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// 中等问题修复的回归测试
// ---------------------------------------------------------------------------

// Set 的 JSON 序列化应为原生数组，而不是被 encoding/json 二次转义成字符串。
// 旧实现用 MarshalText 承载 JSON，导致 Set 作为结构体字段时被序列化成 "[1,2,3]"（字符串）。
func TestSet_MarshalJSON_ProducesNativeArray(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// 1) 直接序列化：应为 JSON 数组
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("marshal set failed: %v", err)
	}
	got := strings.TrimSpace(string(b))
	if !strings.HasPrefix(got, "[") || !strings.HasSuffix(got, "]") {
		t.Errorf("Set 应序列化为 JSON 数组，实际=%s", got)
	}

	// 2) 关键回归：Set 作为结构体字段时，应内联为数组而非字符串
	type wrapper struct {
		Name string    `json:"name"`
		IDs  *Set[int] `json:"ids"`
	}
	b2, err := json.Marshal(wrapper{Name: "x", IDs: s})
	if err != nil {
		t.Fatalf("marshal wrapper failed: %v", err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b2, &m); err != nil {
		t.Fatalf("unmarshal wrapper failed: %v", err)
	}
	ids := strings.TrimSpace(string(m["ids"]))
	if !strings.HasPrefix(ids, "[") {
		t.Errorf("嵌套 Set 字段应为 JSON 数组，实际=%s（完整=%s）", ids, string(b2))
	}

	// 3) 反序列化回 Set，元素应完整
	var back Set[int]
	if err := json.Unmarshal([]byte(ids), &back); err != nil {
		t.Fatalf("unmarshal set failed: %v", err)
	}
	if back.Len() != 3 || !back.Contains(1) || !back.Contains(2) || !back.Contains(3) {
		t.Errorf("反序列化后元素丢失: %v", back.ToArray())
	}
}

// CopyFields 遇到"同 Kind 但类型不同"的字段（如 int 与 int64）应跳过而非 panic。
// 旧实现只比较 Kind，会尝试 reflect.Set(int -> int64) 从而 panic。
func TestCopyFields_IncompatibleSameKind_NoPanic(t *testing.T) {
	type Src struct {
		N int // Kind = Int
		S string
	}
	type Dst struct {
		N int64 // 同为 Kind = Int，但类型不同，不可赋值
		S string
	}
	src := Src{N: 42, S: "ok"}

	// 不应 panic
	dst := CopyFields[Src, Dst](src)

	if dst.S != "ok" {
		t.Errorf("类型相同的字段应被复制，got S=%q", dst.S)
	}
	if dst.N != 0 {
		t.Errorf("类型不兼容的字段应被跳过，got N=%d", dst.N)
	}
}

// fieldName 为空字符串时，三个结构体工具函数应安全返回零值，而不是 fieldName[0] 越界 panic。
func TestStruct_EmptyFieldName_NoPanic(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	users := []User{{Id: 1, Name: "a"}, {Id: 2, Name: "b"}}

	if got := JoinStructsField(users, ""); got != "" {
		t.Errorf("JoinStructsField 空字段名应返回空串，got %q", got)
	}
	if got := PickStructsField[User, int](users, ""); len(got) != 0 {
		t.Errorf("PickStructsField 空字段名应返回空切片，got %v", got)
	}
	if got := SliceGroupBy[User, int](users, ""); len(got) != 0 {
		t.Errorf("SliceGroupBy 空字段名应返回空 map，got %v", got)
	}
}

// Intersection/Union/Diff 现约束为 comparable，应支持可比较的结构体，并正确处理去重。
// （不可比较类型无法通过编译，故此处只验证可比较类型的正确行为。）
func TestCollectionOps_ComparableStructAndDedup(t *testing.T) {
	type Point struct{ X, Y int }
	s1 := []Point{{1, 1}, {2, 2}}
	s2 := []Point{{2, 2}, {3, 3}}

	if inter := Intersection(s1, s2); len(inter) != 1 || inter[0] != (Point{2, 2}) {
		t.Errorf("Intersection 结果错误: %v", inter)
	}
	if diff := Diff(s1, s2); len(diff) != 1 || diff[0] != (Point{1, 1}) {
		t.Errorf("Diff 结果错误: %v", diff)
	}
	if u := Union(s1, s2); len(u) != 3 {
		t.Errorf("Union 应去重后得到 3 个元素，got %v", u)
	}
	if u := Union([]int{1, 1, 2}, []int{2, 3, 3}); len(u) != 3 {
		t.Errorf("Union 去重失败，got %v", u)
	}
}

// IsExists：存在的文件/目录返回 true，不存在的路径返回 false（含不存在目录下的路径）。
func TestIsExists_NegativeAndPositive(t *testing.T) {
	if !IsExists("dir.go") {
		t.Error("dir.go 应存在")
	}
	if IsExists("definitely_not_exists_file_xyz.go") {
		t.Error("不存在的文件应返回 false")
	}
	if IsExists("no_such_dir_xyz/inner.go") {
		t.Error("不存在目录下的路径应返回 false")
	}
}
