package main

import "sort"

func minimumRemoval(beans []int) int64 {
	// 排序后, 遍历
	// 将小于v的魔法豆清空, 然后算剩余的豆子
	sort.Ints(beans)
	mx := 0
	sum := 0
	for i := 0; i < len(beans); i++ {
		sum += beans[i]
		mx = max(mx, (len(beans)-i)*beans[i])
	}
	return int64(sum - mx)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
