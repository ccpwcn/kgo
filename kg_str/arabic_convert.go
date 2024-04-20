package kg_str

import (
	"github.com/ccpwcn/kgo"
	"regexp"
)

type reCheck struct {
	pattern     *regexp.Regexp
	replacement string
}

var (
	chineseNumbers = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	chineseUnit    = []string{"", "十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
	regExpList     = []reCheck{
		{pattern: regexp.MustCompile(`^一十`), replacement: "十"},
		{pattern: regexp.MustCompile(`零[千百十]`), replacement: "零"},
		{pattern: regexp.MustCompile(`零+`), replacement: "零"},
		{pattern: regexp.MustCompile(`零零`), replacement: "零"},
		{pattern: regexp.MustCompile(`零$`), replacement: ""},
		{pattern: regexp.MustCompile(`零+万`), replacement: "万"},
		{pattern: regexp.MustCompile(`零+亿`), replacement: "亿"},
		{pattern: regexp.MustCompile(`亿万`), replacement: "亿零"},
	}
)

// ArabicToChinese 阿里数字转中文数字
func ArabicToChinese[T kgo.IntUintBig](arabic T) string {
	if arabic < 0 {
		return ""
	}
	if arabic > 9999_9999_9999 {
		return ""
	}
	var (
		dst   = ""
		count = 0
	)
	for arabic > 0 {
		dst = (chineseNumbers[arabic%10] + chineseUnit[count]) + dst
		arabic /= 10
		count++
	}
	for _, re := range regExpList {
		dst = re.pattern.ReplaceAllString(dst, re.replacement)
	}
	if dst == "" {
		dst = "零"
	}
	return dst
}
