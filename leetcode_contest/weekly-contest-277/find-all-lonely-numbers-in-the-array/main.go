package main

import "fmt"

func findLonely(nums []int) []int {

	m := make(map[int]int)
	var ret []int
	for _, v := range nums {
		m[v] += 1
		m[v-1] += 1
		m[v+1] += 1
	}
	for _, v := range nums {
		if m[v] == 1 {
			ret = append(ret, v)
		}
	}
	return ret
}

func main() {
	nums := []int{61, 83, 92, 92, 42, 60, 16, 45, 32, 14, 40, 7, 10, 34, 62, 33, 65, 79, 7, 14, 85, 21, 36, 5, 99, 25, 0, 14, 52, 41, 40}
	fmt.Println(findLonely(nums))
}
