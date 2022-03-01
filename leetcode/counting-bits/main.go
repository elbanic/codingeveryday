package main

import "fmt"

func countBits(n int) []int {

	var ret []int
	for i := 0; i <= n; i++ {
		ret = append(ret, intBitsCount(i))
	}
	return ret
}

func intBitsCount(n int) int {
	if n == 0 {
		return n
	}

	var cnt int
	for n >= 2 {
		if n%2 == 1 {
			cnt++
		}
		n /= 2
	}
	return cnt + 1
}

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
	fmt.Println(countBits(7))
	fmt.Println(countBits(10))
}
