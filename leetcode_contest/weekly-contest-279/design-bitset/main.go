package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Bitset struct {
	arr       []int
	oneCount  int
	all       bool
	isUpdated bool
}

func Constructor(size int) Bitset {
	return Bitset{arr: make([]int, size)}
}

func (this *Bitset) Fix(idx int) {
	if idx >= len(this.arr) {
		return
	}
	this.arr[idx] = 1
	this.isUpdated = true
}

func (this *Bitset) Unfix(idx int) {
	if idx >= len(this.arr) {
		return
	}
	this.arr[idx] = 0
	this.isUpdated = true
}

func (this *Bitset) Flip() {

	if len(this.arr)/10 < 5 {
		for i := range this.arr {
			if this.arr[i] == 0 {
				this.arr[i] = 1
			} else {
				this.arr[i] = 0
			}
		}
	} else {
		var wg sync.WaitGroup
		var subArrs [][]int
		chunkSize := len(this.arr) / 2

		for i := 0; i < len(this.arr); i += chunkSize {
			end := i + chunkSize
			if end > len(this.arr) {
				end = len(this.arr)
			}
			subArrs = append(subArrs, this.arr[i:end])
		}

		for _, v := range subArrs {
			wg.Add(1)
			go func(slice []int) {
				for j := range slice {
					if slice[j] == 0 {
						slice[j] = 1
					} else {
						slice[j] = 0
					}
				}
				wg.Done()
			}(v)
		}
		wg.Wait()
	}
	this.isUpdated = true
}

func (this *Bitset) All() bool {
	if !this.isUpdated {
		return this.all
	}
	this.all = true
	for i := 0; i < len(this.arr); i++ {
		if this.arr[i] == 0 {
			this.all = false
			break
		}
	}
	return this.all
}

func (this *Bitset) One() bool {
	if !this.isUpdated {
		return this.oneCount > 0
	}
	for i := 0; i < len(this.arr); i++ {
		if this.arr[i] == 1 {
			return true
		}
	}
	return false
}

func (this *Bitset) Count() int {
	if !this.isUpdated {
		return this.oneCount
	}
	this.oneCount = 0
	for i := 0; i < len(this.arr); i++ {
		if this.arr[i] == 1 {
			this.oneCount++
		}
	}
	return this.oneCount
}

func (this *Bitset) ToString() string {
	var ret string
	for _, v := range this.arr {
		ret += strconv.Itoa(v)
	}
	return ret
}

/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */

func main() {
	bs := Constructor(5)
	bs.Fix(3)
	bs.Fix(1)
	bs.Flip()
	fmt.Println(bs.All())
	bs.Unfix(0)
	bs.Flip()
	fmt.Println(bs.One())
	bs.Unfix(0)
	fmt.Println(bs.Count())
	fmt.Println(bs.ToString())
}
