package queue

// 队列

type Queue struct {
	data     []interface{} // 保存数据
	count    int           // 记录数量
	capacity int           // 队列的容量
}

// 入队列
func (q *Queue) Push(item interface{}) {
	q.data = append(q.data, item)
	q.count++
}

// 出队列
func (q *Queue) Pop() (ret interface{}) {
	if q.count > 0 {
		ret = q.data[0]
		q.data = q.data[1:]
		q.count--
	}

	return
}

func (q *Queue) Len() int {
	return q.count
}

func (q *Queue) Empty() bool {
	return q.count == 0
}

func (q *Queue) Full() bool {
	return q.count == q.capacity
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		data:     make([]interface{}, 0),
		count:    0,
		capacity: capacity,
	}
}
