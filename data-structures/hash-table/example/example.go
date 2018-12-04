package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/hash-table"
)

func main() {
	hashTable := hash_table.NewHashTable()

	hashTable.Add("name", "tfbrother")

	fmt.Println(hashTable.Contains("name"))
	val := hashTable.Get("name")
	fmt.Println(val)
}
