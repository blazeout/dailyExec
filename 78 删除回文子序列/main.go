package _8_删除回文子序列

func removePalindromeSub(s string) int {
	// 脑筋急转弯, 最多删除两次, 第一次删除全部的a, 第二次删除全部的b
	// 删除1次的情况就是, 此串本身就是回文串
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return 2
		}
	}
	return 1
}
