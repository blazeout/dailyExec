package _7_重新格式化字符串

func reformat(s string) string {
	charCount, numCount := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			charCount++
		} else {
			numCount++
		}
	}
	if charCount-numCount > 1 || charCount-numCount < -1 {
		return ""
	}
	// i 用来记录字符的位置, j用来记录数字的位置
	i, j := 0, 0
	if charCount >= numCount {
		i, j = 0, 1
	} else {
		i, j = 1, 0
	}
	arr := make([]byte, len(s))
	for z := 0; z < len(s); z++ {
		if s[z] >= 'a' && s[z] <= 'z' {
			arr[i] = s[z]
			i += 2
		} else {
			arr[j] = s[z]
			j += 2
		}
	}
	return string(arr)
}
