package trees

// Breadth First Search Time:O(n) Space: O(n)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	res := make([][]int, 0)
	for len(list) > 0 {
		n := len(list)
		tmp := make([]int, 0)
		for i := 0; i < n; i++ {
			tmp = append(tmp, list[i].Val)
			if list[i].Left != nil {
				list = append(list, list[i].Left)
			}
			if list[i].Right != nil {
				list = append(list, list[i].Right)
			}
		}
		res = append(res, tmp)
		list = list[n:]
	}
	return res
}
