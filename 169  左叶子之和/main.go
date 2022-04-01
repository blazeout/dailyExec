package _69__左叶子之和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	// 需要输入一个flag代表这个节点是否是做左子节点
	var dfs func(root *TreeNode, flag bool)
	dfs = func(root *TreeNode, flag bool) {
		if root == nil {
			return
		}
		if flag && root.Left == nil && root.Right == nil {
			res += root.Val
		}
		dfs(root.Left, true)
		dfs(root.Right, false)
	}
	dfs(root, false)
	return res
}
