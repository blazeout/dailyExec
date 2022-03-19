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
	// 因为要是同一条调用链上的panic才能被defer捕捉并且调用recover方法
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err, "Recover")
			}
		}()

	}()
	panic("test")
}
