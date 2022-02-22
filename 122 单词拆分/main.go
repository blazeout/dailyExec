package _22_单词拆分

func wordBreak(s string, wordDict []string) bool {
	// 动态规划dp[i], 代表下标以i-1结尾的字符串s是否能被拆分为子串
	dp := make([]bool, len(s)+1)
	wordMap := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = true
	}
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
