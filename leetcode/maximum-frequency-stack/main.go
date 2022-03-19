package main

import (
	"container/heap"
	"fmt"
)

type freqVal struct {
	Val  int
	Freq int
}

type maxHeap []freqVal

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].Freq > h[j].Freq }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxHeap) Peak() interface{} {
	return h[0]
}

func (h *maxHeap) Get(val int) interface{} {
	for i, v := range *h {
		if v.Val == val {
			old := *h
			*h = append(old[:i], old[i+1:]...)
			return v
		}
	}
	return freqVal{Val: -1, Freq: 0}
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(freqVal))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type FreqStack struct {
	h       maxHeap
	freq    map[int]int
	ordered []int
}

func Constructor() FreqStack {
	h := &maxHeap{}
	heap.Init(h)
	return FreqStack{*h, make(map[int]int), []int{}}
}

func (this *FreqStack) Push(val int) {

	find := this.h.Get(val).(freqVal)
	if find.Val != -1 {
		find.Freq++
		heap.Push(&this.h, find)
	} else {
		heap.Push(&this.h, freqVal{val, 1})
	}
	this.freq[val]++
	if this.freq[val] == 1 {
		this.ordered = append(this.ordered, val)
	}
}

func (this *FreqStack) Pop() int {

	var initVal int
	for _, v := range this.freq {
		initVal = v
		break
	}
	allSameVal := true
	for _, v := range this.freq {
		if v != initVal {
			allSameVal = false
		}
	}
	if allSameVal {
		var top int
		if len(this.ordered) > 0 {
			top = this.ordered[len(this.ordered)-1]
			this.ordered = this.ordered[:len(this.ordered)-1]
			this.h.Get(top)
			delete(this.freq, top)
		}
		return top
	} else {
		top := this.h.Peak().(freqVal)
		heap.Pop(&this.h)
		top.Freq--
		if top.Freq > 0 {
			heap.Push(&this.h, top)
		}
		if val, ok := this.freq[top.Val]; ok {
			if val == 1 {
				delete(this.freq, top.Val)
			} else {
				this.freq[top.Val]--
			}
		}
		return top.Val
	}
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(0)
	obj.Push(0)
	obj.Push(1)
	obj.Push(5)
	obj.Push(4)
	obj.Push(1)
	obj.Push(5)
	obj.Push(1)
	obj.Push(6)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
}
