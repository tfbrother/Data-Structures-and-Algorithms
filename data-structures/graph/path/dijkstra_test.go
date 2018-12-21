package path

import (
	"fmt"
	. "github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph"
	"os"
	"strings"
	"testing"
)

func TestGraph_Dijkstra_03(t *testing.T) {
	f, err := os.Open("../testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_03")
	if err != nil {
		t.Error(err)
	}
	path, distance := Dijkstra(g, StringID("S"), StringID("T"))

	ts := []string{}
	for i := len(path) - 1; i >= 0; i-- {
		ts = append(ts, fmt.Sprintf("%s(%.2f)", path[i], distance[path[i]]))
	}
	if strings.Join(ts, " → ") != "S(0.00) → B(14.00) → E(32.00) → F(38.00) → T(44.00)" {
		t.Errorf("Expected the shortest path S(0.00) → B(14.00) → E(32.00) → F(38.00) → T(44.00) but %s", strings.Join(ts, " → "))
	}
	if distance[StringID("T")] != 44.0 {
		t.Errorf("Expected 44.0 but %f", distance[StringID("T")])
	}
	fmt.Println("graph_03:", strings.Join(ts, " → "))
}
