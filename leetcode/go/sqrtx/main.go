package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	for i := 1; i <= x; i++ {
		if i*i > x {
			return i - 1
		} else if i*i == x {
			return i
		}
	}
	return 0
}

//binary search
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}

	start := 1
	end := x

	for start <= end {
		cur := start + (end-start)/2
		if x == cur*cur {
			return cur
		}
		if x > (cur-1)*(cur-1) && x < (cur)*(cur) {
			return cur - 1
		} else if x > cur*cur {
			start = cur + 1
		} else {
			end = cur - 1
		}
	}
	return 0
}

func main() {
	fmt.Println(mySqrt2(8))
}
