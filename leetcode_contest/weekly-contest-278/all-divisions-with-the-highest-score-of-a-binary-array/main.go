package main

import "fmt"

//dp
func maxScoreIndices2(nums []int) []int {
	scores := make([]int, len(nums)+1)
	scores[0] = getScore(1, nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			scores[i+1] = scores[i] - 1
		} else {
			scores[i+1] = scores[i] + 1
		}
	}

	max := scores[0]
	var maxVal []int
	for i, v := range scores {
		if v == max {
			maxVal = append(maxVal, i)
		} else if v > max {
			max = v
			maxVal = []int{i}
		}
	}
	return maxVal
}

//brute force
func maxScoreIndices(nums []int) []int {

	var max int
	var maxVal []int

	for i := 0; i < len(nums)+1; i++ {
		left := nums[:i]
		var right []int
		if i < len(nums) {
			right = nums[i:]
		} else {
			right = []int{}
		}
		score := getScore(0, left) + getScore(1, right)
		if score != 0 && score == max {
			maxVal = append(maxVal, i)
		} else if score > max {
			max = score
			maxVal = []int{i}
		}
	}
	return maxVal
}

func getScore(targetVal int, nums []int) int {
	var sum int
	for _, v := range nums {
		if v == targetVal {
			sum++
		}
	}
	return sum
}

func main() {
	nums := []int{0, 0, 1, 0}
	fmt.Println(maxScoreIndices2(nums))

	nums2 := []int{0, 0, 0}
	fmt.Println(maxScoreIndices2(nums2))

	nums3 := []int{1, 1}
	fmt.Println(maxScoreIndices2(nums3))
}
