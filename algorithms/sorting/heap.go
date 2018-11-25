package sorting

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

func (m *maxHeap) shiftUp(index int) {
	if index == 1 {
		return
	}

	parentIndex := index / 2
	if m.data[index] > m.data[parentIndex] { // 比父结点大，就交换位置
		m.data[index], m.data[parentIndex] = m.data[parentIndex], m.data[index]
		m.shiftUp(parentIndex)
	}

	return
}

func (m *maxHeap) shiftDown(index int) {
	if index == m.count {
		return
	}

	left := 2 * index   // 左结点索引
	if left > m.count { // 左右结点都越界了
		return
	}
	if left == m.count && m.data[index] < m.data[left] { //只有左结点，则当左结点大于父结点时才交换
		m.data[index], m.data[left] = m.data[left], m.data[index]
		return
	}
	if left < m.count { //左右结点都有
		if m.data[left] < m.data[left+1] && m.data[index] < m.data[left+1] {
			m.data[index], m.data[left+1] = m.data[left+1], m.data[index]
			m.shiftDown(left + 1)
			return
		} else if m.data[left] >= m.data[left+1] && m.data[index] > m.data[left] {
			m.data[index], m.data[left] = m.data[left], m.data[index]
			m.shiftDown(left)
			return
		}
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
