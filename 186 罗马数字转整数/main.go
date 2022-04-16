package _86_罗马数字转整数

// @lc code=start

// 罗马数字转整数
func romanToInt(s string) int {
	result := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && m[s[i]] < m[s[i+1]] {
			result -= m[s[i]]
		} else {
			result += m[s[i]]
		}
	}
	return result
}

// @lc code=end
