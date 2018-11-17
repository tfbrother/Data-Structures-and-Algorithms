package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/Map"
)

var (
	node1, node2, node3, node4, node5, node6, node7, node8 Map.Node
	gmap                                                   *Map.Gmap
)

func main() {
	gmap = Map.NewGmap(8)
	node1 = Map.NewNode("A")
	node2 = Map.NewNode("B")
	node3 = Map.NewNode("C")
	node4 = Map.NewNode("D")
	node5 = Map.NewNode("E")
	node6 = Map.NewNode("F")
	node7 = Map.NewNode("G")
	node8 = Map.NewNode("H")

	gmap.AddNode(node1)
	gmap.AddNode(node2)
	gmap.AddNode(node3)
	gmap.AddNode(node4)
	gmap.AddNode(node5)
	gmap.AddNode(node6)
	gmap.AddNode(node7)
	gmap.AddNode(node8)

	//gmap.Dump()

	gmap.SetUnGraphValue(0, 1, 1)
	gmap.SetUnGraphValue(0, 3, 2)
	gmap.SetUnGraphValue(1, 3, 3)
	gmap.SetUnGraphValue(1, 5, 4)
	gmap.SetUnGraphValue(2, 4, 7)
	gmap.SetUnGraphValue(5, 4, 8)
	gmap.SetUnGraphValue(3, 6, 3)
	gmap.SetUnGraphValue(3, 7, 5)
	gmap.SetUnGraphValue(6, 7, 3)

	gmap.Dump()
	fmt.Println("深度优先")
	gmap.DepthFirstTraverse(0)
	fmt.Println()
	gmap.ResetNode()
	fmt.Println("广度优先")
	gmap.BreadthFirstTraverse(0)
	fmt.Println()
	gmap.ResetNode()
	fmt.Println("普里姆算法最小生成数")
	gmap.PrimTree(0)
}
