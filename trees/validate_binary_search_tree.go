package trees

// 方法一：中序遍历后比较大小
func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode)
	var seq []int
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		seq = append(seq, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	for i := 0; i < len(seq)-1; i++ {
		if seq[i] >= seq[i+1] {
			return false
		}
	}
	return true
}

// 方法二：利用 BST 左侧节点 < 当前节点< 右侧节点的原理，
// 递归检查，深度优先，param 传入子树需要满足的最小值和最大值
func isValidBSTV2(root *TreeNode) bool {
	return false
}

// 方法三：利用 BST 左侧节点 < 当前节点< 右侧节点的原理，
// 广度优先，队列元素记录当前树需要满足的最小值和最大值
func isValidBSTV3(root *TreeNode) bool {
	return false
}
