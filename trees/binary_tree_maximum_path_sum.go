package trees

import "math"

func maxPathSum(root *TreeNode) int {
	maxPathSumValue := math.MinInt
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		leftMax := getMax(root.Left)
		rightMax := getMax(root.Right)
		sum := leftMax + rightMax + root.Val
		maxPathSumValue = max(maxPathSumValue, sum)
		dfs(root.Left)
		dfs(root.Right)
		return
	}
	dfs(root)
	return maxPathSumValue
}

func getMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getMax(root.Left)
	right := getMax(root.Right)
	v := max(left, right) + root.Val
	if v > 0 {
		return v
	}
	return 0
}

func maxPathSumV2(root *TreeNode) int {
	maxPathSumValue := math.MinInt
	var dfs func(root *TreeNode) int
	// return maxDownPath, 从当前节点往下的最大路径，包括当前节点
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := max(dfs(root.Left), 0)
		rightMax := max(dfs(root.Right), 0)

		maxPathSumValue = max(maxPathSumValue, leftMax+rightMax+root.Val)
		return root.Val + max(leftMax, rightMax)
	}
	dfs(root)
	return maxPathSumValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
