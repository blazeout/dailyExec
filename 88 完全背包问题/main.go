package main

import "math"

func coinChange(coins []int, amount int) int {
	// dp[j]代表总额为j的需要的最小硬币数
	dp := make([]int, amount+1)
	// 总额为0的话金币数就是0, dp[0] = 0
	// dp[j]只与dp[j-coins[i]]有关, 因为只能由coins中的硬币组成, 那些不能组成的dp[j] == math.MaxInt32
	// 如果dp[j] == math.MaxInt32 return -1;
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i < len(coins); i++ {
		// 外层循环遍历物品, 即coins[i]
		for j := coins[i]; j <= amount; j++ {
			// 内层循环遍历背包
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}

	// 现在是外层循环遍历背包, 内层循环遍历物品
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
