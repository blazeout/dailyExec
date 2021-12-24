package _2_吃苹果的最大数目

import "container/heap"

func eatenApples(apples []int, days []int) int {
	// 策略贪心 : 每次都吃一个快要过期的苹果
	// 这个快要过期的苹果可以用最小堆来实现
	i := 0
	ans := 0
	h := &hp{}
	heap.Init(h)
	for i < len(apples) || h.Len() != 0 {
		for h.Len() > 0 && (*h)[0].day <= i {
			heap.Pop(h)
		}
		if i < len(apples) && apples[i] > 0 {
			heap.Push(h, pair{apples[i], i + days[i]})
		}
		// 每次循环前要将hp[0].apple--和day--, 然后判断是否day <= 0
		if h.Len() > 0 {
			ans++
			(*h)[0].apple--
			if (*h)[0].apple == 0 {
				heap.Pop(h)
			}
		}
		i++
	}
	return ans
}

type pair struct {
	apple int
	day   int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}
func (h hp) Less(i, j int) bool {
	return h[i].day < h[j].day
}
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(pair))
}
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
