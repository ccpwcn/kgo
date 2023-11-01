package kgo

import (
	"testing"
)

func TestAnyMonthEndTime(t *testing.T) {
	// 测下个月
	next := AnyMonthEndTime(1)
	t.Logf("下个月结束时间：%+v", next)
	// 测下两个月
	nextNext := AnyMonthEndTime(2)
	t.Logf("下两个月结束时间：%+v", nextNext)
	// 测上个月
	before := AnyMonthEndTime(-1)
	t.Logf("上个月结束时间：%+v", before)
	// 测上两个月
	beforeBefore := AnyMonthEndTime(-2)
	t.Logf("上两个月结束时间：%+v", beforeBefore)
}

func TestAnyMonthStartTime(t *testing.T) {
	// 测下个月
	next := AnyMonthStartTime(1)
	t.Logf("下个月开始时间：%+v", next)
	// 测下两个月
	nextNext := AnyMonthStartTime(2)
	t.Logf("下两个月开始时间：%+v", nextNext)
	// 测上个月
	before := AnyMonthStartTime(-1)
	t.Logf("上个月开始时间：%+v", before)
	// 测上两个月
	beforeBefore := AnyMonthStartTime(-2)
	t.Logf("上两个月开始时间：%+v", beforeBefore)
}

func TestMonthEndTime(t *testing.T) {
	monthEnd := MonthEndTime()
	t.Logf("本月结束时间：%+v", monthEnd)
}

func TestMonthStartTime(t *testing.T) {
	monthStart := MonthStartTime()
	t.Logf("本月开始时间：%+v", monthStart)
}

func TestNowStr(t *testing.T) {
	current := NowStr()
	t.Logf("当前时间：%+v", current)
}
