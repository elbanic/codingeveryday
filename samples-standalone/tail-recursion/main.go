package main

import "fmt"

func helperTailRecursion(start int, num []int, acc int) int {
	if start >= len(num) {
		return acc
	}
	return helperTailRecursion(start + 1, num, acc + num[start])
}

func sumTailRecursion(num []int) int {
	if num == nil || len(num) == 0 {
		return 0
	}
	sum := helperTailRecursion(0, num, 0)
	return sum
}

func helperNonTailRecursion(start int, num []int) int {
	if start >= len(num) {
		return 0
	}
	return num[start] + helperNonTailRecursion(start + 1, num)
}
func sumNonTailRecursion(num []int) int{
	if num == nil || len(num) == 0 {
		return 0
	}
	sum := helperNonTailRecursion(0, num)
	return sum
}

func main() {

	a1 := []int{1,2,3,4,5}
	s1 := sumNonTailRecursion(a1)
	s2 := sumTailRecursion(a1)

	fmt.Println(s1, s2)
}
