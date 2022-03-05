package _35_排序链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 这是自顶向下的归并排序 空间不达标
func sortList(head *ListNode) *ListNode {
	// 归并排序 O(nlogn)
	// 如果链表只有一个元素了, 那么就return
	if head == nil || head.Next == nil {
		return head
	}
	// 然后找到链表的中间节点, 从中断开
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	zzz := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(zzz)
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	dummyHead := &ListNode{0, nil}
	cur := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return dummyHead.Next
}

// 需要自底向上的归并排序. 空间才是O(1)
