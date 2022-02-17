package _16_缺失的第一个正数

func firstMissingPositive(nums []int) int {
	// 原地哈希, 因为我们只关注1~N+1的正整数, 所以我们可以将数组看成一个哈希表
	// if num[i] == 1 我们就将nums[1-1]置为负数, 如果已经是负数就不用置换了
	// 对于数组中的负数我们应该如何处理呢?, 我们可以将他们置为一个N+1的数, 遍历到他就不用管了
	N := len(nums)
	for i := 0; i < N; i++ {
		if nums[i] <= 0 {
			nums[i] = N + 1
		}
	}
	for i := 0; i < N; i++ {
		num := abs(nums[i])
		if num <= N && num >= 1 {
			if nums[num-1] > 0 {
				nums[num-1] *= -1
			}
		}
	}
	for i := 0; i < N; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return N + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
