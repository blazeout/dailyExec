package _58_如果相邻两个颜色均相同则删除当前颜色

func winnerOfGame(colors string) bool {
	// 统计colors中可删除的A的数量, 和可删除的B的数量
	get := func(i int, char byte) bool {
		if i < 0 || i >= len(colors) {
			return false
		}
		if colors[i] != char {
			return false
		}
		return true
	}
	countA, countB := 0, 0
	for i := 1; i < len(colors)-1; i++ {
		if colors[i] == 'A' {
			if get(i-1, 'A') && get(i+1, 'A') {
				countA++
			}
		}
		if colors[i] == 'B' {
			if get(i-1, 'B') && get(i+1, 'B') {
				countB++
			}
		}
	}
	return countA > countB
}
