package main

import "fmt"

func addDigits(num int) int {

	for num >= 10 {
		cur := num
		var sum int
		for cur >= 1 {
			digits := cur % 10
			cur = cur / 10
			sum += digits
		}
		num = sum
	}
	return num
}

func main() {
	fmt.Println(addDigits(38))
}
