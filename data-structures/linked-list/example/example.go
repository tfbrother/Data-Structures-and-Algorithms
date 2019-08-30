package main

import (
	"github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/linked-list"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	head1 := linked_list.CreateLinkedList(arr)
	head2 := linked_list.CreateLinkedList1(arr)

	linked_list.PrintLinkedList(head1)
	linked_list.PrintLinkedList(head2)
}
