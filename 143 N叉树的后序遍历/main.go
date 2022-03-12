package main

import (
	"fmt"
)

type Node struct {
	Val      int
	Children []*Node
}

func postOrder(root *Node) []int {
	res := []int{}
	// 错了
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}
		for i := 0; i < len(root.Children); i++ {
			dfs(root.Children[i])
		}
		res = append(res, root.Val)
	}
	dfs(root)
	return res
}

func postorder(root *Node) []int {
	// BFS层序遍历的逆序
	if root == nil {
		return nil
	}
	st := []*Node{root}
	ans := []int{}
	// 左 右 根  反转一下就是 根右左, 再将根右左的顺序反转即可
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		ans = append(ans, node.Val)
		st = append(st, node.Children...)
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return ans
}

func main() {
	node5 := &Node{5, nil}
	node6 := &Node{6, nil}
	node3 := &Node{3, []*Node{node5, node6}}
	node2 := &Node{2, nil}
	node4 := &Node{4, nil}
	node1 := &Node{1, []*Node{node3, node2, node4}}
	ret := postorder(node1)
	fmt.Println(ret)
}
