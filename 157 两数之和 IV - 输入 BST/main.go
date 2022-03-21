package _57_两数之和_IV___输入_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	// 使用map
	mp := map[int]struct{}{}
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if _, ok := mp[k-root.Val]; ok {
			return true
		}
		mp[root.Val] = struct{}{}
		left, right := dfs(root.Left), dfs(root.Right)
		return left || right
	}
	return dfs(root)

}
