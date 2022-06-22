package _03_找树左下角的值

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	// 层序遍历
	queue := []*TreeNode{root}
	ans := root.Val
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		ans = node.Val
	}
	return ans
}
