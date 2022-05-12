package _98__删列造序

// 删列造序
func minDeletionSize(strs []string) int {
	count := 0
next:
	for j := 0; j < len(strs[0]); j++ {
		before := strs[0][j]
		for i := 1; i < len(strs); i++ {
			if before > strs[i][j] {
				count++
				continue next
			} else {
				before = strs[i][j]
			}
		}
	}
	return count
}
