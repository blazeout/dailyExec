package main

func maximumDifference(nums []int) int {
	maxRes := -1
	// 我们去固定一个nums[i]的最小值, j 往后面走, 当遇到一个小于preMin的值, 就将它更新为preMin
	for i, preMin := 1, nums[0]; i < len(nums); i++ {
		if nums[i] <= preMin {
			preMin = nums[i]
		} else {
			maxRes = max(maxRes, nums[i]-preMin)
		}
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
