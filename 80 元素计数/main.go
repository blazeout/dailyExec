package _0_元素计数

// 统计最大值和最小值的个数 长度减去个数总和
func countElements(nums []int) int {
	min, max, cntMin, cntMax := nums[0], nums[0], 1, 1
	for _, v := range nums[1:] {
		if v < min {
			min, cntMin = v, 1
		} else if v == min {
			cntMin++
		}
		if v > max {
			max, cntMax = v, 1
		} else if v == max {
			cntMax++
		}
	}
	if min == max {
		return 0
	}
	return len(nums) - cntMin - cntMax
}
