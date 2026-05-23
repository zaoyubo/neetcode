package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	tree := buildTree([]int{1, 2, 4, 6, 3, 5, 7}, []int{4, 6, 2, 1, 3, 7, 5})
	assert.Equal(t, output(tree), "[1,2,3,4,null,null,5,null,6,7]")
}
