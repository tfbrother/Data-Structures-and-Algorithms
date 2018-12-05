package avl_tree

import (
	"fmt"
	"math"
)

// avl树一种平衡二叉搜索树，所以都是基于二叉搜索树来实现。由于golang不支持泛型，此处不复用之前实现的BST，从新底层实现avl树。

// TODO(tfbrother) 此处接口设计还有问题，后期修正2018-12-05
type Node interface {
	Less(a node) int
}

// 设置成为私有结构体，不对外部暴露node的结构
type node struct {
	key         string
	value       string
	height      int   // 高度
	Left, Right *node // 左右子树
}

func (n node) Less(a node) int {
	if n.value > a.value {
		return 1
	} else if n.value == a.value {
		return 0
	}

	return -1
}

type AVL struct {
	root *node // 根结点
	size int   // 当前AVL树中的结点数量
}

func (a *AVL) Size() int {
	return a.size
}

func (a *AVL) Empty() bool {
	return a.size == 0
}

// 获取结点的高度
func (a *AVL) GetHeight(n *node) int {
	if n == nil {
		return 0
	}

	return n.height
}

// 中序遍历
func (a *AVL) InOrder() []string {
	return a.inOrder(a.root)
}

func (a *AVL) inOrder(n *node) []string {
	retStr := make([]string, 0)
	if n == nil {
		return retStr
	}

	retStr = append(retStr, a.inOrder(n.Left)...)
	retStr = append(retStr, n.key)
	retStr = append(retStr, a.inOrder(n.Right)...)

	return retStr
}

// 判断该树是否是二叉搜索树
// 实际上就是判断中序遍历的结果是否是升序的即可。
func (a *AVL) IsBST() bool {
	keys := a.inOrder(a.root)

	for i := 1; i < len(keys); i++ {
		if keys[i] < keys[i-1] {
			return false
		}
	}

	return true
}

// 判断该树是否是平衡的
func (a *AVL) IsBalanced() bool {
	return a.isBalanced(a.root)
}

func (a *AVL) isBalanced(n *node) bool {
	if n == nil {
		return true
	}

	balanceFactor := a.GetBalanceFactor(n)
	if math.Abs(float64(balanceFactor)) > 1.0 {
		return false
	}

	return a.isBalanced(n.Left) && a.isBalanced(n.Right)
}

// 获取结点的平衡因子
func (a *AVL) GetBalanceFactor(n *node) int {
	return a.GetHeight(n.Left) - a.GetHeight(n.Right)
}

// 获取元素
func (a *AVL) Get(key string) string {
	retNode := a.getNode(a.root, key)
	if retNode == nil {
		return ""
	}

	return retNode.value
}

func (a *AVL) contains(key string) bool {
	return a.getNode(a.root, key) != nil
}

// 在以n为根结点的树中搜索key
func (a *AVL) getNode(n *node, key string) *node {
	if n == nil {
		return nil
	}

	if n.key == key {
		return n
	} else if n.key > key {
		return a.getNode(n.Left, key)
	} else {
		return a.getNode(n.Right, key)
	}
}

// 添加
func (a *AVL) Add(key string, value string) {
	a.root = a.add(a.root, key, value)
}

// 在以node为根结点的avl树中添加结点，返回添加后的树的根结点
func (a *AVL) add(n *node, key string, value string) *node {
	if n == nil {
		a.size++
		return NewNode(key, value)
	}
	if n.key > key {
		n.Left = a.add(n.Left, key, value)
	} else if n.key < key {
		n.Right = a.add(n.Right, key, value)
	} else { // TODO(tfbrother) 相等时更新value值，此时可以直接返回，因为并没有影响高度和平衡因子，即使不返回也没有bug。
		n.value = value
		return n
	}

	// 更新高度：
	l := a.GetHeight(n.Left)
	r := a.GetHeight(n.Right)
	n.height = 1 + int(math.Max(float64(l), float64(r)))

	// 获取平衡因子
	balanceFactor := a.GetBalanceFactor(n)
	if math.Abs(float64(balanceFactor)) > 1.0 {
		fmt.Println(n.key, "==balanceFactor==", balanceFactor)
	}

	return n
}

// 对节点n进行向右旋转操作，返回旋转后新的根节点x
//        y                              x
//       / \                           /   \
//      x   T4     向右旋转 (y)        z     y
//     / \       - - - - - - - ->    / \   / \
//    z   T3                       T1  T2 T3 T4
//   / \
// T1   T2
func (a *AVL) rightRotate(y *node) *node {
	// 先把要移动的结点暂存下来
	x := y.Left
	t3 := x.Right
	// 向右旋转
	x.Right = y
	y.Left = t3

	// 更新高度Y，X的高度
	y.height = int(math.Max(float64(a.GetHeight(y.Left)), float64(a.GetHeight(y.Right)))) + 1
	x.height = int(math.Max(float64(a.GetHeight(x.Left)), float64(a.GetHeight(x.Right)))) + 1
	return x
}

// 对节点y进行向左旋转操作，返回旋转后新的根节点x
//    y                             x
//  /  \                          /   \
// T1   x      向左旋转 (y)       y     z
//     / \   - - - - - - - ->   / \   / \
//   T2  z                     T1 T2 T3 T4
//      / \
//     T3 T4
func (a *AVL) leftRotate(y *node) *node {
	// 先把要移动的结点暂存下来
	x := y.Right
	t3 := x.Left
	// 向右旋转
	x.Left = y
	y.Right = t3

	// 更新高度Y，X的高度
	y.height = int(math.Max(float64(a.GetHeight(y.Left)), float64(a.GetHeight(y.Right)))) + 1
	x.height = int(math.Max(float64(a.GetHeight(x.Left)), float64(a.GetHeight(x.Right)))) + 1
	return x
}

func NewNode(key string, value string) *node {
	return &node{key: key, value: value, height: 1, Left: nil, Right: nil}
}

func NewAVL() *AVL {
	return &AVL{
		root: nil,
		size: 0,
	}
}
