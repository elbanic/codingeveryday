package main

import "fmt"

func countEven(num int) int {

	var count int
	for i := 2; i <= num; i++ {
		if sumDigit(i)%2 == 0 {
			count++
		}
	}
	return count
}

func sumDigit(num int) int {

	var sum int
	for num >= 1 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func main() {
	fmt.Println(countEven(14))
}
