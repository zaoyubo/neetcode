package trees

// Both arrays are of the same size and consist of unique values.
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	node := &TreeNode{
		Val: preorder[0],
	}
	inOrderIdx := findIdx(inorder, preorder[0])

	node.Left = buildTree(preorder[1:1+inOrderIdx], inorder[:inOrderIdx])
	node.Right = buildTree(preorder[1+inOrderIdx:], inorder[1+inOrderIdx:])
	return node
}

func findIdx(inorder []int, cur int) int {
	for idx, item := range inorder {
		if item == cur {
			return idx
		}
	}
	return -1
}
