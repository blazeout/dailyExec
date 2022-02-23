package _23_仅仅反转字母

func reverseOnlyLetters(s string) string {
	// 双指针
	i, j := 0, len(s)-1
	b := []byte(s)
	for i < j {
		if isDigit(s[i]) && isDigit(s[j]) {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		} else if isDigit(s[i]) && !isDigit(s[j]) {
			// j 位置不是字母
			j--
		} else if !isDigit(s[i]) && isDigit(s[j]) {
			// i 位置不是字母
			i++
		} else {
			i++
			j--
		}
	}
	return string(b)
}

func isDigit(x byte) bool {
	if x <= 'Z' && x >= 'A' || x <= 'z' && x >= 'a' {
		return true
	}
	return false
}
