package class08

type Node struct {
	pass  int
	end   int
	nexts []*Node
}

type Trie struct {
	root *Node
}

// Insert 往trie树里加一个word
func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	str := []byte(word)
	node := t.root
	node.pass++
	for _, item := range str {
		path := item - 'a'
		if node.nexts[path] == nil {
			node.nexts[path] = &Node{}
		}
		node = node.nexts[path]
		node.pass++
	}
	node.end++
}

// Search word这个单词之前加入过几次
func (t *Trie) Search(word string) int {
	if len(word) == 0 {
		return 0
	}
	str := []byte(word)
	node := t.root
	for _, item := range str {
		path := item - 'a'
		if node.nexts[path] == nil {
			return 0
		}
		node = node.nexts[path]
	}
	return node.end
}

// PrefixNumber 所有加入的字符串中，有几个是以pre这个字符串作为前缀的
func (t *Trie) PrefixNumber(pre string) int {
	if len(pre) == 0 {
		return 0
	}
	str := []byte(pre)
	node := t.root
	for _, item := range str {
		path := item - 'a'
		if node.nexts[path] == nil {
			return 0
		}
		node = node.nexts[path]
	}
	return node.pass
}

// Delete 删除word, 注意避免内存泄露问题
func (t *Trie) Delete(word string) {
	if t.Search(word) == 0 {
		return
	}
	str := []byte(word)
	node := t.root
	node.pass--
	for _, item := range str {
		path := item - 'a'
		if node.nexts[path].pass == 1 {
			node.nexts[path] = nil
		}
		node = node.nexts[path]
	}
	node.end--
}
