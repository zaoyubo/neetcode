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
	stack := NewTreeStack()
	cur := root
	for !stack.Empty() || cur != nil {
		// 沿着每一个节点的左节点走到头
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		// 当前节点一定是叶子节点 -> 访问
		cur = stack.Pop()
		k--
		if k == 0 {
			return cur.Val
		}
		// cur.Right != nil 将 cur.Right 的左支访问到底
		// cur.Right = nil，引起回溯
		cur = cur.Right
	}
	return 0
}

// morris 遍历，使用空节点的右节点记录next指针，空间O(1)
func kthSmallestV3(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	cur := root
	for cur != nil {
		predecessor := getPredecessor(cur, cur)
		if predecessor != nil {
			// 回溯
			if predecessor.Right == cur {
				predecessor.Right = nil
				// process
				k--
				if k == 0 {
					return cur.Val
				}
				cur = cur.Right
			} else {
				// 往左走
				predecessor.Right = cur
				cur = cur.Left
			}
		} else { // 往左走到头了
			// 输出
			k--
			if k == 0 {
				return cur.Val
			}
			// 右边
			cur = cur.Right
		}
	}
	return 0
}

// 有可能 right 指针存储了 parent 节点，所以要传入 cur 做截断
func getPredecessor(root, cur *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return nil
	}
	tmp := root.Left
	for tmp.Right != nil && tmp.Right != cur {
		tmp = tmp.Right
	}
	return tmp
}
