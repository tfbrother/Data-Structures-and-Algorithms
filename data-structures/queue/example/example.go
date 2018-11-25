package main

import "github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/queue"

func main() {
	var capacity, i queue.Node
	capacity = 10
	ring := queue.NewRingQueue(int(capacity))
	for i = 0; i < capacity; i++ {
		ring.PushNode(i)
	}

	ring.Dump() //0,1,2,3,4,5,6,7,8,9

	ring.PopNode()
	ring.PopNode()
	ring.PopNode()
	ring.PopNode()
	ring.Dump() //4,5,6,7,8,9
}
