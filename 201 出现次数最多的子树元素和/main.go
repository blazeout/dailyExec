package _01_出现次数最多的子树元素和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	// 计算结果出现的次数
	numberOfOcc := map[int]int{}
	// maxNumber保存出现次数最大的那一个
	maxNumber := 0
	// 从下至上计算减少时间复杂度
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sum := root.Val
		sum += dfs(root.Left)
		sum += dfs(root.Right)
		numberOfOcc[sum]++
		if numberOfOcc[sum] > maxNumber {
			maxNumber = numberOfOcc[sum]
		}
		return sum
	}
	dfs(root)
	ans := make([]int, 0)
	for key, value := range numberOfOcc {
		if value == maxNumber {
			ans = append(ans, key)
		}
	}
	return ans
}
