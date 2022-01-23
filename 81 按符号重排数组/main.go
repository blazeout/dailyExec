package main

import "fmt"

func rearrangeArray(nums []int) []int {
	arr := make([]int, 0)
	mp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			// 正数
			for j := 0; j < len(nums); j++ {
				if nums[j] > 0 && mp[nums[j]] > 0 {
					arr = append(arr, nums[j])
					mp[nums[j]]--
					break
				}
			}
		} else {
			// 负数
			for j := 0; j < len(nums); j++ {
				if nums[j] < 0 && mp[nums[j]] > 0 {
					arr = append(arr, nums[j])
					mp[nums[j]]--
					break
				}
			}
		}
	}
	return arr
}

func main() {
	nums := []int{28, -41, 22, -8, -37, 46, 35, -9, 18, -6, 19, -26, -37, -10, -9, 15, 14, 31}
	res := rearrangeArray(nums)
	fmt.Println(res)
}
