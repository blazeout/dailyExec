package main

import (
	"fmt"
	"math"
	"math/rand"
)

func minimumDifference(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	ans1 := math.MaxInt32
	for i := 0; i < len(nums)-k+1; i++ {
		ans1 = min(ans1, nums[i+k-1]-nums[i])
	}
	return ans1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func quickSort(nums []int, left, right int) {
	if left < right {
		mid := patition(nums, left, right)
		quickSort(nums, left, mid-1)
		quickSort(nums, mid+1, right)
	}
}

func patition(nums []int, left, right int) int {
	random := rand.Intn(right-left+1) + left
	nums[random], nums[right] = nums[right], nums[random]
	l, r := left-1, left
	pivot := nums[right]
	for ; r < right; r++ {
		if nums[r] <= pivot {
			l++
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[l+1], nums[right] = nums[right], nums[l+1]
	return l + 1
}

func main() {
	nums := []int{9, 4, 1, 7}
	quickSort(nums, 0, 3)
	fmt.Println(nums)
}
