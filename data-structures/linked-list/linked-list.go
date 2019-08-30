package linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建链表(顺序)
func CreateLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	curNode := head

	for i := 1; i < len(arr); i++ {
		curNode.Next = &ListNode{Val: arr[i]}
		curNode = curNode.Next
	}

	return head
}

// 创建链表(倒序)
func CreateLinkedList1(arr []int) *ListNode {
	n := len(arr)
	if n == 0 {
		return nil
	}

	next := &ListNode{Val: arr[n-1]}
	curNode := next
	for i := n - 2; i >= 0; i-- {
		curNode = &ListNode{Val: arr[i]}
		curNode.Next = next
		next = curNode
	}

	return curNode
}

// 打印链表
func PrintLinkedList(head *ListNode) {
	curNode := head
	for curNode != nil {
		fmt.Printf("%d -> ", curNode.Val)
		curNode = curNode.Next
	}

	fmt.Print("nil\n")
}

// 反转链表
func ReverseList(head *ListNode) *ListNode {
	var prev, curNode, next *ListNode

	curNode = head
	for curNode != nil {
		next = curNode.Next
		curNode.Next, prev, curNode = prev, curNode, next
	}

	return prev
}

// 获取链表的中间结点(如果有两个中间结点，则返回第二个中间结点。)
// 快慢指针思想，注意slow、fast的初始值都为head
// https://leetcode-cn.com/problems/middle-of-the-linked-list/
func GetMiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow
}
