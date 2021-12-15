package main

import (
	"fmt"
)

func mySqrt(x int) int {

	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r - l) / 2
		if mid * mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}


type TopVotedCandidate struct {
	// 记录次数
	mp map[int]int
	// 在每一个时刻就记录
	everyTime []int
	// 最近获得选举的人
	latest []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	top := TopVotedCandidate{
		mp:        make(map[int]int),
		everyTime: make([]int, times[len(times)-1]+1),
		latest: make([]int, len(persons)),
	}
	for i := 0; i < len(persons); i++ {
		// 先更新
		if top.mp[persons[i]] > 0 {
			top.mp[persons[i]] += 1
		} else {
			top.mp[persons[i]] = 1
		}
		top.latest[i] = persons[i]
		// 还需要实现一个票数相同时最近获得选举的人获胜的
		if i < len(persons)-1 {
			for j := times[i]; j < times[i+1]; j++ {
				max := -1
				maxPeople := -1
				for key, value := range top.mp {
					if value >= max {
						max = value
						// 实现最近获得选举的人
						if top.latest[i] == key {
							maxPeople = key
						}else {
							maxPeople = top.latest[i]
						}
					}
				}
				top.everyTime[j] = maxPeople
			}
		} else if i == len(persons)-1 {
			max := -1
			maxPeople := -1
			for key, value := range top.mp {
				if value > max {
					max = value
					maxPeople = key
				}
			}
			top.everyTime[times[len(times)-1]] = maxPeople
		}
	}
	return top
}

func (this *TopVotedCandidate) Q(t int) int {
	return this.everyTime[t]
}

func main() {
	person := []int{0, 1, 1, 0, 0, 1, 0}
	times := []int{0, 5, 10, 15, 20, 25, 30}
	top := Constructor(person, times)
	fmt.Println(top.Q(24))
}
