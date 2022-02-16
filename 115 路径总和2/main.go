package _15_路径总和2

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	// 深度BFS
	if root == nil {
		return nil
	}
	res := [][]int{}
	var dfs func(root *TreeNode, targetSum int, path []int)
	dfs = func(root *TreeNode, targetSum int, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if targetSum == root.Val && root.Left == nil && root.Right == nil {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		dfs(root.Left, targetSum-root.Val, path)
		dfs(root.Right, targetSum-root.Val, path)
		path = path[:len(path)-1]
	}
	dfs(root, targetSum, []int{})
	return res
}
