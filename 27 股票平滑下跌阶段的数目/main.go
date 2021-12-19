package main

import "fmt"

func getDescentPeriods(prices []int) int64 {
	count := int64(len(prices))
	left, right := 0, 0
	length := len(prices)
	for left < length {
		// left 和 right 双指针 遍历一遍, 即可
		// 找到一个区间, 严格符合单调递减为1
		// 然后算出这个区间的长度, 按照公式来计算
		falg := false
		for right < length-1 && prices[right]-prices[right+1] == 1 {
			falg = true
			right++
		}
		if falg {
			length1 := right - left + 1
			count += int64((length1 * (length1 - 1)) / 2)
		}

		left = right
		right = right + 1
	}
	return count
}

// 动态规划
func getDescentPeriods2(prices []int) int64 {
	dp := make([]int, len(prices))
	for i := range dp {
		dp[i] = 1
	}
	ans := 1
	for i := 1; i < len(prices); i++ {
		if prices[i-1]-prices[i] == 1 {
			dp[i] = dp[i-1] + 1
		}
		ans += dp[i]
	}
	return int64(ans)
}

func main() {
	fmt.Println("-----------")
	println(getDescentPeriods([]int{3, 2, 1, 4}))
}
