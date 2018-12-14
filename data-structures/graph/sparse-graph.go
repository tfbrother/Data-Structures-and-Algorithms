package graph

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/queue"
)

// 稀疏图，用领接表存储
type SparseGraph struct {
	nodeCount int            // 结点数量
	edgeCount int            // 边数量
	edges     []map[int]bool // 保存边数据
	nodes     []int          // 保存所有的顶点
	directed  bool           // 是否是有向图
	visited   []bool         // 记录结点是否被访问过，用于遍历时使用
}

func (s *SparseGraph) NodeCount() int {
	return s.nodeCount
}

func (s *SparseGraph) EdgecCount() int {
	return s.edgeCount
}

// 添加边
func (s *SparseGraph) AddEdge(v int, w int) bool {
	if (s.hasEdge(v, w) || v == w) || (v < 0 || v >= s.nodeCount || w < 0 || w >= s.nodeCount) {
		return false
	}

	s.edges[v][w] = true
	if !s.directed {
		s.edges[w][v] = true
	}
	s.edgeCount++
	return true
}

// 判断v，w两个顶点之间是否有边
func (s *SparseGraph) hasEdge(v int, w int) bool {
	return s.edges[v][w]
}

// 从v结点开始进行深度优先遍历
func (d *SparseGraph) DepthFirstTraverse(v int) {
	if v < 0 || v >= d.nodeCount {
		return
	}

	d.visited[v] = true // 入队列时就设置为已经访问过了
	fmt.Print(v, " ")
	for i, value := range d.edges[v] {
		// 找到与V相邻且未被访问的第一个结点就开始递归进行深度访问
		if value && d.visited[i] != true {
			d.DepthFirstTraverse(i)
		}
	}
}

// 思路：借助队列实现，先将结点放入队列，然后出队列时把和他所有的相邻且未访问的结点都放入队列
func (d *SparseGraph) BreadthFirstTraverse(v int) {
	if v < 0 || v >= d.nodeCount {
		return
	}
	q := queue.NewQueue(10)
	q.Push(v)
	fmt.Print(v, " ")
	d.visited[v] = true // 入队列时就设置为已经访问过了

	for !q.Empty() {
		v = q.Pop().(int)
		for i, value := range d.edges[v] {
			// 将所有与V相邻且未访问的结点都放入队列
			if value && d.visited[i] != true {
				q.Push(i)
				fmt.Print(i, " ")
				d.visited[i] = true
			}
		}
	}
}

// 重置所有的访问状态，用于测试用
func (d *SparseGraph) ResetStatus() {
	d.visited = make([]bool, d.nodeCount)
}

func NewSparseGraph(count int, directed bool) *SparseGraph {
	s := SparseGraph{
		nodeCount: count,
		edgeCount: 0,
		nodes:     make([]int, count),
		edges:     make([]map[int]bool, count),
		directed:  directed,
		visited:   make([]bool, count),
	}

	for i := 0; i < count; i++ {
		s.nodes[i] = i
		s.edges[i] = make(map[int]bool)
	}

	return &s
}
