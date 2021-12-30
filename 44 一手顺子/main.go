package _4_一手顺子

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	mp := map[int]int{}
	for _, v := range hand {
		mp[v]++
	}
	for _, v := range hand {
		if mp[v] == 0 {
			continue
		}
		// 这里就代表mp[v] > 0, 可以进行下面的步骤
		for num := v; num < v+groupSize; num++ {
			if mp[num] == 0 {
				return false
			}
			mp[num]--
		}
	}
	return true
}
