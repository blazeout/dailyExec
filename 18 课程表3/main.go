package _8_课程表3

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	// 记录一个count 代表现在已经修了多久的课
	// 每次根据这个count + course[i][0] <= course[i][1] , 那么就可以修够
	// 现在主要的就是确定按照什么来排序, 修学分
	// 先按照course[i][1]来排序, lastDay大的放在后面
	sort.Slice(courses, func(i, j int) bool { return courses[i][1] < courses[j][1] })
	count := 0
	h := &hp{}
	for _, course := range courses {
		if d := course[0]; d+count <= course[1] {
			count += d
			heap.Push(h, d)
		} else if h.Len() > 0 && h.IntSlice[0] > d {
			count += d - h.IntSlice[0]
			h.IntSlice[0] = d
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Pop() interface{} {
	length := len(h.IntSlice)
	old := h.IntSlice
	value := old[length-1]
	h.IntSlice = old[:length-1]
	return value
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

// Less 大跟堆
func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}
