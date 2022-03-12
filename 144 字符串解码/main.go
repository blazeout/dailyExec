package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	// 栈 ?
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			// 把其他字符全部压入栈中
			stack = append(stack, s[i])
		} else {
			b1 := make([]byte, 0)
			// 取出所有字符
			for len(stack) != 0 && unicode.IsLetter(rune(stack[len(stack)-1])) {
				b1 = append(b1, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// 吐出 '['
			stack = stack[:len(stack)-1]
			reverse(b1)
			tempStr := string(b1)
			// 现在是数字
			b2 := make([]byte, 0)
			for len(stack) != 0 && unicode.IsDigit(rune(stack[len(stack)-1])) {
				b2 = append(b2, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			reverse(b2)
			count, _ := strconv.Atoi(string(b2))
			str := strings.Repeat(tempStr, count)
			// 然后再将这些东西压入栈中
			b3 := []byte(str)
			for i := 0; i < len(b3); i++ {
				stack = append(stack, b3[i])
			}
		}
	}
	// 最后就是拼接啦
	return string(stack)
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < len(b)/2; i++ {
		b[i], b[j-i] = b[j-i], b[i]
	}
}

func main() {
	ret := decodeString("3[a]2[bc]")
	fmt.Println(ret)
}
