package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	insertNode := &Node{x, nil}
	// 1. 没有node 直接插入返回
	if aNode == nil {
		insertNode.Next = insertNode
		return insertNode
	}
	// 2. 有一个node,也是直接插入就行
	if aNode == aNode.Next {
		aNode.Next = insertNode
		insertNode.Next = aNode
		return aNode
	}
	// 3. 有两个node,就插在比他小的后面, 或者是整个列表的最大值或者最小值
	// 或者如果值都一样,那么就随便找个地方插入就行
	cur, next := aNode, aNode.Next
	for next != aNode {
		// 如果x的值位于cur和next两者中间那么就应该插入
		// 如果x的值等于cur和next 那么也应该插入
		if cur.Val <= x && x <= next.Val {
			break
		}
		// 如果x是整个链表的最大值或者最小值,那么 cur 和 next 值的方向就应该相反
		if cur.Val > next.Val {
			if x > cur.Val || x < next.Val {
				break
			}
		}
		cur = cur.Next
		next = next.Next
	}
	insertNode.Next = next
	cur.Next = insertNode
	return aNode
}

func main() {
	node2 := &Node{1, nil}
	node1 := &Node{4, node2}
	head := &Node{3, node1}
	node2.Next = head

	node := insert(head, 2)
	fmt.Println(node.Val)
}
