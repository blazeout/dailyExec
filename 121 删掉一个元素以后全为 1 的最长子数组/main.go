package main

func longestSubarray(nums []int) int {
	l, r, maxRes := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			r++
		} else {
			l = r
			r = 0
		}
		maxRes = max(maxRes, l+r)
	}
	if maxRes == len(nums) {
		return maxRes - 1
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
	println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
}
