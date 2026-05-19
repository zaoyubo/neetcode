package trees

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSubtree(t *testing.T) {
	root := input("[1,2,3,4,5]")
	subRoot := input("[2,4,5]")
	assert.Equal(t, true, isSubtree(root, subRoot))

	root = input("[1,2,3,4,5,null,null,6]")
	subRoot = input("[2,4,5]")
	assert.Equal(t, false, isSubtree(root, subRoot))

}

func Test_generateNextVal(t *testing.T) {
	fmt.Println(generateNextVal("abaabaacb"))
	assert.Equal(t, generateNextVal("abaabaacb"), getNextVal("abaabaacb"))
}

func Test_generateNext(t *testing.T) {
	fmt.Println(generateNext("abaabaacb"))
}

func Test_kmp(t *testing.T) {
	res := kmp("abaabaacb", "aac")
	assert.Equal(t, true, res)
}
