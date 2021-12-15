package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	var dfs func(root *TreeNode)
	var res *TreeNode
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val == val {
			res = root
			return
		}
		if root.Val < val {
			dfs(root.Left)
		}else if root.Val > val {
			dfs(root.Right)
		}
	}
	dfs(root)
	return res
}

func main() {
	root1 := &TreeNode{1, nil, nil}
	root2 := &TreeNode{3, nil, nil}
	root3 := &TreeNode{2, root1, root2}
	root4 := &TreeNode{7, nil, nil}
	root5 := &TreeNode{4, root3, root4}
	ret := searchBST(root5, 2)
	fmt.Println(ret.Val)
}
