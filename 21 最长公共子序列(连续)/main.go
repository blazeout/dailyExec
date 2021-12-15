package main

import (
	"fmt"
	"strings"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	// 动态规划 dp[i][j]代表以text1[0:i-1]和text2[0:j-1]的最长公共子序列长度
	// 当 text1[i] == text2[j]时 dp[i][j] = dp[i-1][j-1] + 1
	// 因为要同时减去长度1所以是dp[i-1][j-1]
	// 当text1[i] != text2[j]时, dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	// 因为当此时的最后一位不相等的时候, 就要在dp[i-1][j]和dp[i][j-1]中找一个大值
	// 比如ace和bc而言, e和c不相等, 那么此时dp[i][j]代表前i和j中的最长子序列最大值为
	// ace和b的长度为0 中 和ac和bc的长度为1
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	sb := &strings.Builder{}
	// 边界条件 dp[0][x]和dp[x][0] 有一个0代表的就是空串和text[0:x-1]的最长公共子序列长度自然就是0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if text1[i-1] == text2[j-1] {
				// 这里为什么是 i-1 和 j-1 因为i-1和j-1才代表字符串中的真实位置
				dp[i][j] = dp[i-1][j-1] + 1
				sb.WriteByte(text1[i-1])
			} else if text1[i-1] != text2[j-1] {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(sb.String())
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	str1, str2 := "abcde", "ace"
	fmt.Println(longestCommonSubsequence(str1, str2))
}
