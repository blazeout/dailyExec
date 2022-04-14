package _83_最富有客户的资产总量

func maximumWealth(accounts [][]int) int {
	maxRes := 0
	for i := 0; i < len(accounts); i++ {
		temp := 0
		for j := 0; j < len(accounts[i]); j++ {
			temp += accounts[i][j]
		}
		if temp > maxRes {
			maxRes = temp
		}
	}
	return maxRes
}
