package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph"
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
	rs := graph.DFSRecursion(g, graph.StringID("S"))
	fmt.Println("DFSRecursion:", rs) // [S C E T A B D F]
	if len(rs) != 8 {
		fmt.Printf("should be 8 vertices but %s\n", g)
	}
}
