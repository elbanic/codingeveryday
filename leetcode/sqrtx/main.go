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

func main() {
	fmt.Println(mySqrt(1))
}
