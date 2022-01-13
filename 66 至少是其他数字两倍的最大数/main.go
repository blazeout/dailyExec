package _6_至少是其他数字两倍的最大数

func dominantIndex(nums []int) int {
	// 遍历一遍找出最大值和次大值
	max, secondMax, index := -1, -1, -1
	// 一直在更新
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			secondMax = max
			max = nums[i]
			index = i
		} else if nums[i] > secondMax {
			secondMax = nums[i]
		}
	}
	if max >= 2*secondMax {
		return index
	}
	return -1
}
