package leetcode

import "testing"

type MyCircularDeque struct {
	q        []int
	capacity int

	head int
	// 指向next 存储的位置
	tail int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		q:        make([]int, k+1),
		capacity: k + 1,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head = prevPos(this.head, this.capacity)
	this.q[this.head] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.q[this.tail] = value
	this.tail = nextPos(this.tail, this.capacity)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = nextPos(this.head, this.capacity)
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.tail = (this.tail + this.capacity - 1) % this.capacity
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.q[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.q[prevPos(this.tail, this.capacity)]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.head == this.tail {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if nextPos(this.tail, this.capacity) == this.head {
		return true
	}
	return false
}

func prevPos(p, size int) int {
	return (p + size - 1) % size
}

func nextPos(p, size int) int {
	return (p + 1) % size
}

func TestCircularDeque(t *testing.T) {
	cd := Constructor(3)
	t.Log(cd.InsertLast(1))
	cd.InsertLast(2)
	cd.InsertFront(3)
	cd.InsertFront(4)

	t.Log(cd.GetRear())
}
