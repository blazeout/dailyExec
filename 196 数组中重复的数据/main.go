package _96_数组中重复的数据

func findDuplicates(nums []int) []int {
	// 原地哈希
	// 已nums[i]-1为下标标记为负数
	res := []int{}
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] < 0 {
			res = append(res, abs(nums[i]))
		}
		nums[index] *= -1
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
