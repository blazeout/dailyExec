package _2_链表随机节点

import "math/rand"

// 水塘抽样法
// 从链表头开始，遍历整个链表，对遍历到的第 i 个节点，随机选择区间 [0,i) 内的一个整数，如果其等于0，则将答案置为该节点值，否则答案不变。

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head,
	}
}

func (this *Solution) GetRandom() int {
	ans := 0
	for i, cur := 1, this.head; cur != nil; i++ {
		if rand.Intn(i) == 0 {
			ans = cur.Val
		}
		cur = cur.Next
	}
	return ans
}
