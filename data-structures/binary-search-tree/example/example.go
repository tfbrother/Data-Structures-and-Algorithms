package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/binary-search-tree"
	"strconv"
)

type Node struct {
	key   int
	value string
}

func (n *Node) Less(b binary_search_tree.Item) bool {
	return n.key < b.(*Node).key
}

func (n *Node) ToString() string {
	return strconv.Itoa(n.key)
}

func main() {

	BST := binary_search_tree.NewBst(&Node{100, "tfbrother100"})
	//fmt.Println(BST.GetSize())

	BST.Add(&Node{200, "tfbrother200"})
	BST.Add(&Node{90, "tfbrother90"})
	BST.Add(&Node{99, "tfbrother99"})
	BST.Add(&Node{98, "tfbrother98"})
	BST.Add(&Node{89, "tfbrother89"})
	BST.Add(&Node{101, "tfbrother101"})
	BST.Add(&Node{299, "tfbrother299"})

	fmt.Println("二叉搜索树前序遍历(循环实现)结果为：")
	BST.PrevOrder()
	fmt.Println()
	fmt.Println("二叉搜索树前序遍历(递归实现)结果为：")
	BST.PrevOrderRecursion()
	fmt.Println()
	fmt.Println("二叉搜索树中序遍历(循环实现)结果为：")
	BST.InOrder()
	fmt.Println()
	fmt.Println("二叉搜索树中序遍历(递归实现)结果为：")
	BST.InOrderRecursion()
	fmt.Println()
	fmt.Println("二叉搜索树层序遍历结果为：")
	BST.LevelOrder()
	fmt.Println()

	//minNum := BST.MinNum()
	//fmt.Println("二叉搜索树的最小值为：", minNum)
	//
	//maxNum := BST.MaxNum()
	//fmt.Println("二叉搜索树的最大值为：", maxNum)
	//
	//BST.RemoveMin()
	//fmt.Println("删除最小值的，前序遍历")
	//BST.PrevOrder()
	//
	//BST.RemoveMax()
	//fmt.Println("删除最大值的，前序遍历")
	//BST.PrevOrder()
	//
	//node := &Node{201, "tfbrother201"}
	//found := BST.Find(node)
	//if found != nil {
	//	fmt.Println("查找数据，", found.ToString())
	//} else {
	//	fmt.Println("没有查找到，", node.ToString())
	//}
	//
	//BST.Remove(&Node{98, "tfbrother200"})
	//fmt.Println("删除任意值，比如98后的，前序遍历")
	//BST.PrevOrder()
}
