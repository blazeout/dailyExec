package _8_检查所有A是否在B之前

import "strings"

func checkString(s string) bool {
	// 只要字符串不包含ba即可
	return !strings.Contains(s, "ba")
}
