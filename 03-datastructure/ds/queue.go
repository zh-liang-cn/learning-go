/*
 * @Author: Liang zh.liang.cn@gmail.com
 * @Date: 2022-08-27 15:14:46
 * @LastEditors: Liang zh.liang.cn@gmail.com
 * @LastEditTime: 2022-08-27 16:41:39
 * @Description: 实现FIFO队列，参考：https://www.tutorialspoint.com/data_structures_algorithms/dsa_queue.htm
 * 		实现 leetcode 上的队列、循环队列、双端循环队列的题目
 */
package ds

type Queue interface {
	Enque(e any) bool
	Deque() bool
	Peek() any
	IsEmpty() bool
	IsFull() bool
}

// TODO 定义 deque 接口

type Node struct {
	prev *Node
	next *Node
	data any
}

type LinkedQueue struct {
	front *Node
	rear  *Node
	size  int
	count int
}

func NewLinkedQueue(n int) *LinkedQueue {
	return &LinkedQueue{size: n}
}

func (q *LinkedQueue) Enque(e any) bool {
	if e == nil || q.count >= q.size {
		return false
	}

	if q.front == nil {
		q.front = &Node{data: e}
		q.rear = q.front
	} else {
		p := q.rear
		q.rear = &Node{data: e, next: p}
		p.prev = q.rear

	}
	q.count++

	return true
}

func (q *LinkedQueue) Deque() bool {
	if q.front == nil {
		return false
	}

	if q.front == q.rear {
		q.front = nil
		q.rear = nil
	} else {
		f := q.front
		q.front = f.prev
		q.front.next = nil

		f.prev = nil
	}
	q.count--

	return true
}

func (q *LinkedQueue) Peek() any {
	if q.count <= 0 {
		return nil
	} else {
		return q.front.data
	}
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.count == 0
}

func (q *LinkedQueue) IsFull() bool {
	return q.count == q.size
}

type ArrayQueue struct {
	arr   []any
	front int // font index
	rear  int // rear index
	size  int
	count int
}

func NewArrayQueue(n int) *ArrayQueue {
	q := &ArrayQueue{size: n, front: -1, rear: -1}
	q.arr = make([]any, q.size)

	return q
}

func (q *ArrayQueue) Enque(e any) bool {
	if e == nil || q.count >= q.size {
		return false
	}

	if q.front+1 < q.size {
		q.front++

		if q.rear == -1 {
			q.rear = q.front
		}
	} else {
		q.front = 0
	}

	q.arr[q.front] = e
	q.count++

	return true
}

func (q *ArrayQueue) Deque() bool {
	if q.count <= 0 {
		return false
	}

	if q.rear == q.front {
		q.rear = -1
		q.front = -1
	}

	if q.rear+1 < q.size {
		q.rear++
	} else {
		q.rear = 0
	}

	q.count--

	return true
}

func (q *ArrayQueue) Peek() any {
	if q.count <= 0 {
		return nil
	}

	return q.arr[q.rear]
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.count == 0
}

func (q *ArrayQueue) IsFull() bool {
	return q.count == q.size
}
