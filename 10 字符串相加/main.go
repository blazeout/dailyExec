package main

import (
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	strconv.Itoa()
	sb  := &strings.Builder{}
	carry := 0 // carry 表示进位
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1{
		var x, y int
		if i >= 0 {
			x = int(num1[i]-'0')
		}
		if j >= 0 {
			y = int(num2[j]-'0')
		}
		result := x+y+carry
		sb.WriteString(strconv.Itoa(result%10))
		carry = result/10
	}
	ans := sb.String()
	return reverser(ans)
}

func reverser(s string) string {
	var a []byte
	for _, k := range []byte(s) {
		defer func(v byte) {
			a = append(a, v)
		}(k)
	}
	return string(a)
}

func main() {

}
