package _9_零钱兑换2_完全背包求解组合数

func change(amount int, coins []int) int {
	// dp[j]代表可以组成的种数
	dp := make([]int, amount+1)
	dp[0] = 1
	// 求组合数, 外层for遍历物品, 内层for遍历容量
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}
