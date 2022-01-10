package _0_判断是否为有效矩阵

func checkValid(matrix [][]int) bool {
	n := len(matrix)
	sum := 1
	for i := 1; i <= n; i++ {
		sum ^= i
	}
	// 横向遍历
	for i := 0; i < n; i++ {
		curSum := 1
		for j := 0; j < n; j++ {
			curSum ^= matrix[i][j]
		}
		if curSum != sum {
			return false
		}
	}
	// 纵向遍历
	for j := 0; j < n; j++ {
		curSum := 1
		for i := 0; i < n; i++ {
			curSum ^= matrix[i][j]
		}
		if curSum != sum {
			return false
		}
	}
	return true
}
