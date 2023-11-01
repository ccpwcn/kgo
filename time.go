package kgo

import "time"

var location = time.Now().Location()

func NowStr() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}

// MonthStartTime 取得本月的开始时间
func MonthStartTime() time.Time {
	baseTime := time.Now()
	// 设置时间为该月1日起始时间
	return time.Date(baseTime.Year(), baseTime.Month(), 1, 0, 0, 0, 0, location)
}

// MonthEndTime 取得本月的结束时间
func MonthEndTime() time.Time {
	baseTime := time.Now()
	// 先前行到下1个月
	tmpMonth := baseTime.AddDate(0, 1, 0)
	// 切换到该月开始时间
	tmpMonth = time.Date(tmpMonth.Year(), tmpMonth.Month(), 1, 0, 0, 0, 0, location)
	// 再倒退1天，得到本月的结束时间
	tmpMonth = tmpMonth.AddDate(0, 0, -1)
	// 设置时间为该月最后一天最后一秒最后一毫秒
	return time.Date(tmpMonth.Year(), tmpMonth.Month(), tmpMonth.Day(), 23, 59, 59, 999999999, location)
}

// AnyMonthStartTime 以当前时间为基点，取任意月的开始时间（比如上个月、下个月），这在一些时间计算的时候很有用
// 参数：
//
//	delta 月份偏移量，1是下个月，-1是上个月
func AnyMonthStartTime(delta int) time.Time {
	baseTime := time.Now()
	// 先前行到指定月
	nextMonth := baseTime.AddDate(0, delta, 0)
	// 设置时间为该月1日起始时间
	return time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, location)
}

// AnyMonthEndTime 以当前时间为基点，取任意月的的结束时间（比如上个月、下个月），这在一些时间计算的时候很有用
// 参数：
//
//	delta 月份偏移量，1是下个月，-1是上个月
func AnyMonthEndTime(delta int) time.Time {
	baseTime := time.Now()
	// 先前行到指定月的下个月
	tmpMonth := baseTime.AddDate(0, delta+1, 0)
	// 切换到该月开始时间
	tmpMonth = time.Date(tmpMonth.Year(), tmpMonth.Month(), 1, 0, 0, 0, 0, location)
	// 再倒退1天，得到前一个月（即指定月）的结束时间
	tmpMonth = tmpMonth.AddDate(0, 0, -1)
	// 设置时间为该月最后一天最后一秒最后一毫秒最后一纳秒
	return time.Date(tmpMonth.Year(), tmpMonth.Month(), tmpMonth.Day(), 23, 59, 59, 999999999, location)
}
