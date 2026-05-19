package trees

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历
func input(s string) *TreeNode {

	s = strings.Replace(s, "null", strconv.Itoa(math.MaxInt), -1)
	var values []int
	cnt := 0
	firstFlag := true
	json.Unmarshal([]byte(s), &values)
	var list []*TreeNode
	var root *TreeNode
	for _, v := range values {
		var node *TreeNode
		if v == math.MaxInt {
		} else {
			node = &TreeNode{
				Val: v,
			}
		}
		if firstFlag {
			root = node
			firstFlag = false
		}
		if len(list) != 0 {
			if cnt%2 == 0 {
				list[0].Left = node
			} else {
				list[0].Right = node
				list = list[1:]
			}
			cnt++
		}
		if node != nil {
			list = append(list, node)
		}
	}
	return root
}

// 层序遍历 + 剪去末尾连续的空节点
func output(root *TreeNode) string {
	var list []*TreeNode
	list = append(list, root)
	idx := 0
	var values []int
	// 广度实现层序遍历，需要 list 辅助
	for len(list) != 0 {
		node := list[idx]
		if node == nil {
			// 空节点不生孩子
			values = append(values, math.MaxInt) // math.MaxInt 表示nil
		} else {
			values = append(values, node.Val)
			list = append(list, node.Left, node.Right)
		}
		list = list[1:]
	}
	// 剪去末尾连续的空节点
	for len(values) > 0 && values[len(values)-1] == math.MaxInt {
		values = values[:len(values)-1]
	}
	b, _ := json.Marshal(values)
	return strings.Replace(string(b), strconv.Itoa(math.MaxInt), "null", -1)
}

// 先序遍历
func printTree(root *TreeNode) {
	if root == nil {
		println("<nil>")
		return
	} else {
		fmt.Println(root.Val)
	}
	printTree(root.Left)
	printTree(root.Right)
}
