package avl_tree

import (
	"math"
)

// avl树一种平衡二叉搜索树，所以都是基于二叉搜索树来实现。
type Item interface {
	Less(a Item) bool
	ToString() string
}

// 设置成为私有结构体，不对外部暴露node的结构
type node struct {
	item        Item  //结点中保存的数据项目，采用类似泛型的思想，方便复用。
	height      int   // 高度
	Left, Right *node // 左右子树
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
func (a *AVL) InOrder() []Item {
	return a.inOrder(a.root)
}

func (a *AVL) inOrder(n *node) []Item {
	retStr := make([]Item, 0)
	if n == nil {
		return retStr
	}

	retStr = append(retStr, a.inOrder(n.Left)...)
	retStr = append(retStr, n.item)
	retStr = append(retStr, a.inOrder(n.Right)...)

	return retStr
}

// 判断该树是否是二叉搜索树
// 实际上就是判断中序遍历的结果是否是升序的即可。
func (a *AVL) IsBST() bool {
	keys := a.inOrder(a.root)

	for i := 1; i < len(keys); i++ {
		if keys[i].Less(keys[i-1]) {
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
func (a *AVL) Get(key Item) Item {
	retNode := a.getNode(a.root, key)
	if retNode == nil {
		return nil
	}

	return retNode.item
}

func (a *AVL) contains(key Item) bool {
	return a.getNode(a.root, key) != nil
}

// 在以n为根结点的树中搜索key
// TODO 这种辅助元素的设计，非常重要
func (a *AVL) getNode(n *node, key Item) *node {
	for n != nil {
		switch {
		case key.Less(n.item):
			return a.getNode(n.Left, key)
		case n.item.Less(key):
			return a.getNode(n.Right, key)
		default:
			return n
		}
	}

	return n
}

// 添加
func (a *AVL) Add(key Item) {
	a.root = a.add(a.root, key)
}

// 在以node为根结点的avl树中添加结点，返回添加后的树的根结点
func (a *AVL) add(n *node, key Item) *node {
	if n == nil {
		a.size++
		return NewNode(key)
	}

	if key.Less(n.item) {
		n.Left = a.add(n.Left, key)
	} else if n.item.Less(key) {
		n.Right = a.add(n.Right, key)
	} else { // TODO(tfbrother) 相等时更新，此时可以直接返回，因为并没有影响高度和平衡因子，即使不返回也没有bug。
		n.item = key
		return n
	}

	// 更新高度：
	l := a.GetHeight(n.Left)
	r := a.GetHeight(n.Right)
	oldHeight := n.height
	n.height = 1 + int(math.Max(float64(l), float64(r)))
	if oldHeight == n.height { // TODO 优化：高度没有变化就不需要平衡了。
		return n
	}

	// 获取平衡因子
	balanceFactor := a.GetBalanceFactor(n)
	//if math.Abs(float64(balanceFactor)) > 1.0 {
	//	fmt.Println(n.item.ToString(), "==balanceFactor==", balanceFactor)
	//}

	// LL
	if balanceFactor > 1 && a.GetBalanceFactor(n.Left) >= 0 {
		return a.rightRotate(n)
	}

	// RR
	if balanceFactor < -1 && a.GetBalanceFactor(n.Right) <= 0 {
		return a.leftRotate(n)
	}

	// LR
	if balanceFactor > 1 && a.GetBalanceFactor(n.Left) < 0 {
		n.Left = a.leftRotate(n.Left)
		return a.rightRotate(n)
	}

	// RL
	if balanceFactor < -1 && a.GetBalanceFactor(n.Right) > 0 {
		n.Right = a.rightRotate(n.Right)
		return a.leftRotate(n)
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

func NewNode(key Item) *node {
	return &node{item: key, height: 1, Left: nil, Right: nil}
}

func NewAVL() *AVL {
	return &AVL{
		root: nil,
		size: 0,
	}
}
