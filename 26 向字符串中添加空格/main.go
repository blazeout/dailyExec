package _6_向字符串中添加空格

import "strings"

func addSpaces(s string, spaces []int) string {
	sb := strings.Builder{}
	begin := 0
	length := len(spaces)
next:
	for j, index := range spaces {
		for i := begin; i < len(s); i++ {
			if i == index {
				sb.WriteByte(' ')
				begin = i
				if j == length-1 {

				} else {
					continue next
				}
			}
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}
