package _04_在每个树行中找最大值

import "math"

// 树节点的结构体

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValuesBFS(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	ans := []int{}
	for len(queue) > 0 {
		length := len(queue)
		max := math.MinInt32
		for length > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			length--
		}
		ans = append(ans, max)
	}
	return ans
}

func largestValuesDFS(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := []int{}
	var dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		// 如果当前层次比之前的层次大. 说明还没有被记录过, 需要append记录
		if len(ans) == level {
			ans = append(ans, root.Val)
		} else {
			ans[level] = max(ans[level], root.Val)
		}
	}
	dfs(root, 0)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
