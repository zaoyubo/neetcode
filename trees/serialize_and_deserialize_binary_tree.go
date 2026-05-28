package trees

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// dfs
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res []string
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			res = append(res, "N")
			return
		}
		res = append(res, strconv.Itoa(root.Val))
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return strings.Join(res, ",")
}

//1 2 4 N 7 N N 5 N N 3 N 6 8 N N N

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(s string) *TreeNode {
	chars := strings.Split(s, ",")
	var idx int
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		char := chars[idx]
		var node *TreeNode
		if char != "N" {
			value, _ := strconv.Atoi(char)
			node = &TreeNode{Val: value}
		}
		idx++
		if node != nil {
			node.Left = dfs()
			node.Right = dfs()
		}
		return node
	}
	return dfs()
}
