package set

import "github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/binary-search-tree"

// 基于二分搜索树为底层的集合实现
type Node binary_search_tree.Item

type BSTSet struct {
	bst *binary_search_tree.BST
}

// 往集合中添加元素
func (b *BSTSet) Add(node Node) {
	b.bst.Add(node)
}

// 删除集合中的元素
func (b *BSTSet) Remove(node Node) {
	b.bst.Remove(node)
}

// 查询集合中是否包含某元素
func (b *BSTSet) Contains(node Node) bool {
	return b.bst.Find(node) == nil
}

// 获取集合中元素的个数
func (b *BSTSet) Size() int {
	return b.bst.GetSize()
}

// 判断集合是否为空
func (b *BSTSet) Empty() bool {
	return b.bst.Empty()
}
