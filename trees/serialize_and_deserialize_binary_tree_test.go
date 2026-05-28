package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {
	c := Codec{}
	s := "[1,2,3,4,5,null,6,null,7,null,null,8]"
	serialized := c.serialize(input(s))
	deserialized := c.deserialize(serialized)
	assert.Equal(t, s, output(deserialized))
}
