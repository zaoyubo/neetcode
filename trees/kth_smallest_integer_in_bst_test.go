package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthSmallest(t *testing.T) {
	root := input("[4,3,5,2,null]")
	assert.Equal(t, kthSmallest(root, 4), 5)
	assert.Equal(t, kthSmallestV2(root, 4), 5)
	assert.Equal(t, kthSmallestV3(root, 4), 5)
}
