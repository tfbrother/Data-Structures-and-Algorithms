package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/graph"
)

var (
	node1, node2, node3, node4, node5, node6 graph.Node
	gmap                                     *graph.Gmap
)

func main() {
	gmap = graph.NewGmap(6)
	node1 = graph.NewNode("A")
	node2 = graph.NewNode("B")
	node3 = graph.NewNode("C")
	node4 = graph.NewNode("D")
	node5 = graph.NewNode("E")
	node6 = graph.NewNode("F")

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
	//gmap.ResetNode()
	//fmt.Println("普里姆算法最小生成数")
	//gmap.PrimTree(0)

	fmt.Println("克鲁斯卡最小生成数算法")
	gmap.ResetNode()
	gmap.KruskalTree()

	fmt.Println("\n\n\n=============稠密无向图测试=============")
	d := graph.NewDenseGraph(10, false)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			d.AddEdge(i, j)
		}
	}
	fmt.Println("边数量：", d.EdgeCount())
	fmt.Println("结点数量：", d.NodeCount())

	fmt.Println("\n\n\n=============广度遍历优先=============")
	d.BreadthFirstTraverse(1)
	fmt.Println()
	fmt.Println()

	fmt.Println("\n\n\n=============深度遍历优先=============")
	d.ResetVisted()
	d.DepthFirstTraverse(1)
	fmt.Println()
	fmt.Println()
}
