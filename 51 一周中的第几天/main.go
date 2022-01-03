package main

import (
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday().String()
}

func main() {
	println(dayOfTheWeek(28, 2, 2021))
}
