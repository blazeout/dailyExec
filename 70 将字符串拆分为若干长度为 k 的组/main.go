package main

import "fmt"

func divideString(s string, k int, fill byte) []string {
	res := make([]string, 0)
	length := len(s)
	remain := length % k
	for remain > 0 {
		s += string(fill)
		remain--
	}
	for i := 0; i < len(s); i += k {
		res = append(res, s[i:i+k])
	}
	return res
}

func main() {
	res := divideString("abcdefghij", 3, 'x')
	fmt.Println(res)
}
