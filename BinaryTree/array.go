// 二叉树，采用数组来存储
package BinaryTree

import (
	"errors"
	"log"
)

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
	if t.nodes[nodeIndex] == 0 { //值为0，约定结点就不存在
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

// 搜索
func (t *Tree) SearchNode(nodeIndex int) (node Node, err error) {
	if nodeIndex < 0 || nodeIndex > t.size-1 { //越界了
		return 0, errors.New("nodeIndex 越界了")
	}
	if t.nodes[nodeIndex] == 0 { //值为0，约定结点就不存在
		return 0, errors.New("该结点不存在")
	}
	return t.nodes[nodeIndex], nil
}

// 删除结点，同时会删除其所有字节点
func (t *Tree) DeleteNode(nodeIndex int) (node Node, err error) {
	if node, err = t.SearchNode(nodeIndex); err != nil {
		return
	}

	t.nodes[nodeIndex] = 0       // 设置为0就代表删除
	t.deleteChildNode(nodeIndex) //删除结点下面所有的子结点

	return 0, nil
}

// 递归删除结点下面所有的子结点
func (t *Tree) deleteChildNode(nodeIndex int) {
	// 依次删除左右字节点
	var lChildIndex, rChildIndex int //左右字节点的索引
	lChildIndex = 2*nodeIndex + 1
	rChildIndex = 2*nodeIndex + 2

	if lChildIndex < t.size-1 && t.nodes[lChildIndex] != 0 { //有左结点，则删除该结点和子结点
		t.nodes[lChildIndex] = 0
		t.deleteChildNode(lChildIndex)
	}

	if rChildIndex < t.size-1 && t.nodes[rChildIndex] != 0 { //有右结点，则删除该结点和子结点
		t.nodes[rChildIndex] = 0
		t.deleteChildNode(rChildIndex)
	}
}

func (t *Tree) Dump() {
	log.Println(t.nodes)
}

// 同时初始化根结点（AddNode无法初始化根结点）
func NewTree(size int, root Node) (tree *Tree) {
	tree = &Tree{
		size:  size,
		nodes: make([]Node, size),
	}

	tree.nodes[0] = root

	return
}
