package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	// 二叉搜索树中序遍历是有序的
	res := 0
	var dfs func(root *TreeNode, target *int)
	dfs = func(root *TreeNode, target *int) {
		if root == nil || *target < 0 {
			return
		}
		dfs(root.Left, target)
		*target--
		if *target == 0 {
			res = root.Val
			return
		}
		dfs(root.Right, target)
		// 左根右
	}
	dfs(root, &k)
	return res
}

func main() {
	root2 := &TreeNode{2, nil, nil}
	root1 := &TreeNode{1, nil, root2}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root1, root4}
	ret := kthSmallest(root3, 1)
	fmt.Println(ret)
}
