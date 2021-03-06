package binary_search_tree

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/queue"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/stack"
)

/**
二叉搜索树
结点大于左子树所有结点，小于右子树所有结点
结点的左右子树都是二叉搜索树

细节：size维护时在何处维护，函数返回值设计等等。
*/

// 采用类似泛型的思想，方便外层复用
type Item interface {
	Less(a Item) bool
	ToString() string
}

// 结点
type node struct {
	item        Item // 保存结点数据
	Left, Right *node
}

func newNode(item Item) *node {
	return &node{item: item}
}

// 二叉搜索树
type BST struct {
	root *node // 根结点
	size int   // 大小
}

// 添加结点
func (b *BST) Add(item Item) {
	b.root = b.add(b.root, item)
}

// 私有方法：添加结点
// 在以node1为根结点的二叉搜索树中添加结点node2
func (b *BST) add(node1 *node, item Item) *node {
	if node1 == nil {
		b.size++
		return newNode(item)
	}

	if item.Less(node1.item) {
		node1.Left = b.add(node1.Left, item)
	} else if node1.item.Less(item) {
		node1.Right = b.add(node1.Right, item)
	} else {
		node1.item = item
	}

	return node1
}

// 前序遍历
func (b *BST) PrevOrderRecursion() {
	if b.Empty() {
		return
	}

	b.prevOrderRecursion(b.root)
}

// 前序遍历：以node为根的二叉搜索树
func (b *BST) prevOrderRecursion(node *node) {
	if node == nil {
		return
	}
	fmt.Print(node.item.ToString(), " ")
	b.prevOrderRecursion(node.Left)
	b.prevOrderRecursion(node.Right)
}

// 前序遍历，循环实现
func (b *BST) PrevOrder() {
	if b.Empty() {
		return
	}

	b.prevOrder(b.root)
}

// 前序遍历：以node为根的二叉搜索树
func (b *BST) prevOrder(n *node) {
	s := stack.NewStack(10)
	s.Push(n)

	for !s.Empty() {
		n = s.Pop().(*node)
		fmt.Print(n.item.ToString(), " ")
		// 前序遍历是先遍历左子树，后遍历右子树，因为栈是先进后出，所以把右子树先压入栈
		if n.Right != nil {
			s.Push(n.Right)
		}
		if n.Left != nil {
			s.Push(n.Left)
		}
	}
}

// 中序遍历，采用循环实现，借助栈实现
// TODO 注意：非递归写法一定会用到栈
func (b *BST) InOrder() {
	if b.Empty() {
		return
	}

	b.inOrder(b.root)
}

// 用栈
func (b *BST) inOrder(n *node) {
	s := stack.NewStack(10)
	for !s.Empty() || n != nil {
		if n != nil {
			s.Push(n)
			n = n.Left
		} else {
			n = s.Pop().(*node)
			fmt.Print(n.item.ToString(), " ")
			n = n.Right
		}
	}
}

// 中序遍历递归实现
func (b *BST) InOrderRecursion() {
	b.inOrderRecursion(b.root)
}

func (b *BST) inOrderRecursion(n *node) {
	if n != nil {
		b.inOrderRecursion(n.Left)
		fmt.Print(n.item.ToString(), " ")
		b.inOrderRecursion(n.Right)
	}
}

// 后序遍历，递归实现
func (b *BST) PostOrderRecursion() {
	b.postOrderRecursion(b.root)
}

func (b *BST) postOrderRecursion(n *node) {
	if n != nil {
		b.postOrderRecursion(n.Left)
		b.postOrderRecursion(n.Right)
		fmt.Print(n.item.ToString(), " ")
	}
}

// 后序遍历，循环实现
func (b *BST) PostOrder() {
	if b.root != nil {
		b.postOrder(b.root)
	}
}

func (b *BST) postOrder(n *node) {
	s := stack.NewStack(10)
	s.Push(n)
	s.Push(n)

	for !s.Empty() {
		n := s.Pop().(*node)
		if !s.Empty() && n == s.Peek() {
			if n.Right != nil {
				s.Push(n.Right)
				s.Push(n.Right)
			}
			if n.Left != nil {
				s.Push(n.Left)
				s.Push(n.Left)
			}
		} else {
			fmt.Print(n.item.ToString(), " ")
		}
	}
}

// 二叉搜索树的层序遍历（广度优先遍历）
// 需要借助队列来实现
func (b *BST) LevelOrder() {
	if b.Empty() {
		return
	}

	q := queue.NewQueue(10)
	q.Push(b.root)

	for !q.Empty() {
		n := q.Pop().(*node)
		fmt.Print(n.item.(Item).ToString(), " ")
		if n.Left != nil {
			q.Push(n.Left)
		}
		if n.Right != nil {
			q.Push(n.Right)
		}
	}
}

// 获取最小值，根据二叉搜索树的定义，最左边的左结点就是这个最小值
func (b *BST) MinNum() Item {
	return (b.minNum(b.root)).item
}

// 获取以node为根的二叉搜索树的最小值
func (b *BST) minNum(node *node) *node {
	if node.Left == nil {
		return node
	}
	return b.minNum(node.Left)
}

// 循环实现
func (b *BST) MinNumCircul() (item Item) {
	if !b.Empty() {
		n := b.root
		for n.Left != nil {
			n = n.Left
		}
		item = n.item
	}

	return
}

// 获取最大值，根据二叉搜索树的定义，最右边的右结点就是这个最大值
func (b *BST) MaxNum() Item {
	return (b.maxNum(b.root)).item
}

// 获取以node为根的二叉搜索树的最大值
func (b *BST) maxNum(node *node) *node {
	if node.Right == nil {
		return node
	}

	return b.maxNum(node.Right)
}

// 循环实现
func (b *BST) MaxNumCircul() (item Item) {
	if !b.Empty() {
		n := b.root
		for n.Right != nil {
			n = n.Right
		}
		item = n.item
	}

	return
}

// 删除二叉树的最小值
func (b *BST) RemoveMin() {
	b.size--
	b.root = b.removeMin(b.root)
}

// 删除以node为根的二叉搜索树，返回的是新树的根结点
func (b *BST) removeMin(node *node) *node {
	//如果最小值是叶子结点，就直接删除即可
	//如果最小值是非叶子结点（只有右子树），则删除后，把右子树作为移动到该元素的位置即可
	if node.Left == nil { //nil可以看成一个二叉搜索树的根结点，所以不用去检测右子树是否为nil
		retNode := node.Right
		node = nil //手动删除node，虽然go有GC
		return retNode
	}

	node.Left = b.removeMin(node.Left)
	return node
}

// 删除二叉树的最大值
func (b *BST) RemoveMax() {
	b.size--
	b.root = b.removeMax(b.root)
}

// 删除以node为根的二叉搜索树，返回的是新树的根结点
func (b *BST) removeMax(node *node) *node {
	if node.Right == nil {
		retNode := node.Left
		node = nil //手动删除node，虽然go有GC
		return retNode
	}

	node.Right = b.removeMax(node.Right)
	return node
}

// 查找二叉搜索树
func (b *BST) Find(item Item) Item {
	n := b.find(b.root, item)
	// 必须要验证返回值是否为nil
	if n == nil {
		return nil
	}
	return n.item
}

// 查找以node为根的二叉搜索树
func (b *BST) find(n *node, item Item) *node {
	if n == nil {
		return n
	}

	if item.Less(n.item) {
		return b.find(n.Left, item)
	} else if n.item.Less(item) {
		return b.find(n.Right, item)
	} else {
		return n
	}
}

// 查找循环实现
func (b *BST) Find1(item Item) Item {
	n := b.root
	for n != nil {
		switch {
		case item.Less(n.item):
			n = n.Left
		case n.item.Less(item):
			n = n.Right
		default:
			return n.item
		}
	}

	return nil
}

// 删除二叉搜索树的值为val的结点（假设值不重复）
func (b *BST) Remove(item Item) {
	b.root = b.remove(b.root, item)
	return
}

// 删除以node为根的二叉搜索树中值为val的这个结点，返回删除后树的根结点
func (b *BST) remove(n *node, item Item) *node {
	// 如果该结点是叶子结点，直接删除。
	if n == nil { //没有找到
		return n
	}

	if item.Less(n.item) {
		n.Left = b.remove(n.Left, item)
		return n
	} else if n.item.Less(item) {
		n.Right = b.remove(n.Right, item)
		return n
	} else { // 找到，则删除
		b.size--
		// 如果该结点只有一个子树，则删除后，子树替代该结点即可
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		} else { // 如果该结点有左右子树，则情况比较复杂。后继结点定义
			// 找到比待删除结点大的最小结点，即待删除结点右子树的最小结点
			// 用过这个结点顶替待删除结点即可，因此上面的查找最小值函数minNum应该修改为返回Node指针

			successor := b.minNum(n.Right)
			successor.Right = b.removeMin(n.Right)
			successor.Left = n.Left
			n = nil //删除n结点
			return successor
		}
	}

}

func (b *BST) Empty() bool {
	return b.size == 0
}

func (b *BST) GetSize() int {
	return b.size
}

func NewBst(item Item) *BST {
	return &BST{root: newNode(item), size: 1}
}
