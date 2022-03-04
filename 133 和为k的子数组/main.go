package _33_和为k的子数组

func subarraySum(nums []int, k int) int {
	// 前缀和, preSum[i] - preSum[j]表示的子数组就是[j+1, i]
	preSum := 0
	// 使用map来存放前缀和出现的次数
	mp := map[int]int{0: 1}
	count := 0
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		count += mp[preSum-k]
		mp[preSum]++
	}
	return count
}
