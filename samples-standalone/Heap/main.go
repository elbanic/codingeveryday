package main

import (
	"fmt"
	"strconv"
)

type minHeap struct {
	minimumHeap []int
	heapSize    int
	realSize    int
}

type maxHeap struct {
	maximumHeap []int
	heapSize    int
	realSize    int
}

func newMinHeap(heapSize int) minHeap {
	return minHeap{minimumHeap: []int{0}, heapSize: heapSize+1, realSize: 0}
}

func newMaxHeap(heapSize int) maxHeap {
	return maxHeap{maximumHeap: []int{0}, heapSize: heapSize+1, realSize: 0}
}

func (h *minHeap) add(e int) {
	h.realSize++
	if h.realSize > h.heapSize {
		fmt.Println("Add too many elements!")
		h.realSize--
		return
	}

	h.minimumHeap = append(h.minimumHeap, e)
	index := h.realSize
	parent := index / 2

	for h.minimumHeap[index] < h.minimumHeap[parent] && index > 1 {
		temp := h.minimumHeap[index]
		h.minimumHeap[index] = h.minimumHeap[parent]
		h.minimumHeap[parent] = temp
		index = parent
		parent = index / 2
	}
}

func (h *maxHeap) add(e int) {
	h.realSize++
	if h.realSize > h.heapSize {
		fmt.Println("Add too many elements!")
		h.realSize--
		return
	}

	h.maximumHeap = append(h.maximumHeap, e)
	index := h.realSize
	parent := index / 2

	for h.maximumHeap[index] > h.maximumHeap[parent] && index > 1 {
		temp := h.maximumHeap[index]
		h.maximumHeap[index] = h.maximumHeap[parent]
		h.maximumHeap[parent] = temp
		index = parent
		parent = index / 2
	}
}

func (h minHeap) peak() int{
	return h.minimumHeap[1]
}

func (h maxHeap) peak() int{
	return h.maximumHeap[1]
}

func (h *minHeap) pop() int{

	if h.realSize < 1 {
		fmt.Println("Don't have any element!")
		return int(^uint(0) >> 1) // integer MAX
	} else {
		removeElement := h.minimumHeap[1]
		h.minimumHeap[1] = h.minimumHeap[h.realSize]
		h.realSize--
		index := 1

		for index < h.realSize && index <= h.realSize / 2 {
			left := index * 2
			right := (index * 2) + 1

			if h.minimumHeap[index] > h.minimumHeap[left] || h.minimumHeap[index] > h.minimumHeap[right] {
				if h.minimumHeap[left] < h.minimumHeap[right] {
					temp := h.minimumHeap[left]
					h.minimumHeap[left] = h.minimumHeap[index]
					h.minimumHeap[index] = temp
					index = left
				} else {
					temp := h.minimumHeap[right]
					h.minimumHeap[right] = h.minimumHeap[index]
					h.minimumHeap[index] = temp
					index = right
				}
			} else {
				break
			}
		}
		h.minimumHeap = h.minimumHeap[0: h.realSize+1]
		return removeElement
	}
}

func (h *maxHeap) pop() int{

	if h.realSize < 1 {
		fmt.Println("Don't have any element!")
		return int(^uint(0) >> 1) // integer MAX
	} else {
		removeElement := h.maximumHeap[1]
		h.maximumHeap[1] = h.maximumHeap[h.realSize]
		h.realSize--
		index := 1

		for index < h.realSize && index <= h.realSize / 2 {
			left := index * 2
			right := (index * 2) + 1

			if h.maximumHeap[index] < h.maximumHeap[left] || h.maximumHeap[index] < h.maximumHeap[right] {
				if h.maximumHeap[left] > h.maximumHeap[right] {
					temp := h.maximumHeap[left]
					h.maximumHeap[left] = h.maximumHeap[index]
					h.maximumHeap[index] = temp
					index = left
				} else {
					temp := h.maximumHeap[right]
					h.maximumHeap[right] = h.maximumHeap[index]
					h.maximumHeap[index] = temp
					index = right
				}
			} else {
				break
			}
		}
		h.maximumHeap = h.maximumHeap[0: h.realSize+1]
		return removeElement
	}
}


func (h minHeap) size() int{
	return h.realSize
}

func (h minHeap) print() {
	if h.realSize == 0 {
		fmt.Println("No element")
	} else {
		fmt.Print("[")
		for i := 1; i<= h.realSize-1; i++ {
			fmt.Print(h.minimumHeap[i], ",")
		}
		fmt.Println(strconv.Itoa(h.minimumHeap[h.realSize]) + "]")
	}
}

func (h maxHeap) size() int{
	return h.realSize
}

func (h maxHeap) print() {
	if h.realSize == 0 {
		fmt.Println("No element")
	} else {
		fmt.Print("[")
		for i := 1; i<= h.realSize-1; i++ {
			fmt.Print(h.maximumHeap[i], ",")
		}
		fmt.Println(strconv.Itoa(h.maximumHeap[h.realSize]) + "]")
	}
}

func main() {
	// Test case
	minH := newMinHeap(3)
	minH.add(3)
	minH.add(1)
	minH.add(2)
	minH.print()

	fmt.Println(minH.peak())
	fmt.Println(minH.pop())
	fmt.Println(minH.size())
	minH.add(4)
	minH.add(5)

	minH.print()

	maxH := newMaxHeap(3)
	maxH.add(3)
	maxH.add(1)
	maxH.add(2)
	maxH.print()

	fmt.Println(maxH.peak())
	fmt.Println(maxH.pop())
	fmt.Println(maxH.size())
	maxH.add(4)
	maxH.add(5)

	maxH.print()
}
