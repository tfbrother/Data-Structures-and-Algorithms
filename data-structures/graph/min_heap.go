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
type nodeDistance struct {
	nd    Node
	e     Edge
	index int
}

type nodeDistaceHeap []nodeDistance

func (d nodeDistaceHeap) Len() int { return len(d) }
func (d nodeDistaceHeap) Less(i, j int) bool {
	return d[i].e.Weight() < d[j].e.Weight()
}

func (d nodeDistaceHeap) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
	d[i].index = i
	d[j].index = j

}

func (d *nodeDistaceHeap) Pop() interface{} {
	heapLen := len(*d)
	lastNode := (*d)[heapLen-1]
	*d = (*d)[:heapLen-1]
	lastNode.index = -1
	return lastNode
}

func (d *nodeDistaceHeap) Push(x interface{}) {
	nd := x.(nodeDistance)
	nd.index = len(*d)
	*d = append(*d, nd)
}

// TODO need improve ,i.e. add a indexOf item(type map[ID]int), can directed find the index by give nd
func (d *nodeDistaceHeap) update(nd Node, e Edge) {
	for i := 0; i < len(*d); i++ {
		if (*d)[i].nd.ID() == nd.ID() {
			(*d)[i].e = e
			heap.Fix(d, (*d)[i].index)
			break
		}
	}
}
