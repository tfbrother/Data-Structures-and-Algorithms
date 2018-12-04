package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/binary-search-tree"
)

func main() {
	BST := binary_search_tree.NewBst(&binary_search_tree.Node{100, nil, nil})
	//fmt.Println(BST.GetSize())

	BST.Add(&binary_search_tree.Node{200, nil, nil})
	BST.Add(&binary_search_tree.Node{90, nil, nil})
	BST.Add(&binary_search_tree.Node{99, nil, nil})
	BST.Add(&binary_search_tree.Node{98, nil, nil})
	BST.Add(&binary_search_tree.Node{89, nil, nil})
	BST.Add(&binary_search_tree.Node{101, nil, nil})
	BST.Add(&binary_search_tree.Node{299, nil, nil})

	BST.PrevOrder()

	minNum := BST.MinNum()
	fmt.Println("二叉搜索树的最小值为：", minNum)

	maxNum := BST.MaxNum()
	fmt.Println("二叉搜索树的最大值为：", maxNum)

	BST.RemoveMin()
	fmt.Println("删除最小值的，前序遍历")
	BST.PrevOrder()

	BST.RemoveMax()
	fmt.Println("删除最大值的，前序遍历")
	BST.PrevOrder()

	node := BST.Find(200)
	fmt.Println("查找200，", node)

	BST.Remove(98)
	fmt.Println("删除任意值，比如98后的，前序遍历")
	BST.PrevOrder()
}
