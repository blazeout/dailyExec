package _01_可以形成最大正方形的矩形数目

// 一次遍历

func countGoodRectangles(rectangles [][]int) int {
	ans := 0
	maxLen := 0
	for i := 0; i < len(rectangles); i++ {
		k := min(rectangles[i][0], rectangles[i][1])
		if k > maxLen {
			maxLen, ans = k, 1
		} else if k == maxLen {
			ans++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
