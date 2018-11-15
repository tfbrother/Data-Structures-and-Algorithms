package main

import "github.com/tfbrother/Data-Structures-and-Algorithms/BinaryTree"

func main() {
	/*
					5(0)					左子节点索引：2*X+1
			7(1)			8(2)			右子节点索引：2*X+2
		8(3)	10(4)	3(5)	2(6)
	*/
	tree := BinaryTree.NewTree(10, 5)
	tree.AddNode(0, 0, 7)
	tree.AddNode(0, 1, 8)
	tree.AddNode(1, 0, 8)
	tree.AddNode(1, 1, 10)
	tree.AddNode(2, 0, 3)
	tree.AddNode(2, 1, 2)

	tree.Dump()
	tree.DeleteNode(1)
	tree.Dump()
}
