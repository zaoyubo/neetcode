package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidBST_true(t *testing.T) {
	root := input("[2,1,3]")
	assert.Equal(t, true, isValidBST(root))
	assert.Equal(t, true, isValidBSTV2(root))
	assert.Equal(t, true, isValidBSTV3(root))
}

func Test_isValidBST_false(t *testing.T) {
	root := input("[1,2,3]")
	assert.Equal(t, false, isValidBST(root))
	assert.Equal(t, false, isValidBSTV2(root))
	assert.Equal(t, false, isValidBSTV3(root))
}

func Test_isValidBST_duplicate(t *testing.T) {
	root := input("[2,2,2]")
	assert.Equal(t, false, isValidBST(root))
	assert.Equal(t, false, isValidBSTV2(root))
	assert.Equal(t, false, isValidBSTV3(root))
}

func Test_isValidBST_complicate(t *testing.T) {
	root := input("[5,3,8,1,4,7,9,null,2]")
	assert.Equal(t, true, isValidBST(root))
	assert.Equal(t, true, isValidBSTV2(root))
	assert.Equal(t, true, isValidBSTV3(root))
}
