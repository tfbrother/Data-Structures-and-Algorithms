package graph_test

import (
	"fmt"
	. "github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph"
	"math/rand"
	"testing"
	"time"
)

// 广度优先遍历测试
func TestDenseGraph_BreadthFirstTraverse1K(t *testing.T) {
	n := 1 << 5
	d := NewDenseGraph(n, false)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		d.AddEdge(a, b)
	}
	d.DepthFirstTraverse(10)
	fmt.Println()
	d.ResetVisted()
	d.BreadthFirstTraverse(10)
	fmt.Println()
}
