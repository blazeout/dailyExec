package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	// 计算偏移量, x = k % n
	// 然后从 n-x 处截断即可
	if head == nil {
		return nil
	}
	n := 1
	tail := head
	for tail.Next != nil {
		n++
		tail = tail.Next
	}
	// 获得其尾指针
	x := k % n
	if x % n == 0 {
		return head
	}
	tail.Next = head // 已经闭环
	offset := n-x-1
	cur := head
	for offset > 0 {
		offset--
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	return newHead
}


func main() {
	node2 := &ListNode{2, nil}
	node1 := &ListNode{1, node2}
	head := &ListNode{0, node1}
	head = rotateRight(head, 4)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
