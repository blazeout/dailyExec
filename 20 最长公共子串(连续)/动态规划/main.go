package main

import "fmt"

func LongestSubString(str1, str2 string) int {
	// 动态规划, 建立一个长为n, 宽为m 的二维矩阵
	m, n := len(str1), len(str2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 双重循环, 如果str1[i] == str2[j] , 那么 dp[i][j] = 1 + dp[i-1][j-1], 不相等就等于0
	// 这样就不用计算对角线的长度了
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if str1[i] == str2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 1 + dp[i-1][j-1]
				}
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	return max
}

func main() {
	str1, str2 := "acbcbcef", "abcbced"
	fmt.Println(LongestSubString(str1, str2))
}
