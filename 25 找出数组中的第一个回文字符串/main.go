package _5_找出数组中的第一个回文字符串

func firstPalindrome(words []string) string {
	// 找到数组中第一个回文字符串
next:
	for i := 0; i < len(words); i++ {
		// string
		b := []byte(words[i])
		length := len(b)
		for i := 0; i < length; i++ {
			if b[i] != b[length-1-i] {
				continue next
			}
		}
		return words[i]
	}
	return ""
}
