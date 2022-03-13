package main

import (
	"fmt"
)

func findKDistantIndices(nums []int, key int, k int) []int {
	mp := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			mp[key] = append(mp[key], i)
		}
	}
	res := []int{}
	exist := map[int]bool{}
	m := len(nums)
	for _, v := range mp[key] {
		// i是index, v是value
		// 在v的左右两边扩散k的单位
		for i := v - k; i <= v+k; i++ {
			if _, ok := exist[i]; !ok && i < m && i >= 0 {
				res = append(res, i)
				exist[i] = true
			}
		}
	}
	if len(res) < m {
		return res
	}
	res2 := []int{}
	for i := 0; i < len(res); i++ {
		if res[i] < 0 {
			continue
		} else {
			res2 = append(res2, res[i])
			m--
			if m == 0 {
				break
			}
		}
	}
	return res2
}

func main() {
	ret := findKDistantIndices([]int{2, 1, 2, 3, 4, 2, 6, 7, 2, 9, 2}, 9, 5)
	fmt.Println(ret)
}
