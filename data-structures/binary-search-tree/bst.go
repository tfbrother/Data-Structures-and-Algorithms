package binary_search_tree

import "fmt"

/**
二叉搜索树
结点大于左子树所有结点，小于右子树所有结点
结点的左右子树都是二叉搜索树
*/

// 结点
type Node struct {
	Value       int
	Left, Right *Node
}

// 二叉搜索树
type bst struct {
	root *Node // 根结点
	size int   // 大小
}

// 添加结点
func (b *bst) Add(node *Node) {
	b.root = b.add(b.root, node)
}

// 私有方法：添加结点
// 在以node1为根结点的二叉搜索树中添加结点node2
func (b *bst) add(node1 *Node, node2 *Node) *Node {
	if node1 == nil {
		return node2
	}

	if node1.Value > node2.Value {
		node1.Left = b.add(node1.Left, node2)
	}
	if node1.Value < node2.Value {
		node1.Right = b.add(node1.Right, node2)
	}

	return node1
}

// 前序遍历
func (b *bst) PrevOrder() {
	if b.Empty() {
		return
	}

	b.prevOrder(b.root)

}

// 前序遍历：以node为根的二叉搜索树
func (b *bst) prevOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	b.prevOrder(node.Left)
	b.prevOrder(node.Right)
}

func (b *bst) Empty() bool {
	return b.size == 0
}

func (b *bst) GetSize() int {
	return b.size
}

func NewBst(root *Node) *bst {
	return &bst{root: root, size: 1}
}
