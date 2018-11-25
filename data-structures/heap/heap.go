package heap

import (
	"errors"
	"fmt"
)

// 堆排序(二叉堆)
// 最大堆：堆中任何一个结点都不大于其父结点，堆总是一个完全二叉树。
// 从数组第一个位置开始存储结点

type maxHeap struct {
	data     []int
	capacity int
	count    int
}

func (m *maxHeap) Size() int {
	return m.count
}

func (m *maxHeap) Empty() bool {
	return m.count == 0
}

func (m *maxHeap) Insert(item int) error {
	if m.count >= m.capacity {
		return errors.New("heap is full")
	}
	m.data[m.count+1] = item
	m.shiftUp(m.count + 1)
	m.count++
	return nil
}

func (m *maxHeap) Delete() (item int, err error) {
	if m.count == 0 {
		return 0, errors.New("heap is empty")
	}

	item = m.data[1]
	m.data[1] = m.data[m.count]
	m.count--
	m.shiftDown(1)

	return item, nil
}

// 采用数组来初始化堆
func (m *maxHeap) Init(arr []int) {
	// 从第一个位置开始填充
	copy(m.data[1:], arr)
	m.count = len(arr)

	for i := m.count / 2; i > 0; i-- {
		m.shiftDown(i)
	}
}

func (m *maxHeap) shiftUp(index int) {
	for index > 1 && m.data[index] > m.data[index/2] {
		m.data[index], m.data[index/2] = m.data[index/2], m.data[index]
		index = index / 2
	}
	return
}

func (m *maxHeap) shiftDown(index int) {
	for 2*index <= m.count {
		j := 2 * index
		if j+1 <= m.count && m.data[j+1] > m.data[j] {
			j += 1
		}

		if m.data[index] >= m.data[j] {
			break
		}

		m.data[index], m.data[j] = m.data[j], m.data[index]
		index = j
	}
}

// TODO(tfbrother) 后期考虑把二叉堆的格式化输出代码实现
func (m *maxHeap) Dump() {
	fmt.Println(m.data[1 : m.count+1])
}

func NewMaxHeap(capacity int) *maxHeap {
	return &maxHeap{
		data:     make([]int, capacity+1),
		capacity: capacity,
		count:    0,
	}
}
