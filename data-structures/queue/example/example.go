package main

import "github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/queue"

func main() {
	capacity := 10
	ring := queue.NewRingQueue(capacity)
	for i := 0; i < capacity; i++ {
		ring.Push(i)
	}

	ring.Dump() //0,1,2,3,4,5,6,7,8,9

	ring.Pop()
	ring.Pop()
	ring.Pop()
	ring.Pop()
	ring.Dump() //4,5,6,7,8,9
}
