package kg_str

import (
	"testing"
	"unicode/utf8"
)

func TestMaskShortStringWithLeft(t *testing.T) {
	masker := NewMasker(WithLeft(1))
	name := "张三"
	excepted := "张*"
	actual := masker.Mask(name)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskShortStringWithRight(t *testing.T) {
	masker := NewMasker(WithRight(1))
	name := "张三"
	excepted := "*三"
	actual := masker.Mask(name)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskShortStringOnlyLeft(t *testing.T) {
	masker := NewMasker(WithLeft(1))
	name := "张三"
	excepted := "张*"
	actual := masker.Mask(name)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskShortStringOnlyRight(t *testing.T) {
	name := "张三三"
	excepted := "*三三"
	masker := NewMasker(WithRight(utf8.RuneCountInString(name) - 1))
	actual := masker.Mask(name)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskShortStringOnlyRight1(t *testing.T) {
	name := "张三三"
	excepted := "**三"
	masker := NewMasker(WithRight(1))
	actual := masker.Mask(name)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskShortStringWithRightExpended(t *testing.T) {
	name := "张三"
	excepted := "***三"
	masker := NewMasker(WithRight(1), WithMaskCount(3))
	actual := masker.Mask(name)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskShortStringWithLeftExpended(t *testing.T) {
	name := "张三"
	excepted := "张***"
	masker := NewMasker(WithLeft(1), WithMaskCount(3))
	actual := masker.Mask(name)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskDefault(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "最简人名", args: args{s: "张三"}, want: "张*三"},
		{name: "常规人名", args: args{s: "张一二"}, want: "张*二"},
		{name: "常规字符串", args: args{s: "一二三四五六七八九十"}, want: "一********十"},
	}

	masker := NewMasker()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := masker.Mask(test.args.s); got != test.want {
				t.Errorf("Mask() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestMaskDefaultLeftAndRight(t *testing.T) {
	masker := NewMasker(WithMaskChar('?'), WithMaskCount(3))
	s := "一二三四五六七八九十"
	excepted := "一???十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskDefaultLeftAndRightAndStar(t *testing.T) {
	masker := NewMasker(WithMaskCount(3))
	s := "一二三四五六七八九十"
	excepted := "一***十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskDefaultLeftAndRightAndMaskChar(t *testing.T) {
	masker := NewMasker(WithMaskChar(','))
	s := "一二三四五六七八九十"
	excepted := "一,,,,,,,,十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskDefaultMaskCharCount(t *testing.T) {
	masker := NewMasker(WithMaskChar('〇'))
	s := "一二三四五六七八九十"
	excepted := "一〇〇〇〇〇〇〇〇十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithLeft(t *testing.T) {
	masker := NewMasker(WithLeft(3))
	s := "一二三四五六七八九十"
	excepted := "一二三*******"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithLeftAndMaskCount(t *testing.T) {
	masker := NewMasker(WithLeft(3), WithMaskCount(3))
	s := "一二三四五六七八九十"
	excepted := "一二三***"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithLeftAndMaskChar(t *testing.T) {
	masker := NewMasker(WithLeft(3), WithMaskChar('❎'))
	s := "一二三四五六七八九十"
	excepted := "一二三❎❎❎❎❎❎❎"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithLeftAndMaskCharAndMaskCount(t *testing.T) {
	masker := NewMasker(WithLeft(3), WithMaskCount(3), WithMaskChar('❎'))
	s := "一二三四五六七八九十"
	excepted := "一二三❎❎❎"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithRight(t *testing.T) {
	masker := NewMasker(WithRight(3))
	s := "一二三四五六七八九十"
	excepted := "*******八九十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithRightAndMaskCount(t *testing.T) {
	masker := NewMasker(WithRight(3), WithMaskCount(3))
	s := "一二三四五六七八九十"
	excepted := "***八九十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithRightAndMaskChar(t *testing.T) {
	masker := NewMasker(WithRight(3), WithMaskChar('?'))
	s := "一二三四五六七八九十"
	excepted := "???????八九十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithRightAndMaskCharAndMaskCount(t *testing.T) {
	masker := NewMasker(WithRight(3), WithMaskChar('?'), WithMaskCount(3))
	s := "一二三四五六七八九十"
	excepted := "???八九十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithLeftAndRight(t *testing.T) {
	masker := NewMasker(WithLeft(3), WithRight(3))
	s := "一二三四五六七八九十"
	excepted := "一二三****八九十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithLeftAndRightAndMaskCount(t *testing.T) {
	masker := NewMasker(WithLeft(3), WithRight(3), WithMaskCount(2))
	s := "一二三四五六七八九十"
	excepted := "一二三**八九十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithLeftAndRightAndMaskChar(t *testing.T) {
	masker := NewMasker(WithLeft(3), WithRight(3), WithMaskChar('.'))
	s := "一二三四五六七八九十"
	excepted := "一二三....八九十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithLeftAndRightAndMaskCharAndMaskCount(t *testing.T) {
	masker := NewMasker(WithLeft(3), WithRight(2), WithMaskChar('❓'), WithMaskCount(3))
	s := "一二三四五六七八九十"
	excepted := "一二三❓❓❓九十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}

func TestMaskWithLeftAndRightAndMaskCharAndMaskCount2(t *testing.T) {
	masker := NewMasker(WithLeft(3), WithRight(3), WithMaskChar('�'), WithMaskCount(2))
	s := "一二三四五六七八九十"
	excepted := "一二三��八九十"
	actual := masker.Mask(s)
	if actual != excepted {
		t.Errorf("预期 %v，实际值：%v", excepted, actual)
	}
}
