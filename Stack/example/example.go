package main

import (
	"github.com/tfbrother/Data-Structures-and-Algorithms/Stack"
)

func main() {
	var capacity, i Stack.Node
	capacity = 10
	s := Stack.NewStack(int(capacity))
	for i = 0; i < capacity; i++ {
		s.PushNode(i)
	}

	s.Dump() //9,8,7,6,5,4,3,2,1,0

	s.PopNode()
	s.PopNode()
	s.PopNode()
	s.PopNode()
	s.Dump() //5,4,3,2,1,0
}
