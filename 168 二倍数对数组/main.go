package _68_二倍数对数组

import "sort"

func canReorderDoubled(arr []int) bool {
	cnt := map[int]int{}
	for i := 0; i < len(arr); i++ {
		cnt[arr[i]]++
	}
	if cnt[0]%2 == 1 {
		return false
	}
	val := make([]int, 0)
	for key, _ := range cnt {
		val = append(val, key)
	}
	sort.Slice(val, func(i, j int) bool { return abs(val[i]) < abs(val[j]) })
	for _, value := range val {
		if cnt[value*2] < cnt[value] {
			return false
		}
		cnt[2*value] -= cnt[value]

	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
