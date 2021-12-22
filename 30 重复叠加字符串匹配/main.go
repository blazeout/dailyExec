package main

import (
	"math"
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	// 先统计，a和b中出现的字母， 如果两个数组都不一样的， 那么就直接返回不需要匹配
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := 0; i < len(a); i++ {
		cnt1[a[i]-'a'] = 1
	}
	for i := 0; i < len(b); i++ {
		cnt2[b[i]-'a'] = 1
	}
	for i := 0; i < 26; i++ {
		if cnt2[i] != 0 && cnt2[i] != cnt1[i] {
			return -1
		}
	}
	before := a
	var num float64
	num = float64(len(b)) / float64(len(a))
	l := (int)(math.Ceil(num))
	a = strings.Repeat(before, l)
	// 接下来就是主干部分
	// 首先利用strings.Contians()方法, 判断a中是否包含了b这个字串， 如果不包含那么就继续for下去
	if strings.Contains(a, b) {
		return l
	}
	a += before
	if strings.Contains(a, b) {
		return l + 1
	}
	return -1
}

func main() {
	str1 := "abcd"
	str2 := "bc"
	println(repeatedStringMatch(str1, str2))
}
