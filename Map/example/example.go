package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/Map"
)

var (
	node1, node2, node3, node4, node5, node6 Map.Node
	gmap                                     *Map.Gmap
)

func main() {
	gmap = Map.NewGmap(6)
	node1 = Map.NewNode("A")
	node2 = Map.NewNode("B")
	node3 = Map.NewNode("C")
	node4 = Map.NewNode("D")
	node5 = Map.NewNode("E")
	node6 = Map.NewNode("F")

	gmap.AddNode(node1)
	gmap.AddNode(node2)
	gmap.AddNode(node3)
	gmap.AddNode(node4)
	gmap.AddNode(node5)
	gmap.AddNode(node6)

	//gmap.Dump()

	gmap.SetUnGraphValue(0, 1, 6)
	gmap.SetUnGraphValue(0, 4, 5)
	gmap.SetUnGraphValue(0, 5, 1)
	gmap.SetUnGraphValue(1, 2, 3)
	gmap.SetUnGraphValue(1, 5, 2)
	gmap.SetUnGraphValue(2, 5, 8)
	gmap.SetUnGraphValue(2, 3, 7)
	gmap.SetUnGraphValue(3, 4, 2)
	gmap.SetUnGraphValue(3, 5, 4)
	gmap.SetUnGraphValue(4, 5, 9)

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
