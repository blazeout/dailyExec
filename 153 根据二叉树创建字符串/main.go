package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	str := ""
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			str += "()"
			return
		}
		// 根左右
		str += "(" + strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			// 啥也不干
		} else if root.Left != nil && root.Right == nil {
			dfs(root.Left)
		} else {
			// 右孩子不为空就要遍历两个子树
			dfs(root.Left)
			dfs(root.Right)
		}
		str += ")"
	}
	dfs(root)
	str = str[1:]
	str = str[:len(str)-1]
	return str
}

// 几个月之前写的
func tree2str2(root *TreeNode) string {
	// 1. 先构递归构建输出结果
	// 2. 如果节点为空那么就返回空
	if root == nil {
		return ""
	}
	// 3. 如果节点两孩子都为空, 那么就只需要返回值不需要加括号
	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val) + ""
	}
	// 4. 如果右孩子为空, 左孩子不为空, 那么就需要为自己加上一层括号并且递归遍历他的左孩子
	if root.Right == nil {
		return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")"
	}
	// 5. 如果右孩子不为空那么需要先加一层括号表示左孩子 然后
	return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")(" + tree2str(root.Right) + ")"
}

func main() {
	node1 := &TreeNode{1, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node4 := &TreeNode{4, nil, nil}
	node1.Left = node2
	node1.Right = node3
	node2.Right = node4
	ret := tree2str(node1)
	fmt.Println(ret)
}
