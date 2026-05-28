package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_design_add_and_search_word_data_structure(t *testing.T) {
	wordDictionary := ConstructorWordDictionary()
	wordDictionary.AddWord("day")
	wordDictionary.AddWord("bay")
	wordDictionary.AddWord("may")
	assert.Equal(t, false, wordDictionary.Search("say"))
	assert.Equal(t, true, wordDictionary.Search("day"))
	assert.Equal(t, true, wordDictionary.Search(".ay"))
	assert.Equal(t, true, wordDictionary.Search("b.."))
}
