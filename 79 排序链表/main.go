package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 对于链表排序需要用归并排序, 其他的排序不适用
// 这样就可以实现O(nlogn)的时间复杂度, O(1)的空间复杂度
func sortList(head *ListNode) *ListNode {
	// 1. 先分再合
	if head.Next == nil {
		// 当链表只有一个元素的时候就需要返回进行合并了
		return head
	}
	// 2. 找到链表的中间节点断开
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	nextHead := slow.Next
	slow.Next = nil
	// 经过上面这一步链表就已经从中间节点断开了
	// 然后一直断开 直到只有一个节点的时候返回
	left := sortList(head)
	right := sortList(nextHead)
	merge(left, right)
	return head
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{0, nil}
	cur := dummyHead
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
	}
	if head1 == nil {
		cur.Next = head2
	} else if head2 == nil {
		cur.Next = head1
	}
	return dummyHead.Next
}

func main() {
	node1 := &ListNode{4, nil}
	node2 := &ListNode{1, nil}
	node3 := &ListNode{2, nil}
	node4 := &ListNode{3, nil}
	node4.Next = node1
	node1.Next = node2
	node2.Next = node3
	sortList(node4)
}
