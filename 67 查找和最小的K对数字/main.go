package _7_查找和最小的K对数字

import "container/heap"

// 自己写的

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// 肯定不需要计算出全部的序列对, 使用最小堆 ?
	res := make([][]int, k)
	h := &hp{}
	minn1 := min(len(nums1), 60)
	minn2 := min(len(nums2), 60)
	for i := 0; i < minn1; i++ {
		for j := 0; j < minn2; j++ {
			tmp := []int{nums1[i], nums2[j]}
			if h.Len() < k {
				heap.Push(h, tmp)
			} else {
				if (tmp[0] + tmp[1]) < ((*h)[0][0] + (*h)[0][1]) {
					// 最大堆 比最大值小就需要加入堆中
					heap.Pop(h)
					heap.Push(h, tmp)
				}
			}
		}
	}
	for i := k - 1; i >= 0; i-- {
		if h.Len() > 0 {
			res[i] = heap.Pop(h).([]int)
		} else {
			// 就需要把前面的剪掉
			res = res[1:]
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type hp [][]int

func (h hp) Len() int {
	return len(h)
}
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h hp) Less(i, j int) bool {
	return (h[i][0] + h[i][1]) > (h[j][0] + h[j][1])
}
func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.([]int))
}

// 官方解法
