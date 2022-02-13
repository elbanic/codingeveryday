package main

import "fmt"

func countOperations(num1 int, num2 int) int {
	return helper(num1, num2, 0)
}

func helper(num1, num2, count int) int {
	if num1 == 0 || num2 == 0 {
		return count
	}
	count++
	if num1 >= num2 {
		num1 -= num2
	} else {
		num2 -= num1
	}
	return helper(num1, num2, count)
}

func main() {
	fmt.Println(countOperations(2, 3))
	fmt.Println(countOperations(10, 10))
}
