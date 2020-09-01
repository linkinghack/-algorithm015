/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */
package main

// @lc code=start
type MyCircularDeque struct {
	Data     []int
	Front    int
	Rear     int
	Capacity int
	Size     int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	deque := MyCircularDeque{
		Data:     make([]int, k),
		Capacity: k,
		Size:     0,
		Front:    -1,
		Rear:     -1,
	}
	return deque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.Size >= this.Capacity {
		return false
	}

	if this.Front == -1 { // initial state
		this.Front = 0
		this.Rear = 0
	} else if this.Front == 0 { // Front edge
		this.Front = this.Capacity - 1
	} else {
		this.Front--
	}

	this.Data[this.Front] = value
	this.Size++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.Size >= this.Capacity {
		return false
	}

	if this.Front == -1 {
		this.Front = 0
		this.Rear = 0
	} else if this.Rear == this.Capacity-1 { // Rear edge
		this.Rear = 0
	} else { // Normal
		this.Rear++
	}

	this.Data[this.Rear] = value
	this.Size++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.Size <= 0 {
		return false
	}

	if this.Size == 1 {
		this.Front = -1
		this.Rear = -1
	} else if this.Front == this.Capacity-1 {
		this.Front = 0
	} else {
		this.Front++
	}

	this.Size--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.Size <= 0 {
		return false
	}

	if this.Size == 1 {
		this.Front = -1
		this.Rear = -1
	} else if this.Rear == 0 {
		this.Rear = this.Capacity - 1
	} else {
		this.Rear--
	}

	this.Size--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.Size <= 0 {
		return -1
	}
	return this.Data[this.Front]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.Size <= 0 {
		return -1
	}
	return this.Data[this.Rear]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.Size <= 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.Size >= this.Capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end
