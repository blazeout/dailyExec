package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	distance := math.MaxInt32
	res := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 首先固定一个位置, 然后保证这次枚举的和上次不一样即可
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			tempSum := nums[i] + nums[left] + nums[right]
			if tempSum == target {
				return target
			}
			if abs(tempSum-target) < distance {
				res = tempSum
				distance = abs(tempSum - target)
			}
			if tempSum < target {
				// 需要移动到下一个不相等的地方
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				// 需要移动到下一下不相等的地方
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(A int) int {
	if A < 0 {
		return -A
	}
	return A
}

func main() {
	threeSumClosest([]int{-1, 2, 1, -4}, 1)
}
