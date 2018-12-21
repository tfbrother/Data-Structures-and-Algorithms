package path

import (
	"container/heap"
	. "github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph"
	"math"
)

// Dijkstra returns the shortest path using Dijkstra
// algorithm with a min-priority queue. This algorithm
// does not work with negative weight edges.
// (https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)
//
//	 0. Dijkstra(G, source, target)
//	 1.
//	 2. 	let Q be a priority queue
//	 3. 	distance[source] = 0
//	 4.
//	 5. 	for each vertex v in G:
//	 6.
//	 7. 		if v ≠ source:
//	 8. 			distance[v] = ∞
//	 9. 			prev[v] = undefined
//	10.
//	11. 		Q.add_with_priority(v, distance[v])
//	12.
//	13. 	while Q is not empty:
//	14.
//	15. 		u = Q.extract_min()
//	16. 		if u == target:
//	17. 			break
//	18.
//	19. 		for each child vertex v of u:
//	20.
//	21. 			alt = distance[u] + weight(u, v)
//	22. 			if distance[v] > alt:
//	23. 				distance[v] = alt
//	24. 				prev[v] = u
//	25. 				Q.decrease_priority(v, alt)
//	26.
//	27. 		reheapify(Q)
//	28.
//	29.
//	30. 	path = []
//	31. 	u = target
//	32. 	while prev[u] is defined:
//	33. 		path.push_front(u)
//	34. 		u = prev[u]
//	35.
//	36. 	return path, prev
//
func Dijkstra(g Graph, src ID, tgt ID) ([]ID, map[ID]float64) {
	minHeap := &dijQueue{}
	distance := make(map[ID]float64)
	from := make(map[ID]ID)
	distance[src] = 0.0

	nodes := g.GetNodes()

	for id, nd := range nodes {
		if id != src {
			distance[id] = math.MaxFloat64
		}
		from[id] = nil
		minHeap.Push(Distance{nd, distance[id], 0})
		//*minHeap = append(*minHeap, Distance{nd, distance[id], 0})
	}

	// heapify
	heap.Init(minHeap)

	for len(*minHeap) > 0 {
		u := heap.Pop(minHeap).(Distance)
		if u.nd.ID() == tgt {
			break
		}

		edges := g.GetNodeEdges(u.nd.ID())

		for _, e := range edges {
			alt := distance[u.nd.ID()] + e.Weight()
			v := e.Other(u.nd)
			if distance[v.ID()] > alt {
				distance[v.ID()] = alt
				from[v.ID()] = u.nd.ID()
				minHeap.Update(v, alt)
			}
		}

		heap.Init(minHeap)
	}

	shortPath := make([]ID, 0, len(from))
	w := tgt
	for from[w] != nil {
		// path.push_front(u)
		shortPath = append(shortPath, w)
		w = from[w]
	}

	shortPath = append(shortPath, src)
	return shortPath, distance
}

type Distance struct {
	nd       Node
	distance float64
	index    int
}

type dijQueue []Distance

func (d dijQueue) Len() int { return len(d) }
func (d dijQueue) Less(i, j int) bool {
	return d[i].distance < d[j].distance
}

func (d dijQueue) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
	d[i].index = i
	d[j].index = j

}

func (d *dijQueue) Pop() interface{} {
	heapLen := len(*d)
	lastNode := (*d)[heapLen-1]
	*d = (*d)[:heapLen-1]
	lastNode.index = -1
	return lastNode
}

func (d *dijQueue) Push(x interface{}) {
	nd := x.(Distance)
	nd.index = len(*d)
	*d = append(*d, nd)
}

// TODO need improve ,i.e. add a indexOf item(type map[ID]int), can directed find the index by give nd
func (d *dijQueue) Update(nd Node, dis float64) {
	for i := 0; i < len(*d); i++ {
		if (*d)[i].nd.ID() == nd.ID() {
			(*d)[i].distance = dis
		}
	}

}
