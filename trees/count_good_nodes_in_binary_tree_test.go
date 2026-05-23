package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_count_good_nodes_in_binary_tree(t *testing.T) {
	root := input("[2,1,1,3,null,1,5]")
	assert.Equal(t, goodNodes(root), 3)
	assert.Equal(t, goodNodesV2(root), 3)
}

func Test_count_good_nodes_in_binary_tree_case_equal(t *testing.T) {
	root := input("[3,3,null,4,2]")
	assert.Equal(t, goodNodes(root), 3)
	assert.Equal(t, goodNodesV2(root), 3)
}
