package trees

// 中序遍历 + 剪枝
// 哪些地方可能需要剪枝？触发递归操作返回之后
func kthSmallest(root *TreeNode, k int) int {
	var dfs func(root *TreeNode)
	var cnt int
	var res int
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		// 剪枝：如果从 dfs(root.Left) 返回后，已经满足条件了，就没必要继续了
		if cnt >= k {
			return
		}

		cnt++
		if cnt == k {
			res = root.Val
			return
		}

		dfs(root.Right)
		// 这里剪不剪都行，后面已经没有递归操作了
	}
	dfs(root)
	return res
}

// 用栈实现非递归中序遍历
func kthSmallestV2(root *TreeNode, k int) int {
	// TODO
	return 0
}

// morris 遍历，使用空节点的右节点记录next指针，空间O(1)
func kthSmallestV3(root *TreeNode, k int) int {
	// TODO
	return 0
}
