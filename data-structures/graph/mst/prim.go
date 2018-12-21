package mst

import (
	"container/heap"
	. "github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph"
)

// need min index heap(priority queue)
type prim struct {
	g         Graph
	pq        *NodeDistaceHeap
	marked    map[ID]bool
	mstWeight float64
	mst       []Edge
	heapEdge  map[ID]Edge //访问的点所对应的边, 算法辅助数据结构
}

// 区别与lazy prim算法，该算法只把与顶点nd相连的最小边放入索引堆
func (l *prim) Visit(nd Node) {
	l.marked[nd.ID()] = true
	edges := l.g.GetNodeEdges(nd.ID())

	for _, e := range edges {
		// 如果边的另一端点未被访问
		o := e.Other(nd)
		if !l.marked[o.ID()] {
			// 如果从没有考虑过这个端点, 直接将这个端点和与之相连接的边加入索引堆
			if _, ok := l.heapEdge[o.ID()]; !ok {
				l.heapEdge[o.ID()] = e
				heap.Push(l.pq, NodeDistance{Nd: o, E: e})
			} else if e.Weight() < l.heapEdge[o.ID()].Weight() { // 如果曾经考虑这个端点, 但现在的边比之前考虑的边更短, 则进行替换
				l.heapEdge[o.ID()] = e
				l.pq.Update(o, e)
			}
		}
	}
}

func (l *prim) primMst(src ID) {
	node, err := l.g.GetNode(src)
	if err != nil {
		return
	}

	// 算法初始化
	for id, _ := range l.g.GetNodes() {
		l.marked[id] = false
	}

	l.mst = make([]Edge, 0, l.g.GetNodeCount()-1)
	l.Visit(node)

	for l.pq.Len() > 0 {
		// 使用最小索引堆找出已经访问的边中权值最小的边
		// 最小索引堆中存储的是点的索引, 通过点的索引找到相对应的边
		d := heap.Pop(l.pq).(NodeDistance)
		l.mst = append(l.mst, d.E)
		l.Visit(d.Nd)
	}

	l.mstWeight = l.mst[0].Weight()
	for i := 1; i < len(l.mst); i++ {
		l.mstWeight += l.mst[i].Weight()
	}
	return
}

func (l *prim) MstEdges(src ID) []Edge {
	l.primMst(src)
	return l.mst
}

func NewPrim(g Graph) *prim {
	l := &prim{
		g:        g,
		pq:       &NodeDistaceHeap{},
		marked:   make(map[ID]bool),
		mst:      make([]Edge, g.GetNodeCount()-1),
		heapEdge: make(map[ID]Edge),
	}
	return l
}
