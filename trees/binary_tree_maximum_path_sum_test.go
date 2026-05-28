package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxPathSum(t *testing.T) {
	root := input("[1,-2,3,4,5,null,6,-1,-2,null,null,8,3]")
	res := maxPathSum(root)
	res2 := maxPathSumV2(root)
	assert.Equal(t, 21, res)
	assert.Equal(t, 21, res2)
}

func Test_maxPathSum_all_negative(t *testing.T) {
	root := input("[-4,-3,-2]")
	res := maxPathSum(root)
	res2 := maxPathSumV2(root)
	assert.Equal(t, -2, res)
	assert.Equal(t, -2, res2)
}
