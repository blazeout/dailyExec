package main

import (
	"fmt"
	"math"
	"sort"
)

var mp map[int]bool
func maxSubsequence(nums []int, k int) []int {
	// 依次找到最大值, 如果没有正数的话, 找到负数的最小值, 同时要记录各个位置的下标
	mp = map[int]bool{}
	res := make([]int, 0)
	indexArr := make([]int, 0)
	m := k
	for k > 0 {
		ret := findMaxIndex(nums)
		indexArr = append(indexArr, ret)
		k--
	}
	// 根据下标重新排序
	sort.Ints(indexArr)
	for i := 0; i < m; i++ {
		res = append(res, nums[indexArr[i]])
	}
	return res
}

func findMaxIndex(nums []int) (index int) {
	max := math.MinInt32
	for i := 0 ; i < len(nums); i++ {
		if !mp[i] && nums[i] > max {
			index = i
			max = nums[i]
		}
	}
	mp[index] = true
	return
}

func main() {
	nums := []int{-1,-2,3,4}
	ret := maxSubsequence(nums, 3)
	fmt.Println(ret)
}
