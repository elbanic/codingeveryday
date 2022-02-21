package main

import "fmt"

func majorityElement(nums []int) int {

	N := len(nums)
	m := make(map[int]int)

	for _, i := range nums {
		m[i]++
		if m[i] > N/2 {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
