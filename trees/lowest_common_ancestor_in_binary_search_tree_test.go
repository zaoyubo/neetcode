package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor(t *testing.T) {
	root := input("[5,3,8,1,4,7,9,null,2]")
	if root == nil {
		t.Fatal("expected non nil")
	}
	p := &TreeNode{Val: 3}
	q := &TreeNode{Val: 8}
	ancestor := lowestCommonAncestor(root, p, q)
	assert.Equal(t, ancestor.Val, 5)
}

func Test_lowestCommonAncestor_2(t *testing.T) {
	root := input("[5,3,8,1,4,7,9,null,2]")
	if root == nil {
		t.Fatal("expected non nil")
	}
	p := &TreeNode{Val: 2}
	q := &TreeNode{Val: 4}
	ancestor := lowestCommonAncestor(root, p, q)
	assert.Equal(t, ancestor.Val, 3)
}

func Test_lowestCommonAncestor_equal(t *testing.T) {
	root := input("[5,3,8,1,4,7,9,null,2]")
	if root == nil {
		t.Fatal("expected non nil")
	}
	p := &TreeNode{Val: 7}
	q := &TreeNode{Val: 8}
	ancestor := lowestCommonAncestor(root, p, q)
	assert.Equal(t, ancestor.Val, 8)
}
