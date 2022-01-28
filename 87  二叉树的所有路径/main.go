package _7__二叉树的所有路径

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			res = append(res, path+"->"+strconv.Itoa(root.Val))
			return
		}
		dfs(root.Left, path+"->"+strconv.Itoa(root.Val))
		dfs(root.Right, path+"->"+strconv.Itoa(root.Val))
	}
	dfs(root, "")
	for i := 0; i < len(res); i++ {
		res[i] = res[i][2:]
	}
	return res
}
