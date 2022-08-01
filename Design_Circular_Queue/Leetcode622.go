package main

type MyCircularQueue struct {
	front, back int
	elements    []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{elements: make([]int, k+1)}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.elements[q.back] = value
	q.back = (q.back + 1) % len(q.elements)
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % len(q.elements)
	return true
}

func (q MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[q.front]
}

func (q MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[(q.back-1+len(q.elements))%len(q.elements)]
}

func (q MyCircularQueue) IsEmpty() bool {
	return q.back == q.front
}

func (q MyCircularQueue) IsFull() bool {
	return (q.back+1)%len(q.elements) == q.front
}
