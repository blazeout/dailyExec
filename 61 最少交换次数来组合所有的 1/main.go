package main

import (
	"fmt"
	"math"
)

func minSwaps(nums []int) (minRes int) {
	// 数出一共有多少个1
	// 制造一个为N大小的滑动窗口
	// 数出滑动窗口中0的个数的最小值
	countOne := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			countOne++
		}
	}
	minRes = math.MaxInt32
	countZero := 0
	// 因为是首尾相连, 所以我们要断环成链
	for i := 0; i < 2*len(nums); i++ {
		if i >= countOne {
			minRes = min(minRes, countZero)
			// 当我们右移时, 如果是0移出去了那么就countZero--
			if nums[(i-countOne)%len(nums)] == 0 {
				countZero--
			}
		}
		if nums[i%len(nums)] == 0 {
			countZero++
		}
	}
	return minRes
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{0, 1, 0, 1, 1, 0, 0}
	ret := minSwaps(nums)
	fmt.Println(ret)
}
