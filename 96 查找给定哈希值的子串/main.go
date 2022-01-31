package main

import (
	"fmt"
	"math"
)

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {

	length := len(s)
	for i := 0; i < length; i++ {
		if i >= 16 {
			return s[:i]
		}
		if i+k <= length && hashFunc(s[i:i+k], power, modulo) == hashValue {
			return s[i : i+k]
		}
	}
	return s
}

// 22, 51, 41, 9
func hashFunc(s string, p, m int) int {
	var sum int64
	for i := 0; i < len(s); i++ {
		d := int64(s[i]-'a') + int64(1)
		sum += d * int64(math.Pow(float64(p), float64(i))) % int64(m)
	}
	return int(sum % int64(m))
}
func main() {
	res := subStrHash("xqgfatvtlwnnkxipmipcpqwbxihxblaplpfckvxtihonijhtezdnkjmmk", 22, 51, 41, 9)
	fmt.Println(res)
}
