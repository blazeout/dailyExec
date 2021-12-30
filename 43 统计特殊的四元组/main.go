package main

import "fmt"

func countQuadruplets(nums []int) int {
	// 用哈希表存储 nums[d] 的 下标
	ans := 0
	mp := map[int]int{}
	for i := 3; i < len(nums); i++ {
		mp[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if x, ok := mp[nums[i]+nums[j]+nums[k]]; ok && x > 0 {
					ans += x
				}
			}
		}
	}
	return ans
}

func countQuadruplets2(nums []int) (ans int) {
	cnt := map[int]int{}
	// nums[a]+nums[b]=nums[d]−nums[c]
	for b := len(nums) - 3; b >= 1; b-- {
		for _, x := range nums[b+2:] {
			cnt[x-nums[b+1]]++
		}
		for _, x := range nums[:b] {
			ans += cnt[x+nums[b]]
		}
	}
	return
}

func main() {
	nums := []int{1, 2, 3, 6}
	fmt.Println("==============")
	println(countQuadruplets2(nums))
}
