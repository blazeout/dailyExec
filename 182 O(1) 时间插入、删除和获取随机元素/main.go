package main

import (
	"fmt"
	"math/rand"
)

// RandomizedSet 数组存储具体值, 实现O(1)插入和查找
// 哈希表存储具体的值的位置, 实现O(1)查找值所在的位置, 然后需要删除就将这个位置上面的值和最后一个交换删除O(1)
// 随机获取元素就random
type RandomizedSet struct {
	arr []int
	mp  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		arr: make([]int, 0),
		mp:  make(map[int]int),
	}
}

// Insert 插入到数组的最后面
func (s *RandomizedSet) Insert(val int) bool {
	// 存在即不插入元素
	if _, ok := s.mp[val]; !ok {
		// 插入
		s.arr = append(s.arr, val)
		s.mp[val] = len(s.arr) - 1
		return true
	}
	//  else 直接return false
	return false
}

func (s *RandomizedSet) Remove(val int) bool {
	// remove之后还需要更新mp
	if _, ok := s.mp[val]; ok {
		// 删除
		last := len(s.arr) - 1
		// 更新mp, 将原来的最后一个元素放到要删除的位置, 所以我们需要更新mp
		s.mp[s.arr[last]] = s.mp[val]
		// 将最后一个元素和要删除的元素交换
		s.arr[s.mp[val]], s.arr[last] = s.arr[last], s.arr[s.mp[val]]
		// 删除最后一个元素
		s.arr = s.arr[:last]
		// 删除mp中的元素
		delete(s.mp, val)
		return true
	}
	return false
}

func (s *RandomizedSet) GetRandom() int {
	random := rand.Intn(len(s.arr))
	return s.arr[random]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	obj := Constructor()
	obj.Insert(0)
	obj.Insert(1)
	obj.Remove(0)
	obj.Insert(2)
	obj.Remove(1)
	ret := obj.GetRandom()
	fmt.Println(ret)
}
