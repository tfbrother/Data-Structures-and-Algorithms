package graph

import (
	"container/heap"
)

type lazyPrim struct {
	g         Graph
	pq        *EdgeSlice
	marked    map[ID]bool
	mstweight float64
	mst       []Edge
}

func (l *lazyPrim) Visit(nd Node) {
	l.marked[nd.ID()] = true
	edges := l.g.GetNodeEdges(nd.ID())

	for _, e := range edges {

		if !l.marked[e.Source().ID()] && e.Source().ID() != nd.ID() {
			heap.Push(l.pq, e)
		} else if !l.marked[e.Target().ID()] && e.Target().ID() != nd.ID() {
			heap.Push(l.pq, e)
		}
	}
}

func (l *lazyPrim) LazyPrim(src ID) {
	node, err := l.g.GetNode(src)
	if err != nil {
		return
	}
	for id, _ := range l.g.GetNodes() {
		l.marked[id] = false
	}

	l.mst = make([]Edge, 0, l.g.GetNodeCount()-1)
	l.Visit(node)

	for l.pq.Len() > 0 {
		e := heap.Pop(l.pq).(Edge)
		if l.marked[e.Target().ID()] == l.marked[e.Source().ID()] {
			continue
		}

		l.mst = append(l.mst, e)
		if !l.marked[e.Target().ID()] {
			l.Visit(e.Target())
		} else {
			l.Visit(e.Source())
		}
	}
	return
	l.mstweight = l.mst[0].Weight()
	for i := 1; i < len(l.mst); i++ {
		l.mstweight += l.mst[i].Weight()
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
