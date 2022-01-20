package _6_第51题一周中的第几天模拟版本

var monthOfDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}
var DayOfString = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func dayOfTheWeek(day int, month int, year int) string {
	// 1970年12月31号是星期四, 我们只需要对时间进行计算
	// 算出此时间距离1970年12月31号有多少天即可
	// 然后再加3代表重回星期日, 最后对7取模即可
	sum := 0
	// 计算这一年到1970年贡献的天数
	for i := 1971; i < year; i++ {
		if isLeapYear(i) {
			sum += 366
		} else {
			sum += 365
		}
	}
	// 计算月份对于天数的贡献
	for i := 0; i < month-1; i++ {
		sum += monthOfDays[i]
	}
	// 计算天数对于月份的贡献
	if month > 2 && isLeapYear(year) {
		sum += day + 1
	} else {
		sum += day
	}
	sum += 3
	return DayOfString[sum%7]
}

func isLeapYear(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}
