package main

import "math/rand"

type RandomizedSet struct {
	arr []int
	mp  map[int]int // key为val, value为数组下标索引
}

// 数组存储值（实现O(1)插入、随机读)
// 哈希表存储值对应的索引（实现O(1)查找）
// 删除时将要删除的数与数组最后一个数交换位置（避免移动。实现O(1)删除）

func Constructor() RandomizedSet {
	return RandomizedSet{
		arr: []int{},
		mp:  map[int]int{},
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.mp[val]; !ok {
		// 不存在那么就加入
		s.arr = append(s.arr, val)
		s.mp[val] = len(s.arr) - 1
		return true
	}
	return false
}

func (s *RandomizedSet) Remove(val int) bool {
	//s.v2i[s.arr[index]] = index
	//delete(s.v2i, s.arr[lastOne])
	if _, ok := s.mp[val]; ok {
		index := s.mp[val]
		s.arr[index], s.arr[len(s.arr)-1] = s.arr[len(s.arr)-1], s.arr[index]
		s.mp[s.arr[index]] = index
		s.arr = s.arr[:len(s.arr)-1]
		delete(s.mp, val)
		return true
	}
	return false
}

func (s *RandomizedSet) GetRandom() int {
	random := rand.Intn(len(s.arr))
	return s.arr[random]
}

func main() {

}
