/*
https://leetcode.com/problems/min-stack/
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
Implement the MinStack class:
	MinStack() initializes the stack object.
	void push(int val) pushes the element val onto the stack.
	void pop() removes the element on the top of the stack.
	int top() gets the top element of the stack.
	int getMin() retrieves the minimum element in the stack.

Example 1:
	Input
	["MinStack","push","push","push","getMin","pop","top","getMin"]
	[[],[-2],[0],[-3],[],[],[],[]]

	Output
	[null,null,null,null,-3,null,0,-2]

	Explanation
	MinStack minStack = new MinStack();
	minStack.push(-2);
	minStack.push(0);
	minStack.push(-3);
	minStack.getMin(); // return -3
	minStack.pop();
	minStack.top();    // return 0
	minStack.getMin(); // return -2
*/

package main

import "fmt"

type MinStack struct {
	vector  []int
	min     int
	minInit bool
}

func Constructor() MinStack {
	return MinStack{[]int{}, 0, false}
}

func (this *MinStack) Push(val int) {
	this.vector = append(this.vector, val)
	this.minInit = false
}

func (this *MinStack) Pop() {
	this.vector = this.vector[:len(this.vector)-1]
	this.minInit = false
}

func (this *MinStack) Top() int {
	if len(this.vector) > 0 {
		return this.vector[len(this.vector)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if this.minInit {
		return this.min
	}

	if len(this.vector) == 0 {
		return -1
	}

	min := this.vector[0]
	for _, v := range this.vector {
		if v < min {
			min = v
		}
	}
	this.min = min
	this.minInit = true
	return this.min
}

func main() {
	minStack := Constructor()
	minStack.Push(2)
	minStack.Push(0)
	minStack.Push(3)
	minStack.Push(0)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	minStack.Pop()
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}
