package main

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	var days = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	str := strings.Split(date, "-")
	// 可能是闰年
	year, _ := strconv.Atoi(str[len(str)-3])
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		days[1] = 29
	}
	month, _ := strconv.Atoi(str[len(str)-2])
	day, _ := strconv.Atoi(str[len(str)-1])
	return sumOfMonth(month, days) + day
}

func sumOfMonth(month int, days [12]int) int {
	sum := 0
	// 定义了一个变量sum, return 时, 创建了一个临时变量栈用来保存返回值 为 s, 所以defer不能更改
	for i := 0; i < month-1; i++ {
		sum += days[i]
	}
	return sum
}

func main() {

}
