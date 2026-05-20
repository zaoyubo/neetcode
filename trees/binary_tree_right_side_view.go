package trees

/*
方法一：层序遍历，再截断. -> 使用了广度优先的思想，还可优化
*/
func rightSideView(root *TreeNode) []int {
	total := levelOrder(root)
	var res []int
	for _, l := range total {
		if len(l) > 0 {
			res = append(res, l[len(l)-1])
		}
	}
	return res
}

/*
方法二：
1. 先序遍历，然后先访问右子节点，再访问左子节点
2. map[depth]struct{}{} 记录每一层是否有访问过 -> 可优化
*/
func rightSideViewV2(root *TreeNode) []int {
	var dfs func(root *TreeNode, depth int)
	var res []int
	depthMap := make(map[int]struct{})
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if _, ok := depthMap[depth]; !ok {
			res = append(res, root.Val)
			depthMap[depth] = struct{}{}
		}
		// need right view, so travel right first
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}
	dfs(root, 0)
	return res
}
