package _2_求根节点到叶节点的数字之和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 自己的解法
func sumNumbers(root *TreeNode) (sum int) {
	// dfs + 回溯
	if root == nil {
		return -1
	}
	var dfs func(root *TreeNode, path []int)
	dfs = func(root *TreeNode, path []int) {
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			sum += getSum(path)
			path = path[:len(path)-1]
			return
		}
		if root.Left != nil {
			dfs(root.Left, path)
		}
		if root.Right != nil {
			dfs(root.Right, path)
		}
		path = path[:len(path)-1]
	}
	dfs(root, []int{})
	return sum
}

func getSum(path []int) int {
	sum := 0
	count := 1
	for i := len(path) - 1; i >= 0; i-- {
		sum += count * path[i]
		count *= 10
	}
	return sum
}

// 官方解法

func sumNumbers2(root *TreeNode) (sum int) {
	return dfs(root, 0)
}

func dfs(root *TreeNode, curSum int) int {
	if root == nil {
		return 0
	}
	sum := curSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}
