package heap

import (
	"errors"
	"fmt"
)

// 堆排序(二叉堆)
// 最大堆：堆中任何一个结点都不大于其父结点，堆总是一个完全二叉树。
// 从数组第一个位置开始存储结点

type Item interface {
	Less(a Item) bool
	Swap(a Item)
	ToString() string
}

type maxHeap struct {
	data     []Item
	capacity int
	count    int
}

func (m *maxHeap) Size() int {
	return m.count
}

func (m *maxHeap) Empty() bool {
	return m.count == 0
}

func (m *maxHeap) Push(item Item) error {
	if m.count >= m.capacity {
		return errors.New("heap is full")
	}
	m.data[m.count+1] = item
	m.siftUp(m.count + 1)
	m.count++
	return nil
}

// 删除的是堆的根结点，对于最小堆来说，就是最小值，对于最大堆来说就是最小值，这也就是可以利用堆排序的原理。
func (m *maxHeap) Pop() (item Item, err error) {
	if m.count == 0 {
		return nil, errors.New("heap is empty")
	}

	item = m.data[1]
	m.data[1] = m.data[m.count]
	m.count--
	m.siftDown(1)

	return item, nil
}

// 采用数组来初始化堆
func (m *maxHeap) Init(arr []Item) {
	// 从第一个位置开始填充
	copy(m.data[1:], arr)
	m.count = len(arr)

	for i := m.count / 2; i > 0; i-- {
		m.siftDown(i)
	}
}

func (m *maxHeap) siftUp(index int) {
	for index > 1 && m.data[index/2].Less(m.data[index]) {
		m.data[index].Swap(m.data[index/2])
		index = index / 2
	}
	return
}

func (m *maxHeap) siftDown(index int) {
	for 2*index <= m.count {
		j := 2 * index
		if j+1 <= m.count && m.data[j].Less(m.data[j+1]) {
			j += 1
		}

		if m.data[j].Less(m.data[index]) {
			break
		}

		m.data[index].Swap(m.data[j])
		index = j
	}
}

// TODO(tfbrother) 后期考虑把二叉堆的格式化输出代码实现
func (m *maxHeap) Dump() {
	for i := 0; i < m.count; i++ {
		fmt.Println(m.data[i+1].ToString())
	}

}

func NewMaxHeap(capacity int) *maxHeap {
	return &maxHeap{
		data:     make([]Item, capacity+1),
		capacity: capacity,
		count:    0,
	}
}
