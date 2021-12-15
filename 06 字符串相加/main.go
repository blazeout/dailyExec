package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	ans := ""
	carry := 0
	// carry 存储进位
	for i, j := len(num1)-1, len(num2)-1; i>=0 || j>=0 || carry != 0; i,j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			// byte - byte 还是 byte
			x = int(num1[i]-'0')
		}
		if j >= 0 {
			y = int(num2[j]-'0')
		}
		result := x+y+carry
		ans = strconv.Itoa(result%10) + ans
		carry = result/10
	}
	return ans
}

func main() {
	fmt.Println('1'-'0')
}
