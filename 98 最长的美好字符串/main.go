package main

import "fmt"

func longestNiceSubstring(s string) string {
	// 用两个数组来记录字符出现的次数
	cnt1, cnt2 := [26]int{}, [26]int{}
	maxRes := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			char := s[j]
			if char >= 'A' && char <= 'Z' {
				cnt2[char-'A'] = 1
			} else {
				cnt1[char-'a'] = 1
			}
			if cnt1 == cnt2 {
				maxRes = max(s[i:j+1], maxRes)
			}
		}
		cnt1, cnt2 = [26]int{}, [26]int{}
	}
	return maxRes
}

func max(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}

func main() {
	ress := longestNiceSubstring("YazaAay")
	fmt.Println(ress)
}
