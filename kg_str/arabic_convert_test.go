package kg_str

import (
	"github.com/ccpwcn/kgo"
	"testing"
)

func TestArabicToChinese(t *testing.T) {
	type args[T kgo.IntUintBig] struct {
		arabic T
	}
	type testCase[T kgo.IntUintBig] struct {
		name string
		args args[T]
		want string
	}
	tests := []testCase[int]{
		{name: "0", args: args[int]{arabic: 0}, want: "零"},
		{name: "1", args: args[int]{arabic: 1}, want: "一"},
		{name: "9", args: args[int]{arabic: 9}, want: "九"},
		{name: "12", args: args[int]{arabic: 12}, want: "十二"},
		{name: "36", args: args[int]{arabic: 36}, want: "三十六"},
		{name: "123", args: args[int]{arabic: 123}, want: "一百二十三"},
		{name: "1234", args: args[int]{arabic: 1234}, want: "一千二百三十四"},
		{name: "12345", args: args[int]{arabic: 12345}, want: "一万二千三百四十五"},
		{name: "123456", args: args[int]{arabic: 123456}, want: "十二万三千四百五十六"},
		{name: "1234567", args: args[int]{arabic: 1234567}, want: "一百二十三万四千五百六十七"},
		{name: "12345678", args: args[int]{arabic: 12345678}, want: "一千二百三十四万五千六百七十八"},
		{name: "10", args: args[int]{arabic: 10}, want: "十"},
		{name: "20", args: args[int]{arabic: 20}, want: "二十"},
		{name: "101", args: args[int]{arabic: 101}, want: "一百零一"},
		{name: "110", args: args[int]{arabic: 110}, want: "一百一十"},
		{name: "118", args: args[int]{arabic: 118}, want: "一百一十八"},
		{name: "120", args: args[int]{arabic: 120}, want: "一百二十"},
		{name: "1002", args: args[int]{arabic: 1002}, want: "一千零二"},
		{name: "1030", args: args[int]{arabic: 1030}, want: "一千零三十"},
		{name: "10_0300", args: args[int]{arabic: 10_0300}, want: "十万零三百"},
		{name: "100_3004", args: args[int]{arabic: 100_3004}, want: "一百万三千零四"},
		{name: "1000_3004", args: args[int]{arabic: 1000_3004}, want: "一千万三千零四"},
		{name: "1001_3004", args: args[int]{arabic: 1001_3004}, want: "一千零一万三千零四"},
		{name: "1010_3004", args: args[int]{arabic: 1010_3004}, want: "一千零一十万三千零四"},
		{name: "1100_3004", args: args[int]{arabic: 1100_3004}, want: "一千一百万三千零四"},
		{name: "1000_3014", args: args[int]{arabic: 1000_3014}, want: "一千万三千零一十四"},
		{name: "1001_3104", args: args[int]{arabic: 1001_3104}, want: "一千零一万三千一百零四"},
		{name: "1001_3140", args: args[int]{arabic: 1001_3140}, want: "一千零一万三千一百四十"},
		{name: "1001_3141", args: args[int]{arabic: 1001_3141}, want: "一千零一万三千一百四十一"},
		{name: "101_3004", args: args[int]{arabic: 101_3004}, want: "一百零一万三千零四"},
		{name: "1_0030_0400", args: args[int]{arabic: 1_0030_0400}, want: "一亿零三十万零四百"},
		{name: "12_1234_5678", args: args[int]{arabic: 12_1234_5678}, want: "十二亿一千二百三十四万五千六百七十八"},
		{name: "131_1234_5678", args: args[int]{arabic: 131_1234_5678}, want: "一百三十一亿一千二百三十四万五千六百七十八"},
		{name: "1231_1234_5678", args: args[int]{arabic: 1231_1234_5678}, want: "一千二百三十一亿一千二百三十四万五千六百七十八"},
		{name: "1002_1234_5678", args: args[int]{arabic: 1002_1234_5678}, want: "一千零二亿一千二百三十四万五千六百七十八"},
		{name: "1020_1234_5678", args: args[int]{arabic: 1020_1234_5678}, want: "一千零二十亿一千二百三十四万五千六百七十八"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArabicToChinese[int](tt.args.arabic); got != tt.want {
				t.Errorf("ArabicToChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}
