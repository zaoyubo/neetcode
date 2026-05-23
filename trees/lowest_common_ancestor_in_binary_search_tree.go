package trees

/*
题目备注：
1. 输入参数中，p、q 不一定是 root 的节点指针，可能是 deep copy
2. root 是二叉搜索树，有如下特性可以利用：
  - leftSubTree.Val < root.Val < righSubTree.Val
  - all node values are unique
*/
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if p == nil || q == nil || root == nil {
		return nil
	}

	cur := root
	for cur != nil {
		if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right
		} else if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left
		} else {
			return cur
		}
	}

	return cur
}
