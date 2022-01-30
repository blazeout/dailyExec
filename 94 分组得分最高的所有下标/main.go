package main

import (
	"fmt"
	"math"
)

func maxScoreIndices(nums []int) []int {
	// 滑动窗口
	// 初始大小为0
	res := []int{}
	arr := []int{}
	maxSum := math.MinInt32
	sumL, sumR, length := 0, num(nums, 1), len(nums)
	for i := 0; i < length; i++ {
		arr = append(arr, sumL+sumR)
		maxSum = max(maxSum, sumL+sumR)
		if nums[i] == 1 {
			sumR--
		} else {
			// == 0
			sumL++
		}
	}
	arr = append(arr, sumL+sumR)
	maxSum = max(maxSum, sumL+sumR)
	for i, v := range arr {
		if maxSum == v {
			res = append(res, i)
		}
	}
	return res
}

func num(arr []int, target int) (res int) {
	for _, v := range arr {
		if target == v {
			res++
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	res := maxScoreIndices([]int{0, 0, 1, 0})
	fmt.Println(res)
}
