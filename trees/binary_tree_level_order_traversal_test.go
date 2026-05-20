package trees

import (
	"fmt"
	"testing"
)

func Test_levelOrder(*testing.T) {
	res := levelOrder(input("[1,2,3,4,5,6,7]"))
	for _, v := range res {
		fmt.Println(v)
	}
}
