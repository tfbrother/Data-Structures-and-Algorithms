package main

import (
	"fmt"
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/trie"
)

func main() {
	words := []string{"w", "wo", "wor", "worl", "world"}
	t := trie.NewTrie()

	for i := 0; i < len(words); i++ {
		t.Add(words[i])
	}

	flag := t.Find("wor")
	fmt.Println(flag)

}
