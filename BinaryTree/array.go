// 二叉树，采用数组来存储
package BinaryTree

import "log"

/*
				5(0)					左子节点索引：2*X+1
		7(1)			8(2)			右子节点索引：2*X+2
	8(3)	10(4)	3(5)	2(6)
*/

type Node int

type Tree struct {
	size  int    //树的大小
	nodes []Node //用数组存放树所有的结点数据
}

// 添加子节点
// 给nodeIndex结点添加子结点（左/右）
func (t *Tree) AddNode(nodeIndex int, direction int, node Node) bool {
	if nodeIndex < 0 || nodeIndex > t.size-1 { //越界了
		return false
	}

	var insertIndex int
	switch direction {
	case 0: //左
		insertIndex = 2*nodeIndex + 1
	case 1: //右
		insertIndex = 2*nodeIndex + 2
	default:
		return false
	}

	if insertIndex < 0 || insertIndex > t.size-1 { //越界了
		return false
	}

	t.nodes[insertIndex] = node
	return true
}

func (t *Tree) Dump() {
	log.Println(t.nodes)
}

func NewTree(size int) *Tree {
	return &Tree{
		size:  size,
		nodes: make([]Node, size),
	}
}
