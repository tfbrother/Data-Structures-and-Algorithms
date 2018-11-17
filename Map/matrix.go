package Map

import "fmt"

// 图结点类型
type Node struct {
	value    string // 顶点的值
	isAccess bool   // 顶点是否被访问
}

func NewNode(value string) Node {
	return Node{value: value}
}

//采用邻接矩阵存储map
type Gmap struct {
	nodes     []Node // 数组存储结点列表
	matrix    []int  // 数组存储边数据矩阵
	capacity  int    // 图的容量
	nodeCount int    // 已添加顶点的个数
}

// 添加顶点
func (g *Gmap) AddNode(node Node) bool {
	if g.nodeCount >= g.capacity {
		return false
	}

	g.nodes = append(g.nodes, node)
	g.nodeCount++
	return true
}

// 重置所有的顶点到初始状态
func (g *Gmap) ResetNode() {
	for i := 0; i < g.nodeCount; i++ {
		g.nodes[i].isAccess = false
	}
}

//设置无向图的边
func (g *Gmap) SetUnGraphValue(row int, col int, weight int) bool {
	if row < 0 || row >= g.capacity || col < 0 || col >= g.capacity {
		return false
	}

	g.matrix[row*g.capacity+col] = weight
	g.matrix[col*g.capacity+row] = weight
	return true
}

// 输出矩阵
func (g *Gmap) Dump() {
	for i := 0; i < g.capacity; i++ {
		for j := 0; j < g.capacity; j++ {
			fmt.Print(g.matrix[i*g.capacity+j], "  ")
		}
		fmt.Println()
	}
}

// 深度优先遍历
func (g *Gmap) DepthFirstTraverse(startIndex int) {
	if startIndex < 0 || startIndex > g.capacity {
		return
	}

	if g.nodeCount != g.capacity { // 图并不完整
		return
	}

	fmt.Print(g.nodes[startIndex].value, "(", startIndex, ")")
	//设置该顶点已经访问过了
	g.nodes[startIndex].isAccess = true
	for i := 0; i < g.capacity; i++ {
		if g.matrix[i*g.capacity+startIndex] != 0 && g.nodes[i].isAccess != true { // 找到一个满足条件第一个边，然后递归搜索下去
			g.DepthFirstTraverse(i)
		}
	}

	return
}

// 广度优先遍历
func (g *Gmap) BreadthFirstTraverse(startIndex int) {

}

// 初始化一个gmap
func NewGmap(capacity int) *Gmap {
	return &Gmap{
		nodes:     make([]Node, capacity),
		matrix:    make([]int, capacity*capacity),
		capacity:  capacity,
		nodeCount: 0,
	}
}
