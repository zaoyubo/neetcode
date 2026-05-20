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
	p := root.Left
	q := root.Right
	ancestor := lowestCommonAncestor(root, p, q)
	assert.Equal(t, ancestor.Val, 5)
}
