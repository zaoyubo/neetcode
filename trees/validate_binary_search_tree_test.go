package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidBST_true(t *testing.T) {
	res := isValidBST(input("[2,1,3]"))
	assert.Equal(t, true, res)
}

func Test_isValidBST_false(t *testing.T) {
	res := isValidBST(input("[1,2,3]"))
	assert.Equal(t, false, res)
}

func Test_isValidBST_duplicate(t *testing.T) {
	res := isValidBST(input("[2,2,2]"))
	assert.Equal(t, false, res)
}
