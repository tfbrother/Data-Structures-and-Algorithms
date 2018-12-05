package main

import (
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/avl-tree"
	"strconv"
)

func main() {
	avl := avl_tree.NewAVL()
	for i := 0; i < 100; i++ {
		avl.Add("tfbrother"+strconv.Itoa(i), strconv.Itoa(i))
	}
}
