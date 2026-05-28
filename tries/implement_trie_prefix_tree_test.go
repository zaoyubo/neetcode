package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixTree(t *testing.T) {
	prefixTree := Constructor()
	prefixTree.Insert("dog")
	assert.Equal(t, true, prefixTree.Search("dog")) // return true
	assert.Equal(t, false, prefixTree.Search("do"))
	assert.Equal(t, true, prefixTree.StartsWith("do"))
	prefixTree.Insert("do")
	assert.Equal(t, true, prefixTree.Search("do"))
}
