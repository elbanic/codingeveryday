package main

import (
	"fmt"
	"math"
)

func isHappy(n int) bool {

	visited := make(map[int]bool)
	MAX_INT := int(^uint(0) >> 1)

	for i := 0; i < MAX_INT; i++ {
		if visited[n] {
			return false
		} else {
			digit := int(math.Log10(float64(n))) + 1
			var sum int
			cur := n
			for i := 0; i < digit; i++ {
				val := cur % 10
				sum += int(math.Pow(float64(val), float64(2)))
				cur = cur / 10
			}
			if sum == 1 {
				return true
			}
			visited[n] = true
			n = sum
		}
	}
	return false
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}
