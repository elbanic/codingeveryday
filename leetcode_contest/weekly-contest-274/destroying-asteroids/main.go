package main

import (
	"fmt"
	"sort"
)

//iteration
func asteroidsDestroyed2(mass int, asteroids []int) bool {
	if len(asteroids) == 0 {
		return false
	}
	sort.Ints(asteroids)
	for len(asteroids) > 0 {
		id := -1
		for i := range asteroids {
			if asteroids[i] <= mass {
				id = i
				break
			}
		}
		if id != -1 {
			mass += asteroids[id]
			asteroids = append(asteroids[:id], asteroids[id+1:]...)
		} else {
			return false
		}
	}
	return true
}

func bsearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] < target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if nums[start] < target {
		return start
	}
	return -1
}

//recursion
func asteroidsDestroyed(mass int, asteroids []int) bool {
	return helper(asteroids, mass)
}

func helper(asteroids []int, sum int) bool {
	if len(asteroids) == 0 {
		return true
	}

	id := -1
	for i := range asteroids {
		if sum >= asteroids[i] {
			id = i
			break
		}
	}
	if id != -1 {
		sum += asteroids[id]
		asteroids = append(asteroids[:id], asteroids[id+1:]...)
		return helper(asteroids, sum)
	}
	return false
}

func main() {
	mass := 10
	asteroids := []int{3, 9, 19, 5, 21}
	fmt.Println(asteroidsDestroyed2(mass, asteroids))

	mass2 := 5
	asteroids2 := []int{4, 9, 23, 4}
	fmt.Println(asteroidsDestroyed2(mass2, asteroids2))

}
