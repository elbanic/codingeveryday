package main

import "fmt"

func minimumOperations(nums []int) int {

	m := make(map[int]int)
	var maxCount, maxZigzag int
	for i := 0; i < len(nums)-1; i++ {
		if _, ok := m[nums[i]]; !ok {
			var count int
			for j := i + 2; j < len(nums); j += 2 {
				if nums[i] == nums[j] {
					count++
				}
			}
			if count > 0 {
				m[nums[i]] = i
			}
			if maxCount < count {
				maxCount = count
				maxZigzag = nums[i]
			}
		}
	}

	fmt.Println(maxZigzag)
	//left alter

	//right alter

	return 0
}

func main() {
	nums := []int{69, 91, 47, 74, 75, 94, 22, 100, 43, 50, 82, 47, 40, 51, 90, 27, 98, 85, 47, 14, 55, 82, 52, 9, 65, 90, 86, 45, 52, 52, 95, 40, 85, 3, 46, 77, 16, 59, 32, 22, 41, 87, 89, 78, 59, 78, 34, 26, 71, 9, 82, 68, 80, 74, 100, 6, 10, 53, 84, 80, 7, 87, 3, 82, 26, 26, 14, 37, 26, 58, 96, 73, 41, 2, 79, 43, 56, 74, 30, 71, 6, 100, 72, 93, 83, 40, 28, 79, 24}
	fmt.Println(minimumOperations(nums))
}
