package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthSmallest(t *testing.T) {
	root := input("[4,3,5,2,null]")
	assert.Equal(t, 5, kthSmallest(root, 4))
	assert.Equal(t, 5, kthSmallestV2(root, 4))
	assert.Equal(t, 5, kthSmallestV3(root, 4))
}
