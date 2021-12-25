package _3_奇偶树

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	// BFS
	level := 0                 // level代表层数
	prev := root.Val           // 每一层的前一个数字, 在开始新的一层时. prev更新为第一个
	queue := []*TreeNode{root} // 将root放入队列中
	theFirstNode := 1
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			node := queue[0]
			queue = queue[1:]
			if level&1 == 1 {
				// 奇数层需要偶数
				if theFirstNode == 1 {
					theFirstNode++
					if node.Val&1 == 1 {
						return false
					}
				} else {
					if node.Val&1 == 1 || node.Val >= prev {
						return false
					}
				}
			} else {
				// 偶数层需要奇数
				if theFirstNode == 1 {
					theFirstNode++
					if node.Val&1 == 0 {
						return false
					}
				} else {
					if node.Val&1 == 0 || node.Val <= prev {
						return false
					}
				}
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			prev = node.Val
			length--
		}
		// 每一层的前一个数字, 在开始新的一层时. prev更新为第一个
		if len(queue) > 0 {
			prev = queue[0].Val
		}
		// theFirstNode的作用是判断这个节点是否为每一层的第一个节点, 如果是第一个节点就不需要和prev判断
		theFirstNode = 1
		level++
	}
	return true
}
