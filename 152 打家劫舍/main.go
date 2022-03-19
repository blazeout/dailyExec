package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	// 偷今天和不偷今天的情况
	// dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	// dp[1] 也许就是我可以选择拿第一天大的, 不拿今天的
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
}
