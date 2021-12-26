package _4_字符串数组中的最大长度

import "strings"

// 统计空格个数 + 1 即可
func mostWordsFound(sentences []string) int {
	ans := 0
	for _, s := range sentences {
		cnt := strings.Count(s, " ") + 1
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}
