package priority_queue

/**
优先级队列，采用索引最大堆实现
*/

type Item interface {
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type node struct {
	item  Item
	index int // 索引
}

type PriorityQueue struct {
	items    []*node
	indexs   []int
	count    int
	capacity int
}

// heapify
func (p *PriorityQueue) Init(items []Item) {

}

func (p *PriorityQueue) Push(item Item) bool {
	return true
}

func (p *PriorityQueue) Pop() Item {
	return nil
}

func (p *PriorityQueue) Remove(index int) {

}

func (p *PriorityQueue) RemoveEq(item Item) {

}

func (p *PriorityQueue) siftDown(index int) bool {
	return true
}

func (p *PriorityQueue) siftUp(index int) {

}

func New(capacity int) *PriorityQueue {
	return &PriorityQueue{
		items:    make([]*node, capacity+1),
		capacity: capacity,
		count:    0,
	}
}
