package kgo

import "testing"

func TestClean(t *testing.T) {
	resources := map[string]string{
		"第一季度-结算": "第一季度_结算",
		"合作。":     "合作__",
		"数据（人）":   "数据_人_",
		"数量/人":    "数量_人",
	}
	for s, excepted := range resources {
		if actual := Clean(s); actual != excepted {
			t.Errorf("字符串清理，输入：%+v，预期：%+v，实际：%+v", s, excepted, actual)
		}
	}
}

func TestJoinElements(t *testing.T) {
	resources1 := map[string][]int{
		`1,2,3`: {1, 2, 3},
	}
	for expected, src := range resources1 {
		if actual := JoinElements(src); actual != expected {
			t.Fatalf("预期 %+v，实际值：%+v", expected, actual)
		}
	}
	var n1, n2, n3 = 1, 2, 3
	resources11 := map[string][]*int{
		`1,2,3`: {&n1, &n2, &n3},
	}
	for expected, src := range resources11 {
		if actual := JoinElements(src); actual != expected {
			t.Fatalf("预期 %+v，实际值：%+v", expected, actual)
		}
	}

	resources2 := map[string][]string{
		`a,b,c`: {"a", "b", "c"},
	}
	for expected, src := range resources2 {
		if actual := JoinElements(src); actual != expected {
			t.Fatalf("预期 %+v，实际值：%+v", expected, actual)
		}
	}
	var s1, s2, s3 = "a", "b", "c"
	resources22 := map[string][]*string{
		`a,b,c`: {&s1, &s2, &s3},
	}
	for expected, src := range resources22 {
		if actual := JoinElements(src); actual != expected {
			t.Fatalf("预期 %+v，实际值：%+v", expected, actual)
		}
	}

	resources3 := map[string][]bool{
		`true,false,true`: {true, false, true},
	}
	for expected, src := range resources3 {
		if actual := JoinElements(src); actual != expected {
			t.Fatalf("预期 %+v，实际值：%+v", expected, actual)
		}
	}
	var b1, b2, b3 = true, false, true
	resources33 := map[string][]*bool{
		`true,false,true`: {&b1, &b2, &b3},
	}
	for expected, src := range resources33 {
		if actual := JoinElements(src); actual != expected {
			t.Fatalf("预期 %+v，实际值：%+v", expected, actual)
		}
	}
}
