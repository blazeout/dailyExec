package _81_写字符串需要的行数

func numberOfLines(widths []int, s string) []int {
	const maxWidths = 100
	line, count := 1, 0
	for i := 0; i < len(s); i++ {
		count += widths[s[i]-'a']
		need := widths[s[i]-'a']
		if count > maxWidths {
			line++
			count = need
		}
	}

	return []int{line, count}
}
