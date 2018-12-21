package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph/mst"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph/traversal"
	"os"
)

func main() {
	f, err := os.Open("/Users/tfbrother/go/src/github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph/testdata/graph.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	g, err := graph.NewGraphFromJSON(f, "graph_00")
	if err != nil {
		fmt.Println(err)
	}
	rs := traversal.DFSRecursion(g, graph.StringID("S"))
	fmt.Println("DFSRecursion:", rs) // [S C E T A B D F]
	if len(rs) != 8 {
		fmt.Printf("should be 8 vertices but %s\n", g)
	}

	f.Seek(0, 0)
	g, err = graph.NewGraphFromJSON(f, "graph_13")
	if err != nil {
		fmt.Println(err)
	}
	A := mst.Kruskal(g)

	total := 0.0
	for _, edge := range A {
		total += edge.Weight()
	}
	if total != 37.0 {
		fmt.Println("Expected total 37.0 but %.2f", total)
	}
	fmt.Println("Kruskal from graph_13:", A)

	fmt.Println("=========start lazyPrim from graph_13:", A)
	for v := range g.GetNodes() {
		l := mst.NewLazyPrim(g)
		A := l.MstEdges(v)

		total := 0.0
		for _, edge := range A {
			total += edge.Weight()
		}
		if total != 37.0 {
			fmt.Println("Expected total 37.0 but %.2f", total)
		}
		fmt.Println("===========end lazyPrim from graph_13:", A, "with", v)
	}

	fmt.Println("=========start Prim from graph_13:", A)
	for v := range g.GetNodes() {
		l := mst.NewPrim(g)
		A := l.MstEdges(v)

		total := 0.0
		for _, edge := range A {
			total += edge.Weight()
		}
		if total != 37.0 {
			fmt.Println("Expected total 37.0 but %.2f", total)
		}
		fmt.Println("===========end Prim from graph_13:", A, "with", v)
	}
}
