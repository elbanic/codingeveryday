package main

import "fmt"

func sumOfThree(num int64) []int64 {

	start := num / 3
	if num%3 == 0 {
		return []int64{start - 1, start, start + 1}
	}

	curSum := sum([]int64{start - 1, start, start + 1})
	prev := int64(1)
	next := int64(4)
	for curSum < num {

		curSum -= prev
		curSum += next

		if curSum == num {
			return []int64{next - 2, next - 1, next}
		}
		prev++
		next++
	}
	return []int64{}
}

func sum(nums []int64) int64 {
	var sum int64
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {

	fmt.Println(sumOfThree(33))
	fmt.Println(sumOfThree(0))
	fmt.Println(sumOfThree(920474140))
}
