package main

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	// 对于每个房屋，要么用前面的暖气，要么用后面的，二者取近的，得到距离
	// 就是对于每一个house, 要找到离他最近的两个heaters, 取两者距离的最小值
	// 对于heaters要排序, 找到第一个heaters[i] <= house的值, 找到第一个 > house的值
	sort.Ints(heaters)
	ans := 0
	for _, house := range houses {
		// 坐标返回的是heaters里面的
		j := Bsearch(heaters, house+1)
		i := j - 1
		minDns := math.MaxInt32
		// j 是第一个大于house的加热器
		if j < len(heaters) {
			minDns = heaters[j] - house
		}
		if i >= 0 && house-heaters[i] < minDns {
			minDns = house - heaters[i]
		}
		if minDns > ans {
			ans = minDns
		}
	}
	return ans
}

func Bsearch(nums []int, x int) int {
	left, right := 0, len(nums)
	//if nums[left] > x {
	//	return -1
	//}
	// 这是找i的位置才对
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] < x {
			left = mid + 1
		} else if nums[mid] > x {
			right = mid
		} else {
			return mid
		}
	}

	return left
}

func main() {
	nums := []int{1, 5}
	println(findRadius(nums, []int{2}))
}
