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

// 二叉搜索树的层序遍历（广度优先遍历）
// 需要借助队列来实现
func (b *bst) LevelOrder() {

}

// 获取最小值，根据二叉搜索树的定义，最左边的左结点就是这个最小值
func (b *bst) MinNum() int {
	return b.minNum(b.root)
}

// 获取以node为根的二叉搜索树的最小值
func (b *bst) minNum(node *Node) int {
	if node.Left == nil {
		return node.Value
	}
	return b.minNum(node.Left)
}

// 获取最大值，根据二叉搜索树的定义，最右边的右结点就是这个最大值
func (b *bst) MaxNum() int {
	return b.maxNum(b.root)
}

// 获取以node为根的二叉搜索树的最大值
func (b *bst) maxNum(node *Node) int {
	if node.Right == nil {
		return node.Value
	}

	return b.maxNum(node.Right)
}

// 删除二叉树的最小值
func (b *bst) RemoveMin() {
	b.root = b.removeMin(b.root)
}

// 删除以node为根的二叉搜索树，返回的是新树的根结点
func (b *bst) removeMin(node *Node) *Node {
	//如果最小值是叶子结点，就直接删除即可
	//如果最小值是非叶子结点（只有右子树），则删除后，把右子树作为移动到该元素的位置即可
	if node.Left == nil { //nil可以看成一个二叉搜索树的根结点，所以不用去检测右子树是否为nil
		return node.Right
	}

	node.Left = b.removeMin(node.Left)
	return node
}

// 删除二叉树的最大值
func (b *bst) RemoveMax() {
	b.root = b.removeMax(b.root)
}

// 删除以node为根的二叉搜索树，返回的是新树的根结点
func (b *bst) removeMax(node *Node) *Node {
	if node.Right == nil {
		return node.Left
	}

	node.Right = b.removeMax(node.Right)
	return node
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
