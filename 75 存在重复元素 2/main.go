package main

func containsNearbyDuplicate(nums []int, k int) bool {
	// map用来记录上一次此value出现的下标位置, 如果 i - j <= k 那么就返回true
	mp := map[int]int{}
	for index, value := range nums {
		if previousIndex, ok := mp[value]; ok {
			if abs(index-previousIndex) <= k {
				return true
			}
		}
		// 无论index-previousIndex > k 都要更新mp[value] = 最新的index
		mp[value] = index
	}
	return false
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	ret := containsNearbyDuplicate([]int{1, 0, 1, 1}, 1)
	println(ret)
}
