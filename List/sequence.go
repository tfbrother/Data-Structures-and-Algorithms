// 顺序表
package List

import (
	"errors"
	"fmt"
)

/*
线性表-顺序表
4	5	8	9	3	7
   前倾    后继
*/

type Element int

type SequenceList struct {
	head     int        // 头部索引
	tail     int        // 尾部索引
	ele      []*Element // 所有的数据
	len      int        // 顺序表的长度
	capacity int        // 顺序表的容量
}

// 返回表的长度
func (s *SequenceList) Len() int {
	return s.len
}

// 返回表是否已满
func (s *SequenceList) Full() bool {
	if s.len >= s.capacity {
		return true
	}

	return false
}

// 返回表是否为空
func (s *SequenceList) Empty() bool {
	if s.len == 0 {
		return true
	}

	return false
}

// 返回表头元素
func (s *SequenceList) Front() *Element {
	if s.len == 0 {
		return nil
	}
	return s.ele[0]
}

// 返回表尾元素
func (s *SequenceList) Back() *Element {
	if s.len == 0 {
		return nil
	}
	return s.ele[s.len-1]
}

// 往表的指定位置插入元素，后面的元素依次往后移动一位
func (s *SequenceList) InsertIndex(index int, ele *Element) (err error) {
	if index > s.len || index < 0 {
		return errors.New("index error")
	}

	if index == s.len { //在线性表尾部插入，就不需要移动元素
		s.ele[index] = ele
		s.len++
		return nil
	}

	//头部或者中间插入，要把之前的元素依次往后移动一位
	for i := index; i <= s.len; i++ {
		s.ele[i+1] = s.ele[i]
	}
	s.ele[index] = ele
	s.len++

	return nil
}

// 遍历队列
func (s *SequenceList) Dump() {
	for i := 0; i < s.len; i++ {
		fmt.Print(*s.ele[i])
	}
	fmt.Println()
}

func NewSequenceList(capacity int) *SequenceList {
	return &SequenceList{
		capacity: capacity,
		len:      0,
		head:     0,
		tail:     0,
		ele:      make([]*Element, capacity),
	}
}
