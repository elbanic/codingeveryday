/*
https://leetcode.com/problems/moving-average-from-data-stream/
Given a stream of integers and a window size, calculate the moving average of all integers in the sliding window.

Implement the MovingAverage class:
	MovingAverage(int size) Initializes the object with the size of the window size.
	double next(int val) Returns the moving average of the last size values of the stream.

Example 1:
	Input
	["MovingAverage", "next", "next", "next", "next"]
	[[3], [1], [10], [3], [5]]
	Output
	[null, 1.0, 5.5, 4.66667, 6.0]

	Explanation
	MovingAverage movingAverage = new MovingAverage(3);
	movingAverage.next(1); // return 1.0 = 1 / 1
	movingAverage.next(10); // return 5.5 = (1 + 10) / 2
	movingAverage.next(3); // return 4.66667 = (1 + 10 + 3) / 3
	movingAverage.next(5); // return 6.0 = (10 + 3 + 5) / 3
*/

package main

import "fmt"

//addition and subtraction
type MyCircularQueue struct {
	vector []int
	head   int
	tail   int
	size   int
}

func NewMyCircularQueue(k int) MyCircularQueue {
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
	this.size++
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
	this.size--
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
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == len(this.vector)
}

type MovingAverage struct {
	cQueue MyCircularQueue
	sum    int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{NewMyCircularQueue(size), 0}
}

func (this *MovingAverage) Next(val int) float64 {

	if this.cQueue.size < len(this.cQueue.vector) {
		this.cQueue.EnQueue(val)
		this.sum += val
	} else {
		this.sum -= this.cQueue.Front()
		this.cQueue.DeQueue()
		this.cQueue.EnQueue(val)
		this.sum += val
	}
	return float64(this.sum) / float64(this.cQueue.size)
}

//queue
//type MovingAverage struct {
//	queue []int
//	loc   int
//	count int
//}
//
//func Constructor(size int) MovingAverage {
//	return MovingAverage{make([]int, size), -1, 0}
//}
//
//func (this *MovingAverage) Next(val int) float64 {
//	this.add(val)
//	var sum int
//	for _, v := range this.queue {
//		sum += v
//	}
//	return float64(sum) / float64(this.count)
//}
//
//func (this *MovingAverage) add(val int) {
//	this.loc++
//	if this.loc == len(this.queue) {
//		this.loc = 0
//	}
//	if this.count >= len(this.queue) {
//		this.count = len(this.queue)
//	} else {
//		this.count++
//	}
//	this.queue[this.loc] = val
//}

func main() {
	obj := Constructor(1)
	fmt.Println(obj.Next(4))
	fmt.Println(obj.Next(0))
}
