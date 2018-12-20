package graph

import (
	"container/heap"
)

type lazyPrim struct {
	g         Graph
	pq        *EdgeSlice
	marked    map[ID]bool
	mstWeight float64
	mst       []Edge
}

func (l *lazyPrim) Visit(nd Node) {
	l.marked[nd.ID()] = true
	edges := l.g.GetNodeEdges(nd.ID())

	for _, e := range edges {
		// 将和节点nd相连接的所有未访问的边放入最小堆中
		if !l.marked[e.Other(nd)] {
			heap.Push(l.pq, e)
		}
	}
}

func (l *lazyPrim) LazyPrim(src ID) {
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
		// 使用最小堆找出已经访问的边中权值最小的边
		e := heap.Pop(l.pq).(Edge)
		// 如果这条边的两端都已经访问过了, 则扔掉这条边
		if l.marked[e.Target().ID()] == l.marked[e.Source().ID()] {
			continue
		}

		// 否则, 这条边则应该存在在最小生成树中
		l.mst = append(l.mst, e)

		// 访问和这条边连接的还没有被访问过的节点
		if !l.marked[e.Target().ID()] {
			l.Visit(e.Target())
		} else {
			l.Visit(e.Source())
		}
	}

	l.mstWeight = l.mst[0].Weight()
	for i := 1; i < len(l.mst); i++ {
		l.mstWeight += l.mst[i].Weight()
	}
	return
}

func (l *lazyPrim) MstEdges(src ID) []Edge {
	l.LazyPrim(src)
	return l.mst
}

func NewLazyPrim(g Graph) *lazyPrim {
	l := &lazyPrim{
		g:      g,
		pq:     &EdgeSlice{},
		marked: make(map[ID]bool),
		mst:    make([]Edge, g.GetNodeCount()-1),
	}
	return l
}
