package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph/mst"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph/path"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph/traversal"
	"os"
	"strings"
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

	f.Seek(0, 0)
	g, err = graph.NewGraphFromJSON(f, "graph_03")
	if err != nil {
		fmt.Println(err)
	}
	path, distance := path.Dijkstra(g, graph.StringID("S"), graph.StringID("T"))
	fmt.Println(path, distance)
	ts := []string{}
	for i := len(path) - 1; i >= 0; i-- {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", path[i], distance[path[i]]))
	}

	if strings.Join(ts, " → ") != "S(0.00) → B(14.00) → E(32.00) → F(38.00) → T(44.00)" {
		fmt.Printf("Expected the shortest path S(0.00) → B(14.00) → E(32.00) → F(38.00) → T(44.00) but %s", strings.Join(ts, " → "))
	}
	if distance[graph.StringID("T")] != 44.0 {
		fmt.Printf("Expected 44.0 but %f", distance[graph.StringID("T")])
	}
	fmt.Println("graph_03:", strings.Join(ts, " → "))
}
