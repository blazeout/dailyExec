package _8_执行所有的后缀指令

// 蔚来周赛第二题

func executeInstructions(n int, startPos []int, s string) []int {
	m := len(s)
	res := make([]int, m)
	for i := 0; i < m; i++ {
		r := startPos[0]
		c := startPos[1]
		for _, ch := range s[i:] {
			switch ch {
			case 'U':
				r--
			case 'D':
				r++
			case 'L':
				c--
			case 'R':
				c++
			}
			if r < 0 || r >= n || c < 0 || c >= n {
				break
			}
			res[i]++
		}
	}
	return res
}
