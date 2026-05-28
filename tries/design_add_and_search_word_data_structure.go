package tries

// word in search consist of '.' or lowercase English letters.
type WordDictionary struct {
	root *Node
}

// word and prefix are made up of lowercase English letters.
//type Node struct {
//	children [26]*Node
//	isEnd    bool
//}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{
		root: &Node{
			children: [26]*Node{},
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
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

func (this *WordDictionary) Search(word string) bool {
	return search(this.root, word)
}

func search(cur *Node, word string) bool {
	if cur == nil {
		return false
	}
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if ch == []uint8(".")[0] {
			for j := 0; j < len(cur.children); j++ {
				if search(cur.children[j], word[i+1:]) {
					return true
				}
			}
			return false
		}
		index := idx(ch)
		if cur.children[index] == nil {
			return false
		}
		cur = cur.children[index]
	}
	return cur.isEnd
}
