package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/binary-search-tree"
)

func main() {
	bst := binary_search_tree.NewBst(&binary_search_tree.Node{100, nil, nil})
	//fmt.Println(bst.GetSize())

	bst.Add(&binary_search_tree.Node{200, nil, nil})
	bst.Add(&binary_search_tree.Node{90, nil, nil})
	bst.Add(&binary_search_tree.Node{99, nil, nil})
	bst.Add(&binary_search_tree.Node{98, nil, nil})
	bst.Add(&binary_search_tree.Node{89, nil, nil})
	bst.Add(&binary_search_tree.Node{101, nil, nil})
	bst.Add(&binary_search_tree.Node{299, nil, nil})
	bst.PrevOrder()

	minNum := bst.MinNum()
	fmt.Println("二叉搜索树的最小值为：", minNum)

	maxNum := bst.MaxNum()
	fmt.Println("二叉搜索树的最大值为：", maxNum)

	bst.RemoveMin()
	fmt.Println("删除最小值的，前序遍历")
	bst.PrevOrder()

	bst.RemoveMax()
	fmt.Println("删除最大值的，前序遍历")
	bst.PrevOrder()
}