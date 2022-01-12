package _5_递增的三元子序列

import "math"

func increasingTriplet(nums []int) bool {
	// 贪心
	// 1. 记录最小的和第二小的
	// 2. 遇到这个数, 有三种情况
	//    2.1 比最小的小, 更新最小值为此值, 更新第二的值为math.MaxINT32
	//    2.2 如果遇到比第二个大的值, 直接返回true
	//    2.3 如果遇到在最小的第二个的中间就更新第二个的值为此值
	//    2.4 我们可能在第一个的 first 和 second 后面又找到一个first2小于first, 此时更新first
	//    2.5 但此时我们只要找到一个num > second 就可以返回, 因为在之前是有一个first < second的
	//    2.6 我们此时更新first是为了能够更大可能找到递增的三元子序列

	if len(nums) < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32
	for _, v := range nums {
		if v <= first {
			first = v
		} else if v <= second {
			second = v
		} else {
			return true
		}
	}
	return false
}
