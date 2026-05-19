package trees

import (
	"fmt"
	"testing"
)

func Test_input(t *testing.T) {
	r := input("[1,2,3,4,5]")
	printTree(r)
	fmt.Println("-----")
	r = input("[1, null, 3, 4, 5]")
	printTree(r)
}

func Test_output(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	res := output(root)
	fmt.Println(res)
}
