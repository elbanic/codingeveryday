package main

import "fmt"

func validMountainArray(arr []int) bool {

	if len(arr) < 3 {
		return false
	}

	incline := arr[0] < arr[1]
	if !incline {
		return false
	}

	decline := false
	j := 0
	for i := 1; i < len(arr)-1; i++ {
		if incline {
			if arr[i] >= arr[i+1] {
				decline = true
				j = i
				break
			}
		}
	}

	if !decline {
		return false
	}
	for j < len(arr)-1 {
		if arr[j] <= arr[j+1] {
			return false
		}
		j++
	}
	return incline && decline
}

func main() {
	arr := []int{0, 1, 3, 2}
	fmt.Println(validMountainArray(arr))

	arr2 := []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(arr2))
}
