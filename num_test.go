package kgo

import "testing"

func TestNumJoinStr(t *testing.T) {
	resources1 := map[string][]int{
		`[1,2,3]`:          {1, 2, 3},
		`[10,20,30]`:       {10, 20, 30},
		`[100,200,300]`:    {100, 200, 300},
		`[1001,2002,3003]`: {1001, 2002, 3003},
	}
	for expected, src := range resources1 {
		if actual := NumJoinStr(src); actual != expected {
			t.Fatalf("预期 %+v，实际值：%+v", expected, actual)
		}
	}
	var n1, n2, n3 = 111, 222, 333
	resources11 := map[string][]*int{
		`[111,222,333]`: {&n1, &n2, &n3},
	}
	for expected, src := range resources11 {
		if actual := NumJoinStr(src); actual != expected {
			t.Fatalf("预期 %+v，实际值：%+v", expected, actual)
		}
	}
	resources2 := map[string][]float64{
		`[1.1,2.2,3.3]`:    {1.1, 2.2, 3.3},
		`[10.1,20.2,30.3]`: {10.1, 20.2, 30.3},
	}
	for expected, src := range resources2 {
		if actual := NumJoinStr(src); actual != expected {
			t.Fatalf("预期 %+v，实际值：%+v", expected, actual)
		}
	}
}
