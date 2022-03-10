package _41_N叉树的前序遍历

type Node struct {
	Val      int
	Children []*Node
}

// dfs 递归
func preorder(root *Node) (ans []int) {
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		for i := 0; i < len(root.Children); i++ {
			dfs(root.Children[i])
		}
	}
	return
}

// 队列迭代, 从最后一个取, 那么就是前序遍历了 , 不错!!!

func PreOrder(root *Node) (ans []int) {
	queue := []*Node{root}
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		ans = append(ans, root.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			queue = append(queue, node.Children[i])
		}
	}
	return
}
