package traversal

import (
	. "github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph"
)

// BFS does breadth-first search, and returns the list of vertices.
// (https://en.wikipedia.org/wiki/Breadth-first_search)
// Breadth first traversal is accomplished by enqueueing each level of a tree sequentially
// as the root of any subtree is encountered. There are 2 cases in the iterative algorithm.
//
// 	1.Root case: The traversal queue is initially empty so
//    the root node must be added before the general case.
// 	2.General case: Process any items in the queue, while also expanding
// 	  their children. Stop if the queue is empty. The general case will halt
//    after processing the bottom level as leaf nodes have no children.
func BFS(g Graph, id ID) []ID {
	if _, err := g.GetNode(id); err != nil {
		return nil
	}

	// push queue
	queue := []ID{id}
	visited := make(map[ID]bool)
	visited[id] = true
	ret := []ID{id}

	for len(queue) > 0 {
		c := queue[0] // queue.Pop()
		queue = queue[1:]

		// for each vertex w adjacent to c:
		if tmap, err := g.GetTargets(c); err == nil {
			for _, w := range tmap {
				if _, ok := visited[w.ID()]; !ok {
					queue = append(queue, w.ID()) // queue.Push(w)
					ret = append(ret, w.ID())
					visited[w.ID()] = true
				}
			}
		}

		if smap, err := g.GetSources(c); err == nil {
			for _, w := range smap {
				if _, ok := visited[w.ID()]; !ok {
					queue = append(queue, w.ID())
					ret = append(ret, w.ID())
					visited[w.ID()] = true
				}
			}
		}
	}

	return ret
}

// DFS does depth-first search, and returns the list of vertices.
// (https://en.wikipedia.org/wiki/Depth-first_search)
//
//1  procedure DFS-iterative(G,v):
//2      let S be a stack
//3      S.push(v)
//4      while S is not empty
//5          v = S.pop()
//6          if v is not labeled as discovered:
//7              label v as discovered
//8              for all edges from v to w in G.adjacentEdges(v) do
//9                  S.push(w)
//
func DFS(g Graph, id ID) []ID {
	if _, err := g.GetNode(id); err != nil {
		return nil
	}

	// push stack
	s := []ID{id}
	visited := make(map[ID]bool)
	ret := []ID{}

	for len(s) > 0 {
		c := s[len(s)-1] // s.Pop()
		s = s[:len(s)-1]
		if _, ok := visited[c]; !ok {
			visited[c] = true
			ret = append(ret, c)

			// for each vertex w adjacent to c:
			if tmap, err := g.GetTargets(c); err == nil {
				for _, w := range tmap {
					if _, ok := visited[w.ID()]; !ok {
						s = append(s, w.ID()) // s.Push(w)
					}
				}
			}

			if smap, err := g.GetSources(c); err == nil {
				for _, w := range smap {
					if _, ok := visited[w.ID()]; !ok {
						s = append(s, w.ID())
					}
				}
			}
		}
	}

	return ret
}

// DFSRecursion does depth-first search recursively.
//1  procedure DFS(G,v):
//2      label v as discovered
//3      for all edges from v to w in G.adjacentEdges(v) do
//4          if vertex w is not labeled as discovered then
//5              recursively call DFS(G,w)

func DFSRecursion(g Graph, id ID) []ID {
	if _, err := g.GetNode(id); err != nil {
		return nil
	}

	visited := make(map[ID]bool)
	ret := make([]ID, 0)
	// TODO see slice usage and internals https://blog.golang.org/go-slices-usage-and-internals
	// To increase the capacity of a slice one must create a new, larger slice and copy the contents of the original slice into it.
	dfsRecursion(g, id, visited, &ret)

	return ret
}

//func dfsRecursion(g Graph, id ID, visited map[ID]bool) (ret []ID) {
//	if _, ok := visited[id]; ok {
//		return
//	}
//	fmt.Printf("ret %p\n", ret)
//	fmt.Printf("&ret %p\n", &ret)
//	fmt.Printf("visited %p\n", visited)
//	fmt.Printf("&visited %p\n", &visited)
//	visited[id] = true
//	ret = append(ret, id)
//	fmt.Printf("ret %p\n", ret)
//	fmt.Printf("&ret %p\n", &ret)
//	fmt.Printf("==========\n")
//	// for each vertex w adjacent to c:
//	if tmap, err := g.GetTargets(id); err == nil {
//		for _, w := range tmap {
//			if _, ok := visited[w.ID()]; !ok {
//				ret = append(ret, dfsRecursion(g, w.ID(), visited)...)
//			}
//		}
//	}
//
//	if smap, err := g.GetSources(id); err == nil {
//		for _, w := range smap {
//			if _, ok := visited[w.ID()]; !ok {
//				ret = append(ret, dfsRecursion(g, w.ID(), visited)...)
//			}
//		}
//	}
//	return ret
//}

func dfsRecursion(g Graph, id ID, visited map[ID]bool, ret *[]ID) {
	if _, ok := visited[id]; ok {
		return
	}
	visited[id] = true
	*ret = append(*ret, id)
	// for each vertex w adjacent to c:
	if tmap, err := g.GetTargets(id); err == nil {
		for _, w := range tmap {
			if _, ok := visited[w.ID()]; !ok {
				dfsRecursion(g, w.ID(), visited, ret)
			}
		}
	}

	if smap, err := g.GetSources(id); err == nil {
		for _, w := range smap {
			if _, ok := visited[w.ID()]; !ok {
				dfsRecursion(g, w.ID(), visited, ret)
			}
		}
	}
	return
}
