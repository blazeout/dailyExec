package main

import (
	"container/heap"
	"fmt"
)

func minMax(a []int, k int, x int) int {
	// write code here
	// 用堆 ?
	var h IntHeap
	h = a
	heap.Init(&h)
	for i := 0; i < k; i++ {
		ret := heap.Pop(&h).(int)
		ret -= x
		heap.Push(&h, ret)
	}
	return h[0]
}

type IntHeap []int

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func main() {
	ret := minMax([]int{7, 2, 1}, 3, 2)
	fmt.Println(ret)
}
