package main

import (
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/stack"
)

func main() {
	var capacity int
	capacity = 10
	s := stack.NewStack(capacity)
	for i := 0; i < capacity; i++ {
		s.Push(i)
	}

	s.Dump() //9,8,7,6,5,4,3,2,1,0

	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Dump() //5,4,3,2,1,0
}
