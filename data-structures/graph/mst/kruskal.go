// minimum spanning tree
package mst

import (
	. "github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph"
	"sort"
)

// Kruskal finds the minimum spanning tree with disjoint-set data structure
// (http://en.wikipedia.org/wiki/Kruskal%27s_algorithm)
//
//	 0. Kruskal(G)
//	 1.
//	 2. 	A = ∅
//	 3.
//	 4. 	for each vertex v in G:
//	 5. 		MakeDisjointSet(v)
//	 6.
//	 7. 	edges = get all edges
//	 8. 	sort edges in ascending order of weight
//	 9.
//	10. 	for each edge (u, v) in edges:
//	11. 		if FindSet(u) ≠ FindSet(v):
//	12. 			A = A ∪ {(u, v)}
//	13. 			Union(u, v)
//	14.
//	15. 	return A
//
func Kruskal(g Graph) []Edge {
	ret := make([]Edge, 0, g.GetNodeCount()-1)

	// get all vertex in g
	nodes := g.GetNodes()

	disSet := NewDisjointSet(len(nodes))
	// push all vertex in disjoint-set
	for id, _ := range nodes {
		disSet.AddElement(id.String())
	}

	// get all edges
	edges := g.GetAllEdges()
	// sort edges in ascending order of weight
	sort.Sort(EdgeSlice(edges))

	for _, e := range edges {
		if !disSet.IsConnected(e.Source().String(), e.Target().String()) {
			ret = append(ret, e)
			disSet.Union(e.Source().String(), e.Target().String())
		}
	}

	return ret
}
