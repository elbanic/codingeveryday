package main

import (
	"fmt"
	"sort"
)

func minimumSum(num int) int {

	var arr []int
	for num >= 1 {
		arr = append(arr, num%10)
		num = num / 10
	}
	sort.Ints(arr)
	return (arr[0]*10 + arr[2]) + (arr[1]*10 + arr[3])
}

func main() {
	fmt.Println(minimumSum(1298))
	fmt.Println(minimumSum(4009))
}
