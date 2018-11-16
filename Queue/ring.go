// 环形队列(采用顺序表数组存储)
package Queue

import "log"

type Node int

type RingQueue struct {
	head     int    // 头部索引
	tail     int    // 尾部索引
	nodes    []Node // 所有的数据
	len      int    // 队列的长度
	capacity int    // 队列的容量
}

// 入队列
func (r *RingQueue) PushNode(node Node) bool {
	if r.Full() {
		return false
	}

	r.nodes[r.tail] = node
	r.tail++
	r.tail = r.tail % r.capacity
	r.len++
	return true
}

// 出队列
func (r *RingQueue) PopNode() (node Node, flag bool) {
	if r.Empty() {
		return
	}
	flag = true
	node = r.nodes[r.head]
	r.nodes[r.head] = 0
	r.head++
	r.head = r.head % r.capacity
	r.len--
	return
}

// 返回队列的长度
func (r *RingQueue) Len() int {
	return r.len
}

// 是否为空
func (r *RingQueue) Empty() bool {
	if r.len == 0 {
		return true
	}
	return false
}

// 是否已满
func (r *RingQueue) Full() bool {
	if r.len == r.capacity {
		return true
	}
	return false
}

// 遍历队列
func (r *RingQueue) Dump() {
	for i := r.head; i < r.head+r.len; i++ {
		log.Println(r.nodes[i%r.capacity])
	}
}

func NewRingQueue(capacity int) *RingQueue {
	return &RingQueue{
		head:     0,
		tail:     0,
		nodes:    make([]Node, capacity),
		len:      0,
		capacity: capacity,
	}
}
