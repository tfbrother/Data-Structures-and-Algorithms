package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/ConsistentHash"
	"strconv"
)

func main() {
	var chash *ConsistentHash.Map = ConsistentHash.New(10)
	chash.Add([]int{1, 2, 3, 4, 5, 6})
	for i := 0; i < 10; i++ {
		fmt.Println(chash.Get(strconv.Itoa(i) + "ijibu"))
	}

	//chash.Dump()
}
