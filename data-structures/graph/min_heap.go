package graph

import (
	"container/heap"
)

// min-heap
// implemented heap.Interface
type EdgeSlice []Edge

func (e EdgeSlice) Len() int { return len(e) }
func (e EdgeSlice) Less(i, j int) bool {
	return e[i].Weight() < e[j].Weight()
}
func (e EdgeSlice) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e *EdgeSlice) Pop() interface{} {
	ret := (*e)[len(*e)-1]
	*e = (*e)[:len(*e)-1]
	return ret
}

func (e *EdgeSlice) Push(x interface{}) {
	d := x.(Edge)
	*e = append(*e, d)
}

// min-index heap
type NodeDistance struct {
	Nd    Node
	E     Edge
	index int
}

type NodeDistaceHeap []NodeDistance

func (d NodeDistaceHeap) Len() int { return len(d) }
func (d NodeDistaceHeap) Less(i, j int) bool {
	return d[i].E.Weight() < d[j].E.Weight()
}

func (d NodeDistaceHeap) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
	d[i].index = i
	d[j].index = j

}

func (d *NodeDistaceHeap) Pop() interface{} {
	heapLen := len(*d)
	lastNode := (*d)[heapLen-1]
	*d = (*d)[:heapLen-1]
	lastNode.index = -1
	return lastNode
}

func (d *NodeDistaceHeap) Push(x interface{}) {
	nd := x.(NodeDistance)
	nd.index = len(*d)
	*d = append(*d, nd)
}

// TODO need improve ,i.e. add a indexOf item(type map[ID]int), can directed find the index by give nd
func (d *NodeDistaceHeap) Update(nd Node, e Edge) {
	for i := 0; i < len(*d); i++ {
		if (*d)[i].Nd.ID() == nd.ID() {
			(*d)[i].E = e
			heap.Fix(d, (*d)[i].index)
			break
		}
	}
}
