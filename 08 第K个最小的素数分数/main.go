package main

import (
	"container/heap"
	"fmt"
)

func kthSmallestPrimeFraction(arr []int, k int) []int {
	// 优先级队列, 构建最大堆, 堆顶是最大的元素
	// 如果遇到比堆顶元素还小的元素,就放入堆中当遍历完成时,就是第K小的分数
	newArrLength := (len(arr)-1)*(len(arr))/2
	newArr := make([]pair, 0)
	// 直接在循环中利用heap即可
	for i := 0; i < len(arr)-1; i++ {
		for j := i+1; j < len(arr); j++ {
			x := float64(arr[i])/float64(arr[j])
			st := pair{arr[i], arr[j], x}
			newArr = append(newArr, st)
		}
	}
	var h hp
	heap.Init(&h)
	for i := 0; i < newArrLength; i++ {
		if h.Len() < k {
			heap.Push(&h, newArr[i])
		}else {
			// 当已经满K个了 就比较堆顶元素和当前元素的大小
			if newArr[i].value < h[0].value {
				heap.Pop(&h)
				heap.Push(&h, newArr[i])
			}
		}
	}
	return []int{h[0].x, h[0].y}
}

type pair struct {
	x int
	y int
	value float64
}

type hp []pair

func (h *hp)Pop() interface{} {
	old := *h
	n := len(old)
	value := old[n-1]
	*h = old[:n-1]
	return value
}

func (h *hp)Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h hp)Less(i, j int) bool {
	// 大堆
	return h[i].value > h[j].value
}

func (h hp)Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hp)Len() int {
	return len(h)
}

func main() {
	arr := []int{1,2,3,5}
	k := 3
	ret := kthSmallestPrimeFraction(arr, k)
	fmt.Println(ret)
}
