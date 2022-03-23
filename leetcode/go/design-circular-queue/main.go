/*
https://leetcode.com/problems/design-circular-queue/
Design your implementation of the circular queue. The circular queue is a linear data structure
in which the operations are performed based on FIFO (First In First Out) principle and
the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue.
In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue.
ut using the circular queue, we can use the space to store new values.

Implementation the MyCircularQueue class:
	MyCircularQueue(k) Initializes the object with the size of the queue to be k.
	int Front() Gets the front item from the queue. If the queue is empty, return -1.
	int Rear() Gets the last item from the queue. If the queue is empty, return -1.
	boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
	boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
	boolean isEmpty() Checks whether the circular queue is empty or not.
	boolean isFull() Checks whether the circular queue is full or not.
You must solve the problem without using the built-in queue data structure in your programming language.
*/

package main

import "fmt"

type MyCircularQueue struct {
	vector []int
	head   int
	tail   int
	cap    int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{make([]int, k), -1, -1, 0}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.head == -1 {
		this.head = 0
	}
	if this.tail == len(this.vector)-1 {
		this.tail = 0
		this.vector[this.tail] = value
	} else {
		this.tail++
		this.vector[this.tail] = value
	}
	this.cap++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == len(this.vector)-1 {
		this.head = 0
	} else {
		this.head++
	}
	this.cap--
	return true
}

func (this *MyCircularQueue) Front() int {
	if !this.IsEmpty() && this.head >= 0 {
		return this.vector[this.head]
	}
	return -1
}

func (this *MyCircularQueue) Rear() int {
	if !this.IsEmpty() && this.tail >= 0 {
		return this.vector[this.tail]
	}
	return -1
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.cap == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.cap == len(this.vector)
}

func main() {
	obj := Constructor(3)
	obj.EnQueue(1)
	obj.EnQueue(2)
	obj.EnQueue(3)
	obj.EnQueue(4)
	fmt.Println(obj.Rear())
	fmt.Println(obj.IsFull())
	obj.DeQueue()
	obj.EnQueue(4)
	fmt.Println(obj.Rear())
}
