package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/heap"
	"github.com/tfbrother/Data-Structures-and-Algorithms/util"
)

var (
	myarr, myarr1 []int
)

func main() {
	// 堆
	myarr = util.GenrateRandomArray(30, 10, 100)
	fmt.Println(myarr)
	myHeap := heap.NewMaxHeap(30)
	for _, v := range myarr {
		myHeap.Insert(v)
	}
	myHeap.Dump()
	myHeap.Delete()
	myHeap.Dump()

	myarr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(myarr1)
	myHeap.Init(myarr1)
	myHeap.Dump()
}
