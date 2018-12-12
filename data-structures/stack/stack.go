package stack

import "log"

type Stack struct {
	head     int           // 栈顶索引
	tail     int           // 栈底索引
	nodes    []interface{} // 所有的数据
	len      int           // 栈的长度
	capacity int           // 栈的容量
}

// 入栈
func (s *Stack) Push(node interface{}) bool {
	if s.Full() {
		return false
	}

	s.nodes[s.head] = node
	s.head++
	s.len++
	return true
}

// 出栈
func (s *Stack) Pop() (node interface{}) {
	if s.Empty() {
		return
	}

	node = s.nodes[s.head-1]
	s.nodes[s.head-1] = 0
	s.head--
	s.len--
	return
}

// 返回栈的长度
func (s *Stack) Len() int {
	return s.len
}

// 是否为空
func (s *Stack) Empty() bool {
	if s.len == 0 {
		return true
	}
	return false
}

// 是否已满
func (s *Stack) Full() bool {
	if s.len == s.capacity {
		return true
	}
	return false
}

// 遍历栈
func (s *Stack) Dump() {
	for i := s.head - 1; i >= s.tail; i-- {
		log.Println(s.nodes[i])
	}
}

func NewStack(capacity int) *Stack {
	return &Stack{
		head:     0,
		tail:     0,
		nodes:    make([]interface{}, capacity),
		len:      0,
		capacity: capacity,
	}
}
