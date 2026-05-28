package tries

type PrefixTree struct {
	root *Node
}

// word and prefix are made up of lowercase English letters.
type Node struct {
	children [26]*Node
	isEnd    bool
}

func Constructor() PrefixTree {
	return PrefixTree{
		root: &Node{
			children: [26]*Node{},
		},
	}
}

func idx(b byte) int {
	return int(b - 'a')
}

func (this *PrefixTree) Insert(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if cur.children[idx(ch)] == nil {
			cur.children[idx(ch)] = &Node{
				children: [26]*Node{},
			}
		}
		cur = cur.children[idx(ch)]
	}
	cur.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	cur := this.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if cur.children[idx(ch)] == nil {
			return false
		}
		cur = cur.children[idx(ch)]
	}
	return cur.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	cur := this.root
	for i := 0; i < len(prefix); i++ {
		ch := prefix[i]
		if cur.children[idx(ch)] == nil {
			return false
		}
		cur = cur.children[idx(ch)]
	}
	return true
}
