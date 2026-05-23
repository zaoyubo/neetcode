package trees

import "math"

// 记录节点对应路径中最大值，如果 val >= 该值，则为 goodNode
// 深度优先
func goodNodes(root *TreeNode) int {
	var dfs func(root *TreeNode, max int) int
	// max: cur 节点到根节点路径中，最大的值 (不包括cur 节点）
	// return int : cur 对应的树 goodNodes 数量
	dfs = func(cur *TreeNode, max int) int {
		if cur == nil {
			return 0
		}
		includeCur := 0
		if cur.Val >= max {
			includeCur = 1
			max = cur.Val
		}
		return dfs(cur.Left, max) + dfs(cur.Right, max) + includeCur
	}
	return dfs(root, math.MinInt)
}

// 记录节点对应路径中最大值，如果 val >= 该值，则为 goodNode
// 广度优先
func goodNodesV2(root *TreeNode) int {
	type advanceNode struct {
		root   *TreeNode
		maxVal int // 节点对应路径中最大值,不包括当前节点
	}

	// 用来广度优先遍历二叉树
	list := make([]advanceNode, 0)
	list = append(list, advanceNode{
		root:   root,
		maxVal: math.MinInt,
	})

	var res int
	for len(list) > 0 {
		node := list[0]
		list = list[1:]
		maxValue := node.maxVal
		if node.root.Val >= node.maxVal {
			res++
			maxValue = node.root.Val
		}
		if node.root.Left != nil {
			list = append(list, advanceNode{node.root.Left, maxValue})
		}
		if node.root.Right != nil {
			list = append(list, advanceNode{node.root.Right, maxValue})
		}
	}
	return res
}
