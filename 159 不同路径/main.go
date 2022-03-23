package _59_不同路径

func uniquePaths(m int, n int) int {
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				matrix[i][j] = 1
			} else if j == 0 {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
			}
		}
	}
	return matrix[m-1][n-1]
}
