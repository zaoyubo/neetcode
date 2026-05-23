package trees

import "math"

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
	var recur func(root *TreeNode, min, max int) bool
	// param min: root 的所有节点都应该大于 min
	// param max:  root 的所有节点都应该小于 max
	// return bool: root 是否是 valid bst
	recur = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}
		if root.Val <= min || root.Val >= max {
			return false
		}
		a := recur(root.Left, min, root.Val)
		b := recur(root.Right, root.Val, max)
		return a && b
	}
	return recur(root, math.MinInt, math.MaxInt)
}

// 方法三：利用 BST 左侧节点 < 当前节点< 右侧节点的原理，
// 广度优先，队列元素记录当前树需要满足的最小值和最大值
func isValidBSTV3(root *TreeNode) bool {
	type advanceNode struct {
		root *TreeNode
		min  int // root 的所有节点都应该大于 min
		max  int // root 的所有节点都应该小于 max
	}
	list := make([]*advanceNode, 0)
	list = append(list, &advanceNode{
		root: root,
		min:  math.MinInt,
		max:  math.MaxInt,
	})
	for len(list) > 0 {
		cur := list[0]
		list = list[1:]
		minValue := cur.min
		maxValue := cur.max
		if cur.root.Val <= minValue || cur.root.Val >= maxValue {
			return false
		}
		if cur.root.Left != nil {
			list = append(list, &advanceNode{
				root: cur.root.Left,
				min:  minValue,
				max:  cur.root.Val,
			})
		}
		if cur.root.Right != nil {
			list = append(list, &advanceNode{
				root: cur.root.Right,
				min:  cur.root.Val,
				max:  maxValue,
			})
		}
	}
	return true
}
