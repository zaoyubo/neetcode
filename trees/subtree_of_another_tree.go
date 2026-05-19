package trees

import "strconv"

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return kmp(convertToStr(root), convertToStr(subRoot))
}

const (
	Separator = "$"
	NullDesc  = "#"
)

func convertToStr(root *TreeNode) string {
	var result string
	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			result = result + NullDesc + Separator
			return
		}
		result = result + strconv.Itoa(node.Val) + Separator
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return result
}

func kmp(t, s string) bool {
	a := []rune(t)
	b := []rune(s)
	nextval := generateNextVal(s)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
			continue
		} else {
			j = nextval[j]
			if j == -1 {
				i++
				j = 0
			}
		}
	}
	if j == len(b) {
		return true
	}
	return false
}

func generateNextVal(needle string) []int {
	runes := []rune(needle)
	n := len(runes)

	if n == 0 {
		return nil
	}

	nextVal := make([]int, n)
	nextVal[0] = -1

	pre := -1 // pre 表示前一个位置的 next 值

	for i := 1; i < n; i++ {
		//fmt.Println(i-1, pre)
		for pre != -1 && runes[pre] != runes[i-1] {
			pre = nextVal[pre]
		}

		pre++
		if runes[pre] == runes[i] {
			nextVal[i] = nextVal[pre]
		} else {
			nextVal[i] = pre
		}
	}
	//fmt.Println(n-1, pre)

	return nextVal
}

// ----- 下面是题目衍生的思考
func getNextVal(needle string) []int {
	n := len(needle)
	// 假设只包含 ascii 字符
	if n == 0 {
		return []int{}
	}
	nextval := make([]int, n)
	nextval[0] = -1
	pre := -1
	for i := 1; i < n; i++ {
		for pre != -1 && needle[i-1] != needle[pre] {
			pre = nextval[pre]
		}
		pre++
		if needle[pre] == needle[i] {
			nextval[i] = nextval[pre]
		} else {
			nextval[i] = pre
		}
	}
	return nextval
}

func generateNext(needle string) []int {
	runes := []rune(needle)
	n := len(runes)

	if n == 0 {
		panic("needle length must be > 0")
	}

	next := make([]int, n)
	next[0] = -1

	pre := -1 // pre 表示前一个位置的 nextVal 值

	for i := 1; i < n; i++ {
		for pre != -1 && runes[pre] != runes[i-1] {
			pre = next[pre]
		}

		pre++
		next[i] = pre
	}

	return next
}
