package trie

// 字典树/前缀树
// 只针对26个英文小写字母实现

type Element struct {
	char   byte // 注意用byte来保存，这样就可以用数组来保存，因为是整型的
	isWord bool
	next   [26]*Element // 用数组来保存，并不用map来保存，方便排序。
}

type Trie struct {
	root  *Element
	count int // 字典树中总共的单词数
}

// 添加一个单词
func (t *Trie) Add(word string) {
	byteStr := []byte(word)
	curr := t.root
	for i := 0; i < len(byteStr); i++ {
		s := byteStr[i] - 'a'
		if curr.next[s] != nil { //找到了，然后继续
			curr = curr.next[s]
		} else { // 没有找到，则添加，然后继续
			curr.next[s] = &Element{s, false, [26]*Element{}}
			curr = curr.next[s]
		}
	}

	// 循环执行完了
	curr.isWord = true
	t.count++
}

func (t *Trie) Find(word string) bool {
	byteStr := []byte(word)
	curr := t.root
	for i := 0; i < len(byteStr); i++ {
		s := byteStr[i] - 'a'
		if curr.next[s] != nil { //找到了，继续
			curr = curr.next[s]
			continue
		} else { // 没有找到，则返回false
			return false
		}
	}

	// 循环执行完了
	return curr.isWord
}

func NewTrie() *Trie {
	return &Trie{
		root:  &Element{0, true, [26]*Element{}},
		count: 0,
	}
}
