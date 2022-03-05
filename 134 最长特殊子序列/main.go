package _34_最长特殊子序列

// 如果不相等, 那么就返回那个最长的即可
// 如果相等, 直接返回 -1

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
