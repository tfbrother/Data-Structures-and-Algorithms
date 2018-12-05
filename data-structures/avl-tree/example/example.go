package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/avl-tree"
	"strconv"
)

func main() {
	avl := avl_tree.NewAVL()
	for i := 0; i < 100; i++ {
		avl.Add("tfbrother"+strconv.Itoa(i), strconv.Itoa(i))
	}

	fmt.Println("树的大小", avl.Size())
	fmt.Println("所有的keys", avl.InOrder())
	fmt.Println("是否是二叉搜索树", avl.IsBST())
	fmt.Println("是否是平衡二叉搜索树", avl.IsBalanced())
}
