package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/union-find"
)

// 并查集
func main() {
	n := 10000
	u := union_find.New(n)

	fmt.Println(u.Find(10))
	u.Union(10, 100)
	fmt.Println(u.Find(10))
	fmt.Println(u.Find(100))
	fmt.Println(u.IsConnected(10, 100))
	fmt.Println(u.IsConnected(10, 101))
}
