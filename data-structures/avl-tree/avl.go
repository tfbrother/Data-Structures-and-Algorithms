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

// 获取结点的平衡因子
func (a *AVL) GetBalanceFactor(n *node) int {
	return a.GetHeight(n.Left) - a.GetHeight(n.Right)
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

func NewNode(key string, value string) *node {
	return &node{key: key, value: value, height: 1, Left: nil, Right: nil}
}

func NewAVL() *AVL {
	return &AVL{
		root: nil,
		size: 0,
	}
}