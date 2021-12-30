package main

import (
	"fmt"
	"math"
)

const MaxInt = int(^uint(0) >> 1)
const MaxInt64 = int(^uint64(0) >> 1)

//brute force
func smallestRepunitDivByK(k int) int {

	num := 1
	exp := 1
	var found bool
	for num > 0 && num < MaxInt {
		if num%k == 0 {
			found = true
			break
		}
		num += int(math.Pow(10, float64(exp)))
		exp++
	}
	if !found {
		return -1
	}
	return exp
}

//Pigeonhole Principle
func smallestRepunitDivByK2(k int) int {
	var remainder int
	for length := 1; length <= k; length++ {
		remainder = (remainder*10 + 1) % k
		if remainder == 0 {
			return length
		}
	}
	return -1
}

func main() {
	fmt.Println(smallestRepunitDivByK2(7))
}
