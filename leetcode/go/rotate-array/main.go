package main

import "fmt"

//reverse
func rotate3(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		temp := nums[len(nums)-1-i]
		nums[len(nums)-1-i] = nums[i]
		nums[i] = temp
	}
}

//Cyclic Replacements
func rotate2(nums []int, k int) {
	k = k % len(nums)
	var count int

	for start := 0; count < len(nums); start++ {
		current := start
		prev := nums[start]

		for {
			next := (current + k) % len(nums)
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++
			if start == current {
				break
			}
		}
	}
}

func rotate(nums []int, k int) {

	if len(nums) <= k {
		if k%len(nums) == 0 {
			return
		}
		k = k % len(nums)
	}

	id := len(nums) - k
	temp := make([]int, k)

	copy(temp, nums[id:])
	temp = append(temp, nums[:id]...)
	copy(nums, temp)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 2

	rotate3(nums, k)
	fmt.Println(nums)
}
