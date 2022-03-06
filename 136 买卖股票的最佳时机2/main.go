package _36_买卖股票的最佳时机2

func maxProfit(prices []int) int {
	// 贪心算法
	maxRes := 0
	// 每次我们只取上升的钱, 下降的钱我们就不要
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			maxRes += tmp
		}
	}
	return maxRes
}
