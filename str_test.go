package kgo

import (
	"testing"
	"unicode/utf8"
)

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
			t.Errorf("预期 %+v，实际值：%+v", expected, actual)
		}
	}
	var n1, n2, n3 = 1, 2, 3
	resources11 := map[string][]*int{
		`1,2,3`: {&n1, &n2, &n3},
	}
	for expected, src := range resources11 {
		if actual := JoinElements(src); actual != expected {
			t.Errorf("预期 %+v，实际值：%+v", expected, actual)
		}
	}

	resources2 := map[string][]string{
		`a,b,c`: {"a", "b", "c"},
	}
	for expected, src := range resources2 {
		if actual := JoinElements(src); actual != expected {
			t.Errorf("预期 %+v，实际值：%+v", expected, actual)
		}
	}
	var s1, s2, s3 = "a", "b", "c"
	resources22 := map[string][]*string{
		`a,b,c`: {&s1, &s2, &s3},
	}
	for expected, src := range resources22 {
		if actual := JoinElements(src); actual != expected {
			t.Errorf("预期 %+v，实际值：%+v", expected, actual)
		}
	}

	resources3 := map[string][]bool{
		`true,false,true`: {true, false, true},
	}
	for expected, src := range resources3 {
		if actual := JoinElements(src); actual != expected {
			t.Errorf("预期 %+v，实际值：%+v", expected, actual)
		}
	}
	var b1, b2, b3 = true, false, true
	resources33 := map[string][]*bool{
		`true,false,true`: {&b1, &b2, &b3},
	}
	for expected, src := range resources33 {
		if actual := JoinElements(src); actual != expected {
			t.Errorf("预期 %+v，实际值：%+v", expected, actual)
		}
	}
}

func TestB2S(t *testing.T) {
	b := []byte{'A', 'B', 'C'}
	s := B2S(b)
	if s != "ABC" {
		t.Errorf("预期 %+v，实际值：%+v", "ABC", s)
	}
}

func TestR2S_ASCII(t *testing.T) {
	b := []rune{'A', 'B', 'C'}
	s := R2S(b)
	if s != "ABC" {
		t.Errorf("预期 %+v，实际值：%+v", "ABC", s)
	}
}

func TestR2S_Chinese(t *testing.T) {
	b := []rune{'一', '二', '三'}
	s := R2S(b)
	if s != "一二三" {
		t.Errorf("预期 %+v，实际值：%+v", "ABC", s)
	}
}

func TestR2S_Mix(t *testing.T) {
	b := []rune{'一', '1', '二', '2', '三', '3'}
	s := R2S(b)
	if s != "一1二2三3" {
		t.Errorf("预期 %+v，实际值：%+v", "ABC", s)
	}
}

func TestS2B(t *testing.T) {
	s := "!@#"
	b := S2B(s)
	if len(b) != 3 || s[0] != '!' || s[1] != '@' || s[2] != '#' {
		t.Errorf("预期 %+v，实际值：%+v", []byte{'!', '@', '#'}, b)
	}
}

func TestS2B2S(t *testing.T) {
	s := "!@#"
	b := S2B(s)
	s2 := B2S(b)
	if s2 != s {
		t.Errorf("预期 %+v，实际值：%+v", s, s2)
	}
}

func TestMaskChineseName(t *testing.T) {
	name := "张三"
	excepted := "张*"
	actual := MaskChineseName(name)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseName1(t *testing.T) {
	name := "张三一"
	excepted := "张**"
	actual := MaskChineseName(name)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseName2(t *testing.T) {
	name := "张三三三"
	excepted := "张***"
	actual := MaskChineseName(name)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseName4(t *testing.T) {
	name := "hello"
	excepted := "h****"
	actual := MaskChineseName(name)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseNameEx(t *testing.T) {
	name := "张三"
	excepted := "*三"
	actual := MaskChineseNameEx(name, 0, 1)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseNameEx1(t *testing.T) {
	name := "张三一"
	excepted := "*三一"
	actual := MaskChineseNameEx(name, 0, 2)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseNameEx2(t *testing.T) {
	name := "张三三三"
	excepted := "*三三三"
	actual := MaskChineseNameEx(name, 0, 3)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseNameEx4(t *testing.T) {
	name := "hello"
	excepted := "*ello"
	actual := MaskChineseNameEx(name, 0, 4)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseNameEx5(t *testing.T) {
	name := "张一二"
	excepted := "张*二"
	actual := MaskChineseNameEx(name, 1, 1)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseNameEx6(t *testing.T) {
	name := "张一二"
	excepted := "*一二"
	actual := MaskChineseNameEx(name, 0, utf8.RuneCountInString(name)-1)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseMobile(t *testing.T) {
	tel := "13012345678"
	excepted := "130****5678"
	actual := MaskChineseMobile(tel)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChsCard34(t *testing.T) {
	idCard := "101101199010101234"
	excepted := "101***********1234"
	actual := MaskChineseIdCard34(idCard)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChsCard64(t *testing.T) {
	idCard := "101101199010101234"
	excepted := "101101********1234"
	actual := MaskChineseIdCard64(idCard)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChsCard11(t *testing.T) {
	idCard := "101101199010101234"
	excepted := "1****************4"
	actual := MaskChineseIdCard11(idCard)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseIdCard(t *testing.T) {
	idCard := "101101199010101234"
	excepted := "101***********1234"
	actual := MaskChineseIdCard(idCard, 3, 4)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseIdCard2(t *testing.T) {
	idCard := "10110119901010123X"
	excepted := "101***********123X"
	actual := MaskChineseIdCard(idCard, 3, 4)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskChineseIdCard3(t *testing.T) {
	idCard := "10110119901010123x"
	excepted := "101***********123x"
	actual := MaskChineseIdCard(idCard, 3, 4)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskAnyString(t *testing.T) {
	s := "一二三四五"
	excepted := "一***五"
	actual := MaskAnyString(s, 1, 1)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskAnyString1(t *testing.T) {
	s := "一二三四五"
	excepted := "一****"
	actual := MaskAnyString(s, 1, 0)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskAnyString2(t *testing.T) {
	s := "一二三四五"
	excepted := "****五"
	actual := MaskAnyString(s, 0, 1)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskAnyString3(t *testing.T) {
	s := "一二三四五六七八九十"
	excepted := "一二三******十"
	actual := MaskAnyString(s, 3, 1)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestReverseStringEnglish(t *testing.T) {
	s := "abcd"
	excepted := "dcba"
	actual := ReverseString(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestReverseStringChinese(t *testing.T) {
	s := "一二三四"
	excepted := "四三二一"
	actual := ReverseString(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestReverseStringMix(t *testing.T) {
	s := "一1二2三3"
	excepted := "3三2二1一"
	actual := ReverseString(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestEnglishWordsCount(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "hello_word",
			args: args{s: "hello_word"},
			want: 2,
		},
		{
			name: "hello word",
			args: args{s: "hello word"},
			want: 2,
		},
		{
			name: "statement",
			args: args{s: "Hello, this is a sample text with some English words."},
			want: 10,
		},
	}
	for _, test := range tests {
		actual := EnglishWordsCount(test.args.s)
		if actual != test.want {
			t.Errorf("预期 %v，实际值：%v", test.want, actual)
		}
	}
}
