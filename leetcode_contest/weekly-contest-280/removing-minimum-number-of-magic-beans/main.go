package main

import (
	"fmt"
	"math"
)

//brute force
func minimumRemoval(beans []int) int64 {

	var maxNum int
	m := make(map[int]int)
	for i := range beans {
		m[beans[i]]++
		if beans[i] > maxNum {
			maxNum = beans[i]
		}
	}

	var sum1, sum2 int
	for i := range beans {
		if beans[i] != maxNum {
			sum1 += beans[i]
		}
	}
	sum2 = sum1
	for i := range beans {
		var tempSum int
		for j := range beans {
			if beans[j] > beans[i] {
				tempSum += beans[j] - beans[i]
			} else if beans[j] < beans[i] {
				tempSum += beans[j]
			}
		}
		if tempSum < sum2 {
			sum2 = tempSum
		}
	}
	return int64(math.Min(float64(sum1), float64(sum2)))
}

func main() {
	beans := []int{63, 43, 12, 94}
	fmt.Println(minimumRemoval(beans))

	beans2 := []int{66, 90, 47, 25, 92, 90, 76, 85, 22, 3}
	fmt.Println(minimumRemoval(beans2))
}
