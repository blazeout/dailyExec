package main

import "fmt"

func reverseWords(s string) string {
	left := 0
	b := []byte(s)
	for right := 0; right < len(b); right++ {
		if b[right] == ' ' {
			for i, j := left, right-1; i < j; i, j = i+1, j-1 {
				b[i], b[j] = b[j], b[i]
			}
			left = right + 1
		}
		if right == len(b)-1 {
			for i, j := left, right; i < j; i, j = i+1, j-1 {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
	return string(b)
}

func main() {
	ret := reverseWords("Let's take LeetCode contest")
	fmt.Println(ret)
}
