package graph_test

import (
	"fmt"
	. "github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph"
	"log"
	"os"
	"testing"
)

func TestNewGraphFromJSON(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_00")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g.String())
}
