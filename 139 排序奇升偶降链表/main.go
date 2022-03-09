package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param head ListNode类
 * @return ListNode类
 */
// 拆分为两个链表然后在逆序然后再归并排即可

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func sortLinkedList(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return nil
	}
	O := head
	G := head.Next
	GHead := G
	for G != nil && G.Next != nil {
		O.Next = G.Next
		O = O.Next
		G.Next = O.Next
		G = G.Next
	}
	O.Next = nil
	newHead := reverse(GHead)
	dummyHead := &ListNode{0, nil}
	cur := dummyHead
	for head != nil && newHead != nil {
		if head.Val <= newHead.Val {
			cur.Next = head
			head = head.Next
		} else {
			cur.Next = newHead
			newHead = newHead.Next
		}
		cur = cur.Next
	}
	if head == nil {
		cur.Next = newHead
	} else {
		cur.Next = head
	}
	return dummyHead.Next
}

func main() {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{3, node1}
	node3 := &ListNode{2, node2}
	node4 := &ListNode{2, node3}
	node5 := &ListNode{3, node4}
	node6 := &ListNode{1, node5}
	ret := sortLinkedList(node6)
	fmt.Println(ret)
}
