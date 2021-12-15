package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	sb := strings.Builder{}
	sb.WriteString("1")
	if needle == "" {
		return 0
	}
	m, n := len(haystack), len(needle)
	for i := 0; i < m; i++ {
		if haystack[i] == needle[0] {
			j := 0
			x := i
			for j < n && x < m &&haystack[x] == needle[j] {
				j++
				x++
			}
			if j == n {
				// 返回第一个出现的下标
				return i
			}
		}
	}
	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}
