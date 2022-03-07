package main

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	flag := false
	if num < 0 {
		flag = true
		num *= -1
	}
	str := ""
	for num > 0 {
		str = strconv.Itoa(num%7) + str
		num /= 7
	}
	if flag {
		str = "-" + str
	}
	return str
}

func main() {

}
