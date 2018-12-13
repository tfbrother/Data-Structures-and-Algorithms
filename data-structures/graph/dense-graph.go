package graph

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/queue"
)

// TODO 先实现功能，后期借鉴数据库驱动的设计，提供统一的接口和迭代器等等，对外只需要传一个参数，确定用稠密图还是稀疏图。
// 稠密图，用领接矩阵存储
type DenseGraph struct {
	nodeCount int      // 结点数量
	edgeCount int      // 边数量
	edges     [][]bool // 保存边数据
	nodes     []int    // 保存所有的顶点
	directed  bool     // 是否是有向图
	visited   []bool   // 记录结点是否被访问过，用于遍历时使用
}

// 返回结点数量
func (d *DenseGraph) NodeCount() int {
	return d.nodeCount
}

// 返回边的数量
func (d *DenseGraph) EdgeCount() int {
	return d.edgeCount
}

// 添加边
func (d *DenseGraph) AddEdge(v int, w int) bool {
	if (d.hasEdge(v, w) || v == w) || (v < 0 || v >= d.nodeCount || w < 0 || w >= d.nodeCount) {
		return false
	}

	d.edges[v][w] = true
	if !d.directed {
		d.edges[w][v] = true
	}
	d.edgeCount++
	return true
}

// 判断v，w两个顶点之间是否有边
func (d *DenseGraph) hasEdge(v int, w int) bool {
	return d.edges[v][w]
}

// 从v结点开始进行深度优先遍历
// TODO 在应用中，一般遍历图，实际上是要对每个结点做一些操作，就类似从数据库中查询数据，要在应用层遍历每条数据用于后面的业务逻辑
// 所以远不止这里实现的这样，只是简单的全部输出，没法对单个结点做后序逻辑，这也是很多实际场景中，数据结构的遍历很少用，搜索却很常用。
// O(n^2)
func (d *DenseGraph) DepthFirstTraverse(v int) {
	if v < 0 || v >= d.nodeCount {
		return
	}

	d.visited[v] = true // 入队列时就设置为已经访问过了
	fmt.Print(v, " ")
	for i := 0; i < len(d.edges[v]); i++ {
		// 找到与V相邻且未被访问的第一个结点就开始递归进行深度访问
		if d.edges[v][i] && d.visited[i] != true {
			d.DepthFirstTraverse(i)
		}
	}
}

// 思路：借助队列实现，先将结点放入队列，然后出队列时把和他所有的相邻且未访问的结点都放入队列
// O(n^2)
func (d *DenseGraph) BreadthFirstTraverse(v int) {
	if v < 0 || v >= d.nodeCount {
		return
	}
	q := queue.NewQueue(10)
	q.Push(v)
	fmt.Print(v, " ")
	d.visited[v] = true // 入队列时就设置为已经访问过了

	for !q.Empty() {
		v = q.Pop().(int)
		for i := 0; i < len(d.edges[v]); i++ {
			// 将所有与V相邻且未访问的结点都放入队列
			if d.edges[v][i] && d.visited[i] != true {
				q.Push(i)
				fmt.Print(i, " ")
				d.visited[i] = true
			}
		}
	}
}

// 重置所有的访问状态，用于测试用
func (d *DenseGraph) ResetVisted() {
	d.visited = make([]bool, d.nodeCount)
}

func NewDenseGraph(count int, directed bool) *DenseGraph {
	d := DenseGraph{
		nodeCount: count,
		edgeCount: 0,
		nodes:     make([]int, count),
		edges:     make([][]bool, count),
		directed:  directed,
		visited:   make([]bool, count),
	}

	for i := 0; i < count; i++ {
		d.nodes[i] = i
		d.edges[i] = make([]bool, count)
	}

	return &d
}
