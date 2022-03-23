package main

import "fmt"

//binary search
func fixedPoint2(arr []int) int {
	left, right := 0, len(arr)-1
	answer := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == mid {
			answer = mid
			right = mid - 1
		} else if arr[mid] < mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return answer
}

//brute force
func fixedPoint(arr []int) int {

	for i := range arr {
		if i == arr[i] {
			return i
		}
		if i < arr[i] {
			return -1
		}
	}
	return -1
}

func main() {
	arr := []int{-10, -5, 0, 3, 7}
	fmt.Println(fixedPoint2(arr))
}
