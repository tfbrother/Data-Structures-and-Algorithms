package read_black_tree

import (
	"fmt"
)

/**
红黑树，底层基于二叉搜索树实现。只是在二叉搜索树的基础上增加来颜色的维护。性质和2-3树等价
结点属性：
1.空结点为黑色
2.新加入的结点默认为红色
3.根结点始终为黑色
4.红色结点的两个子结点为黑色，黑色结点的右孩子一定是黑色的
5.根结点到所有叶子结点经过的黑色结点数量相同===黑平衡
*/

// 采用类似泛型的思想，方便外层复用
type Item interface {
	Less(a Item) bool
	ToString() string
}

// 结点
type node struct {
	item        Item // 保存结点数据
	black       bool // 颜色
	Left, Right *node
}

// 辅助函数
// 空结点为黑色
func isRed(n *node) bool {
	if n == nil || n.black {
		return false
	}

	return true
}

// 新创建的为红色
func newNode(item Item) *node {
	return &node{item: item, black: false}
}

// 二叉搜索树
type RBTree struct {
	root *node // 根结点
	size int   // 大小
}

// 添加结点
func (r *RBTree) Add(item Item) {
	r.root = r.add(r.root, item)
	// 维护根结点始终为黑色
	r.root.black = true
}

// 私有方法：添加结点
// 在以node1为根结点的二叉搜索树中添加结点node2
func (r *RBTree) add(n *node, item Item) *node {
	if n == nil {
		r.size++
		return newNode(item)
	}

	if item.Less(n.item) {
		n.Left = r.add(n.Left, item)
	} else if n.item.Less(item) {
		n.Right = r.add(n.Right, item)
	} else {
		n.item = item
	}

	// 颜色维护
	// 左旋转
	if !isRed(n.Left) && isRed(n.Right) { // 说明是往一个二结点的右边插入了一个新结点
		n = r.leftRotate(n)
	}

	// 右旋转
	if isRed(n.Left) && isRed(n.Left.Left) { // 说明是往一个三结点的左边插入了一个新结点
		n = r.rightRotate(n)
	}

	// 颜色反转
	if isRed(n.Left) && isRed(n.Right) { // 往一个三结点右边插入了一个新结点，保持红黑树属性
		r.flipColors(n)
	}

	return n
}

//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func (r *RBTree) leftRotate(n *node) *node {
	x := n.Right

	// 旋转
	n.Right = x.Left
	x.Left = n
	x.black = n.black
	n.black = false

	return x
}

//     node                   x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   node
//  / \                       /  \
// y  T1
func (r *RBTree) rightRotate(n *node) *node {
	x := n.Left

	// 旋转
	n.Left = x.Right
	x.Right = n
	x.black = n.black
	n.black = false

	return x
}

// 颜色反转
func (r *RBTree) flipColors(n *node) {
	n.black = !n.black
	n.Left.black = !n.Left.black
	n.Right.black = !n.Right.black
}

// 前序遍历
func (r *RBTree) PrevOrder() {
	if r.Empty() {
		return
	}

	r.prevOrder(r.root)
}

// 前序遍历：以node为根的二叉搜索树
func (r *RBTree) prevOrder(node *node) {
	if node == nil {
		return
	}
	fmt.Println(node.item.ToString())
	r.prevOrder(node.Left)
	r.prevOrder(node.Right)
}

// 中序遍历，采用循环实现，借助栈实现
// TODO 注意：非递归写法一定会用到栈
func (r *RBTree) InOrder() {
	if r.Empty() {
		return
	}

	r.inOrder(r.root)
}

// 用栈模拟递归
func (r *RBTree) inOrder(node *node) {
	for node != nil {

	}
}

// 二叉搜索树的层序遍历（广度优先遍历）
// 需要借助队列来实现
func (r *RBTree) LevelOrder() {

}

// 获取最小值，根据二叉搜索树的定义，最左边的左结点就是这个最小值
func (r *RBTree) MinNum() Item {
	return (r.minNum(r.root)).item
}

// 获取以node为根的二叉搜索树的最小值
func (r *RBTree) minNum(node *node) *node {
	if node.Left == nil {
		return node
	}
	return r.minNum(node.Left)
}

// 获取最大值，根据二叉搜索树的定义，最右边的右结点就是这个最大值
func (r *RBTree) MaxNum() Item {
	return (r.maxNum(r.root)).item
}

// 获取以node为根的二叉搜索树的最大值
func (r *RBTree) maxNum(node *node) *node {
	if node.Right == nil {
		return node
	}

	return r.maxNum(node.Right)
}

// 删除二叉树的最小值
func (r *RBTree) RemoveMin() {
	r.size--
	r.root = r.removeMin(r.root)
}

// 删除以node为根的二叉搜索树，返回的是新树的根结点
func (r *RBTree) removeMin(node *node) *node {
	//如果最小值是叶子结点，就直接删除即可
	//如果最小值是非叶子结点（只有右子树），则删除后，把右子树作为移动到该元素的位置即可
	if node.Left == nil { //nil可以看成一个二叉搜索树的根结点，所以不用去检测右子树是否为nil
		return node.Right
	}

	node.Left = r.removeMin(node.Left)
	return node
}

// 删除二叉树的最大值
func (r *RBTree) RemoveMax() {
	r.size--
	r.root = r.removeMax(r.root)
}

// 删除以node为根的二叉搜索树，返回的是新树的根结点
func (r *RBTree) removeMax(node *node) *node {
	if node.Right == nil {
		return node.Left
	}

	node.Right = r.removeMax(node.Right)
	return node
}

// 查找二叉搜索树
func (r *RBTree) Find1(item Item) Item {
	n := r.find(r.root, item)
	// 必须要验证返回值是否为nil
	if n == nil {
		return nil
	}
	return n.item
}

// 查找以node为根的二叉搜索树
func (r *RBTree) find(n *node, item Item) *node {
	if n == nil {
		return n
	}

	if item.Less(n.item) {
		return r.find(n.Left, item)
	} else if n.item.Less(item) {
		return r.find(n.Right, item)
	} else {
		return n
	}
}

// 查找循环实现
func (r *RBTree) Find(item Item) Item {
	n := r.root
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
func (r *RBTree) Remove(item Item) {
	r.root = r.remove(r.root, item)
	return
}

// 删除以node为根的二叉搜索树中值为val的这个结点，返回删除后树的根结点
func (r *RBTree) remove(n *node, item Item) *node {
	// 如果该结点是叶子结点，直接删除。
	if n == nil { //没有找到
		return n
	}

	if item.Less(n.item) {
		n.Left = r.remove(n.Left, item)
		return n
	} else if n.item.Less(item) {
		n.Right = r.remove(n.Right, item)
		return n
	} else { // 找到，则删除
		r.size--
		// 如果该结点只有一个子树，则删除后，子树替代该结点即可
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		} else { // 如果该结点有左右子树，则情况比较复杂。后继结点定义
			// 找到比待删除结点大的最小结点，即待删除结点右子树的最小结点
			// 用过这个结点顶替待删除结点即可，因此上面的查找最小值函数minNum应该修改为返回Node指针

			successor := r.minNum(n.Right)
			successor.Right = r.removeMin(n.Left)
			successor.Left = n.Left
			return successor
		}
	}

}

func (r *RBTree) Empty() bool {
	return r.size == 0
}

func (r *RBTree) GetSize() int {
	return r.size
}

func NewBst(item Item) *RBTree {
	return &RBTree{root: newNode(item), size: 1}
}
