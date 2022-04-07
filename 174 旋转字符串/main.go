package main

import "strings"

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	for i := 0; i < len(s); i++ {
		s = s[1:] + string(s[0])
		if s == goal {
			return true
		}
	}
	return false
}

// 字符串旋转判断是否相等 可以用两个A拼接起来, 判断B是否是A的子串即可

func rotateString2(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}

func main() {

}
